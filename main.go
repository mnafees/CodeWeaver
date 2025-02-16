package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-git/go-git/v5"
)

var version = "v0.0.8"

func main() {
	// Define command line flags
	dirPath := flag.String("input", ".", "Directory to scan")
	outputFileName := flag.String("output", "codebase.md", "Output file name")
	ignorePatterns := flag.String("ignore", `\.git.*`, "Comma-separated list of regular expression patterns that match the paths to be ignored")
	includedPathsFile := flag.String("included-paths-file", "", "File to save included paths (optional). If provided, the included paths will be saved to the file and not printed to the console.")
	excludedPathsFile := flag.String("excluded-paths-file", "", "File to save excluded paths (optional). If provided, the excluded paths will be saved to the file and not printed to the console.")
	respectGitIgnore := flag.Bool("respect-git-ignore", false, "If true, the git ignore file will be respected and the respective paths will be ignored.")
	showVersion := flag.Bool("version", false, "Show version and exit")
	showHelp := flag.Bool("help", false, "Show help message and exit")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	// Check if help flag is set or no arguments are provided
	if *showHelp || len(os.Args) == 1 {
		printHelp()
		return
	}

	// Split ignore patterns string into a slice
	ignoreListString := strings.Split(*ignorePatterns, ",")
	ignoreList := make([]*regexp.Regexp, len(ignoreListString))

	if *respectGitIgnore {
		// Check if we are in a git repository
		repo, err := git.PlainOpen(".")
		if err != nil {
			fmt.Println("Not a git repository:", err)
		} else {
			// Get the .gitignore patterns
			_, err := repo.Worktree()
			if err != nil {
				fmt.Println("Error accessing worktree:", err)
			} else {
				// Read .gitignore files from the repository
				gitIgnorePatterns, err := readGitIgnorePatterns()
				if err != nil {
					fmt.Println("Error reading .gitignore files:", err)
				} else {
					for _, pattern := range gitIgnorePatterns {
						ignoreList = append(ignoreList, regexp.MustCompile(pattern))
					}
				}
			}
		}
	}

	fmt.Println("Patterns to ignore:")
	for i, pattern := range ignoreListString {
		fmt.Println(ignoreListString[i])
		ignoreList[i] = regexp.MustCompile(strings.TrimSpace(pattern))
	}

	// Create the output file
	outputFile, err := os.Create(*outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the codebase tree to the output file
	fmt.Fprintln(outputFile, "# Tree View:\n```")
	fmt.Fprintf(outputFile, "%s\n", *dirPath)

	var depthOpen map[int]bool
	depthOpen = make(map[int]bool)
	err = printTree(*dirPath, 0, depthOpen, ignoreList, outputFile)
	if err != nil {
		fmt.Println("Error printing codebase tree:", err)
		return
	}
	fmt.Fprintln(outputFile, "```")

	// Write the code content to the output file
	fmt.Fprintln(outputFile, "\n# Content:\n")
	err = writeCodeContent(*dirPath, ignoreList, outputFile, *includedPathsFile, *excludedPathsFile)
	if err != nil {
		fmt.Println("Error writing code content:", err)
		return
	}

	fmt.Println("Codebase documentation generated successfully!")
}

// printTree recursively walks the directory tree and prints the structure to the output file
func printTree(dirPath string, depth int, depthOpen map[int]bool, ignoreList []*regexp.Regexp, outputFile *os.File) error {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	// Filter out ignored files and directories
	var filteredFiles []fs.DirEntry
	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		relPath, _ := filepath.Rel(".", filePath)
		if !shouldIgnore(relPath, ignoreList) {
			filteredFiles = append(filteredFiles, file)
		}
	}

	for i, file := range filteredFiles { // Iterate over filtered files
		filePath := filepath.Join(dirPath, file.Name())

		var pipe string = "├─"
		depthOpen[depth] = true
		if i == len(filteredFiles)-1 { // Use filteredFiles length
			pipe = "└─"
			depthOpen[depth] = false
		}

		var indent = []rune("")
		if depth > 0 {
			indent = []rune(strings.Repeat("  ", depth))
			for j := 0; j < depth; j++ {
				if depthOpen[j] {
					indent[j*2] = '│'
				}
			}
		}

		if file.IsDir() {
			fmt.Fprintf(outputFile, "%s%s%s\n", string(indent), pipe, file.Name())
			printTree(filePath, depth+1, depthOpen, ignoreList, outputFile)
			depthOpen[depth] = false
		} else {
			fmt.Fprintf(outputFile, "%s%s%s\n", string(indent), pipe, file.Name())
		}
	}

	return nil
}

// writeCodeContent reads the content of each file and writes it to the output file within a code block
func writeCodeContent(dirPath string, ignoreList []*regexp.Regexp, outputFile *os.File, includedPathsFile, excludedPathsFile string) error {
	var Red = "\033[31m"
	var Green = "\033[32m"
	var Reset = "\033[0m"
	var includedPaths []string
	var excludedPaths []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Check if the file should be ignored
		relPath, _ := filepath.Rel(".", path)
		if shouldIgnore(relPath, ignoreList) {
			if excludedPathsFile == "" {
				fmt.Println(Red + "- " + path + Reset)
			} else {
				excludedPaths = append(excludedPaths, path)
			}
			return nil
		}

		if includedPathsFile == "" {
			fmt.Println(Green + "+ " + path + Reset)
		} else {
			includedPaths = append(includedPaths, path)
		}

		if !d.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			extension := filepath.Ext(path)
			extension = strings.ToLower(extension)
			extension = strings.TrimPrefix(extension, ".")
			fmt.Fprintf(outputFile, "## %s\n", path)
			fmt.Fprintf(outputFile, "```%s\n%s\n```\n\n", extension, content)
		}

		return nil
	})

	// Save included paths to file (if filename provided)
	if includedPathsFile != "" {
		err = savePathsToFile(includedPathsFile, includedPaths)
		if err != nil {
			return fmt.Errorf("error saving included paths to file: %w", err)
		}
	}

	// Save excluded paths to file (if filename provided)
	if excludedPathsFile != "" {
		err = savePathsToFile(excludedPathsFile, excludedPaths)
		if err != nil {
			return fmt.Errorf("error saving excluded paths to file: %w", err)
		}
	}

	return err
}

// shouldIgnore checks if a given path should be ignored based on the ignore patterns
func shouldIgnore(path string, ignoreList []*regexp.Regexp) bool {
	if path == "." {
		return true
	}
	for _, pattern := range ignoreList {
		if pattern.MatchString(path) {
			return true
		}
	}
	return false
}

// printHelp prints the help message
func printHelp() {
	fmt.Println("Usage: go run codemerge.go [options]")
	fmt.Println("\nOptions:")
	flag.PrintDefaults()
}

// savePathsToFile saves a list of paths to a file, one per line
func savePathsToFile(filename string, paths []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, path := range paths {
		_, err := fmt.Fprintln(file, path)
		if err != nil {
			return err
		}
	}

	return nil
}

// readGitIgnorePatterns reads .gitignore files and returns a list of patterns
func readGitIgnorePatterns() ([]string, error) {
	var patterns []string
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			gitIgnorePath := filepath.Join(path, ".gitignore")
			if _, err := os.Stat(gitIgnorePath); err == nil {
				file, err := os.Open(gitIgnorePath)
				if err != nil {
					return err
				}
				defer file.Close()

				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					line := strings.TrimSpace(scanner.Text())
					if line != "" && !strings.HasPrefix(line, "#") {
						patterns = append(patterns, line)
					}
				}
				if err := scanner.Err(); err != nil {
					return err
				}
			}
		}
		return nil
	})
	return patterns, err
}
