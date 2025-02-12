# Codebase Structure

.
├─LICENSE
├─README.md
├─build_and_run.ps1
├─go.mod
├─goreleaser.yaml
├─main.go
└─test_root
  ├─File at root A.txt
  ├─File at root B.md
  ├─folder 01
  │ ├─File at folder 01 I.txt
  │ ├─File at folder 01 II.md
  │ └─File at folder 01 III.csv
  ├─folder 02
  │ ├─File at folder 02 I.txt
  │ ├─File at folder 02 II.md
  │ ├─File at folder 02 III.csv
  │ └─folder 02 01
  │   ├─File at folder 02 01 I.txt
  │   ├─File at folder 02 01 II.md
  │   └─File at folder 02 01 III.csv
  └─folder 03

# Code Content

## LICENSE
```
MIT License

Copyright (c) 2024 Carlos Tarjano

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```

## README.md
```md
## CodeWeaver: Weave your codebase into a single, navigable Markdown document.

CodeWeaver is a command-line tool that helps you generate comprehensive Markdown documentation from your code repositories. Easily share your codebase structure and content with colleagues, technical writers, or for use in AI/ML code analysis tools. 

### Features

* **Structured Documentation:**  CodeWeaver creates a Markdown file with a clear tree-like structure representing the directories and files in your project.
* **Code Highlighting:** Code content is embedded in code blocks with syntax highlighting for improved readability.
* **Path Filtering:** Exclude unwanted directories or files using regular expression patterns.
* **Log Included/Excluded Paths:** Optionally save lists of included and excluded paths to files for reference.
* **Easy to Use:**  Simple command-line interface with clear options.

### Installation 

1. **Make sure you have Go installed:** [https://go.dev/doc/install](https://go.dev/doc/install)
2. **Download CodeWeaver:**  You can download the pre-built binary from the releases page (or build it from source).
3. **Make it executable (if necessary):**  `chmod +x codeweaver`

### Usage

```bash
codeweaver [options]
```

**Options:**

* `-dir <directory>`: The directory to scan (default: current directory).
* `-output <filename>`:  The name of the output Markdown file (default: "codebase.md").
* `-ignore "<regex pattern>,<regex pattern>,..."`: Comma-separated list of regular expression patterns to exclude files or directories.
* `-included-paths-file <filename>`:  Save the list of included paths to the specified file.
* `-excluded-paths-file <filename>`: Save the list of excluded paths to the specified file.
* `-help`:  Show this help message. 

**Example:**

To generate documentation for the "myproject" directory and save it to "docs.md", ignoring any ".log" files and the "temp" directory:

```bash
codeweaver -dir=myproject -output=docs.md -ignore="\.log,temp"
```

### Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on GitHub.

### License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details. 

```

## build_and_run.ps1
```ps1
# mit the symbol table, debug information and the DWARF symbol table by passing -s and -w go build -ldflags="-s -w" .
go build # -ldflags="-s -w" .

./codeweaver -h

./codeweaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"

# "\.git.*,.+\.exe,.+\.[Ii][Cc][Oo],.+\.[Pp][Nn][Gg],.+\.[Jj][Pp][Ee]?[Gg],.+\.[Mm][Pp][34],.+\.[Aa][Vv][Ii],.+\.[Mm][Oo][Vv],.+\.[Ww][Mm][Aa],.+\.[Pp][Dd][Ff],.+\.[Dd][Oo][Cc][Xx]?,.+\.[Xx][Ll][Ss][Xx]?,.+\.[Pp][Pp][Tt][Xx]?,.+\.[Zz][Ii][Pp],.+\.[Rr][Aa][Rr],.+\.[7][Zz],.+\.[Ii][Ss][Oo],.+\.[Bb][Ii][Nn],.+\.[Dd][Aa][Tt],.+\.[Dd][Mm][Gg],.+\.[Gg][Ii][Ff],.+\.[Tt][Ii][Ff][Ff]?,.+\.[Pp][Ss][Dd],.+\.[Mm][Kk][Vv],.+\.[Ff][Ll][Aa][Cc],.+\.[Oo][Gg][Gg],.+\.[Ww][Ee][Bb][Pp],.+\.[Aa][Vv][Ii][Ff],.+\.[Hh][Ee][Ii][Cc],.+\.[Jj][Xx][Ll]"
```

## go.mod
```mod
module github.com/tesserato/CodeWeaver

go 1.23.0

```

## goreleaser.yaml
```yaml
# vim: set ts=2 sw=2 tw=0 fo=cnqoj

version: 2

builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin

archives:
  - format: tar.gz
    # this name template makes the OS and Arch compatible with the results of `uname`.
    name_template: >-
      {{ .ProjectName }}_
      {{- title .Os }}_
      {{- if eq .Arch "amd64" }}x86_64
      {{- else if eq .Arch "386" }}i386
      {{- else }}{{ .Arch }}{{ end }}
      {{- if .Arm }}v{{ .Arm }}{{ end }}
    # use zip for windows archives
    format_overrides:
      - goos: windows
        format: zip

changelog:
  sort: asc
  filters:
    exclude:
      - "^docs:"
      - "^test:"

release:
  footer: >-

    ---

    Released by [GoReleaser](https://github.com/goreleaser/goreleaser).

```

## main.go
```go
package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Define command line flags
	dirPath := flag.String("input", ".", "Directory to scan")
	outputFileName := flag.String("output", "codebase.md", "Output file name")
	ignorePatterns := flag.String("ignore", `\.git.*`, "Comma-separated list of regular expression patterns that match the paths to be ignored")
	includedPathsFile := flag.String("included-paths-file", "", "File to save included paths (optional). If provided, the included paths will be saved to the file and not printed to the console.")
	excludedPathsFile := flag.String("excluded-paths-file", "", "File to save excluded paths (optional). If provided, the excluded paths will be saved to the file and not printed to the console.")
	showHelp := flag.Bool("help", false, "Show help message")

	flag.Parse()

	// Check if help flag is set or no arguments are provided
	if *showHelp || len(os.Args) == 1 {
		printHelp()
		return
	}

	// Split ignore patterns string into a slice
	ignoreListString := strings.Split(*ignorePatterns, ",")
	ignoreList := make([]*regexp.Regexp, len(ignoreListString))

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
	fmt.Fprintln(outputFile, "# Codebase Structure\n")
	fmt.Fprintf(outputFile, "%s\n", *dirPath)

	var depthOpen map[int]bool
	depthOpen = make(map[int]bool)
	err = printTree(*dirPath, 0, depthOpen, ignoreList, outputFile)
	if err != nil {
		fmt.Println("Error printing codebase tree:", err)
		return
	}

	// Write the code content to the output file
	fmt.Fprintln(outputFile, "\n# Code Content\n")
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

	for i, file := range files {
		filePath := filepath.Join(dirPath, file.Name())
		relPath, _ := filepath.Rel(".", filePath)

		// Check if the file/directory should be ignored
		if shouldIgnore(relPath, ignoreList) {
			continue
		}

		var pipe string = "├─"
		depthOpen[depth] = true
		if i == len(files)-1 {
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

```

## test_root\File at root A.txt
```txt

```

## test_root\File at root B.md
```md

```

## test_root\folder 01\File at folder 01 I.txt
```txt

```

## test_root\folder 01\File at folder 01 II.md
```md

```

## test_root\folder 01\File at folder 01 III.csv
```csv

```

## test_root\folder 02\File at folder 02 I.txt
```txt

```

## test_root\folder 02\File at folder 02 II.md
```md

```

## test_root\folder 02\File at folder 02 III.csv
```csv

```

## test_root\folder 02\folder 02 01\File at folder 02 01 I.txt
```txt

```

## test_root\folder 02\folder 02 01\File at folder 02 01 II.md
```md

```

## test_root\folder 02\folder 02 01\File at folder 02 01 III.csv
```csv

```

