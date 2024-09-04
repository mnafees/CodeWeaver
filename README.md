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
