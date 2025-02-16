# Tree View:
```
.
├─LICENSE
├─README.md
├─build_and_run.ps1
├─codebase.md
├─go.mod
├─go.sum
├─goreleaser.yaml
├─main.go
└─test_root
  ├─File at root A.txt
  ├─File at root B.md
  ├─folder 01
  │ ├─File at folder 01 I.txt
  │ ├─File at folder 01 II.md
  │ └─File at folder 01 III.csv
  └─folder 02
    ├─File at folder 02 I.txt
    ├─File at folder 02 II.md
    ├─File at folder 02 III.csv
    └─folder 02 01
      ├─File at folder 02 01 I.txt
      ├─File at folder 02 01 II.md
      └─File at folder 02 01 III.csv
```

# Content:

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

# CodeWeaver: Generate a Markdown Document of Your Codebase Structure and Content

CodeWeaver is a command-line tool designed to weave your codebase into a single, easy-to-navigate Markdown document. It recursively scans a directory, generating a structured representation of your project's file hierarchy and embedding the content of each file within code blocks. This tool simplifies codebase sharing, documentation, and integration with AI/ML code analysis tools by providing a consolidated and readable Markdown output.
The output for the current repository can be found [here](https://github.com/tesserato/CodeWeaver/blob/main/codebase.md).

# Key Features

* **Comprehensive Codebase Documentation:** Generates a Markdown file that meticulously outlines your project's directory and file structure in a clear, tree-like format.
* **Code Content Inclusion:** Embeds the complete content of each file directly within the Markdown document, enclosed in syntax-highlighted code blocks based on file extensions.
* **Flexible Path Filtering:**  Utilize regular expressions to define ignore patterns, allowing you to exclude specific files and directories from the generated documentation (e.g., `.git`, build artifacts, specific file types).
* **Optional Path Logging:** Choose to save lists of included and excluded file paths to separate files for detailed tracking and debugging of your ignore rules.
* **Simple Command-Line Interface:**  Offers an intuitive command-line interface with straightforward options for customization.

# Installation

If you have Go installed, run `go install github.com/tesserato/CodeWeaver@latest`to install the latest version of CodeWeaver or `go install github.com/tesserato/CodeWeaver@vX.Y.Z` to install a specific version.

Alternatively, download the appropriate pre built executable from the [releases page](https://github.com/tesserato/CodeWeaver/releases).

If necessary, make the `codeweaver` executable by using the `chmod` command:

```bash
chmod +x codeweaver
```

# Usage

## For help, run
```bash
codeweaver -h
```

## For actual usage, run
```bash
codeweaver [options]
```

**Options:**

| Option                            | Description                                                               | Default Value           |
| --------------------------------- | ------------------------------------------------------------------------- | ----------------------- |
| `-dir <directory>`                | The root directory to scan and document.                                  | Current directory (`.`) |
| `-output <filename>`              | The name of the output Markdown file.                                     | `codebase.md`           |
| `-ignore "<regex patterns>"`      | Comma-separated list of regular expression patterns for paths to exclude. | `\.git.*`               |
| `-included-paths-file <filename>` | File to save the list of paths that were included in the documentation.   | None                    |
| `-excluded-paths-file <filename>` | File to save the list of paths that were excluded from the documentation. | None                    |
| `-help`                           | Display this help message and exit.                                       |                         |

# Examples

## **Generate documentation for the current directory:**

   ```bash
   ./codeweaver
   ```
   This will create a file named `codebase.md` in the current directory, documenting the structure and content of the current directory and its subdirectories (excluding paths matching the default ignore pattern `\.git.*`).

## **Specify a different input directory and output file:**

   ```bash
   ./codeweaver -dir=my_project -output=project_docs.md
   ```
   This command will process the `my_project` directory and save the documentation to `project_docs.md`.

## **Ignore specific file types and directories:**

   ```bash
   ./codeweaver -ignore="\.log,temp,build" -output=detailed_docs.md
   ```
   This example will generate `detailed_docs.md`, excluding any files or directories with names containing `.log`, `temp`, or `build`. Regular expression patterns are comma-separated.

## **Save lists of included and excluded paths:**

   ```bash
   ./codeweaver -ignore="node_modules" -included-paths-file=included.txt -excluded-paths-file=excluded.txt -output=code_overview.md
   ```
   This command will create `code_overview.md` while also saving the list of included paths to `included.txt` and the list of excluded paths (due to the `node_modules` ignore pattern) to `excluded.txt`.

# Contributing

Contributions are welcome! If you encounter any issues, have suggestions for new features, or want to improve CodeWeaver, please feel free to open an issue or submit a pull request on the project's GitHub repository.

# License

CodeWeaver is released under the [MIT License](LICENSE). See the `LICENSE` file for complete license details.
```

## build_and_run.ps1
```ps1
go build .

git describe --tags --abbrev=0

./CodeWeaver -version

./CodeWeaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"
```

## codebase.md
```md
# Tree View:
```
.
├─LICENSE
├─README.md
├─build_and_run.ps1
├─codebase.md
├─go.mod
├─go.sum
├─goreleaser.yaml
├─main.go
└─test_root
  ├─File at root A.txt
  ├─File at root B.md
  ├─folder 01
  │ ├─File at folder 01 I.txt
  │ ├─File at folder 01 II.md
  │ └─File at folder 01 III.csv
  └─folder 02
    ├─File at folder 02 I.txt
    ├─File at folder 02 II.md
    ├─File at folder 02 III.csv
    └─folder 02 01
      ├─File at folder 02 01 I.txt
      ├─File at folder 02 01 II.md
      └─File at folder 02 01 III.csv
```

# Content:

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

# CodeWeaver: Generate a Markdown Document of Your Codebase Structure and Content

CodeWeaver is a command-line tool designed to weave your codebase into a single, easy-to-navigate Markdown document. It recursively scans a directory, generating a structured representation of your project's file hierarchy and embedding the content of each file within code blocks. This tool simplifies codebase sharing, documentation, and integration with AI/ML code analysis tools by providing a consolidated and readable Markdown output.
The output for the current repository can be found [here](https://github.com/tesserato/CodeWeaver/blob/main/codebase.md).

# Key Features

* **Comprehensive Codebase Documentation:** Generates a Markdown file that meticulously outlines your project's directory and file structure in a clear, tree-like format.
* **Code Content Inclusion:** Embeds the complete content of each file directly within the Markdown document, enclosed in syntax-highlighted code blocks based on file extensions.
* **Flexible Path Filtering:**  Utilize regular expressions to define ignore patterns, allowing you to exclude specific files and directories from the generated documentation (e.g., `.git`, build artifacts, specific file types).
* **Optional Path Logging:** Choose to save lists of included and excluded file paths to separate files for detailed tracking and debugging of your ignore rules.
* **Simple Command-Line Interface:**  Offers an intuitive command-line interface with straightforward options for customization.

# Installation

If you have Go installed, run `go install github.com/tesserato/CodeWeaver@latest`to install the latest version of CodeWeaver or `go install github.com/tesserato/CodeWeaver@vX.Y.Z` to install a specific version.

Alternatively, download the appropriate pre built executable from the [releases page](https://github.com/tesserato/CodeWeaver/releases).

If necessary, make the `codeweaver` executable by using the `chmod` command:

```bash
chmod +x codeweaver
```

# Usage

## For help, run
```bash
codeweaver -h
```

## For actual usage, run
```bash
codeweaver [options]
```

**Options:**

| Option                            | Description                                                               | Default Value           |
| --------------------------------- | ------------------------------------------------------------------------- | ----------------------- |
| `-dir <directory>`                | The root directory to scan and document.                                  | Current directory (`.`) |
| `-output <filename>`              | The name of the output Markdown file.                                     | `codebase.md`           |
| `-ignore "<regex patterns>"`      | Comma-separated list of regular expression patterns for paths to exclude. | `\.git.*`               |
| `-included-paths-file <filename>` | File to save the list of paths that were included in the documentation.   | None                    |
| `-excluded-paths-file <filename>` | File to save the list of paths that were excluded from the documentation. | None                    |
| `-help`                           | Display this help message and exit.                                       |                         |

# Examples

## **Generate documentation for the current directory:**

   ```bash
   ./codeweaver
   ```
   This will create a file named `codebase.md` in the current directory, documenting the structure and content of the current directory and its subdirectories (excluding paths matching the default ignore pattern `\.git.*`).

## **Specify a different input directory and output file:**

   ```bash
   ./codeweaver -dir=my_project -output=project_docs.md
   ```
   This command will process the `my_project` directory and save the documentation to `project_docs.md`.

## **Ignore specific file types and directories:**

   ```bash
   ./codeweaver -ignore="\.log,temp,build" -output=detailed_docs.md
   ```
   This example will generate `detailed_docs.md`, excluding any files or directories with names containing `.log`, `temp`, or `build`. Regular expression patterns are comma-separated.

## **Save lists of included and excluded paths:**

   ```bash
   ./codeweaver -ignore="node_modules" -included-paths-file=included.txt -excluded-paths-file=excluded.txt -output=code_overview.md
   ```
   This command will create `code_overview.md` while also saving the list of included paths to `included.txt` and the list of excluded paths (due to the `node_modules` ignore pattern) to `excluded.txt`.

# Contributing

Contributions are welcome! If you encounter any issues, have suggestions for new features, or want to improve CodeWeaver, please feel free to open an issue or submit a pull request on the project's GitHub repository.

# License

CodeWeaver is released under the [MIT License](LICENSE). See the `LICENSE` file for complete license details.
```

## build_and_run.ps1
```ps1
go build .

git describe --tags --abbrev=0

./CodeWeaver -version

./CodeWeaver -ignore="\.git.*,.+\.exe,codebase.md,excluded_paths.txt" -excluded-paths-file="excluded_paths.txt"
```


```

## go.mod
```mod
module github.com/tesserato/CodeWeaver

go 1.23.0

require (
	github.com/go-git/go-billy/v5 v5.6.2
	github.com/go-git/go-git/v5 v5.13.2
)

require (
	dario.cat/mergo v1.0.1 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/ProtonMail/go-crypto v1.1.5 // indirect
	github.com/cloudflare/circl v1.6.0 // indirect
	github.com/cyphar/filepath-securejoin v0.4.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/pjbgf/sha1cd v0.3.2 // indirect
	github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 // indirect
	github.com/skeema/knownhosts v1.3.1 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
)

```

## go.sum
```sum
dario.cat/mergo v1.0.1 h1:Ra4+bf83h2ztPIQYNP99R6m+Y7KfnARDfID+a+vLl4s=
dario.cat/mergo v1.0.1/go.mod h1:uNxQE+84aUszobStD9th8a29P2fMDhsBdgRYvZOxGmk=
github.com/Microsoft/go-winio v0.5.2/go.mod h1:WpS1mjBmmwHBEWmogvA2mj8546UReBk4v8QkMxJ6pZY=
github.com/Microsoft/go-winio v0.6.2 h1:F2VQgta7ecxGYO8k3ZZz3RS8fVIXVxONVUPlNERoyfY=
github.com/Microsoft/go-winio v0.6.2/go.mod h1:yd8OoFMLzJbo9gZq8j5qaps8bJ9aShtEA8Ipt1oGCvU=
github.com/ProtonMail/go-crypto v1.1.5 h1:eoAQfK2dwL+tFSFpr7TbOaPNUbPiJj4fLYwwGE1FQO4=
github.com/ProtonMail/go-crypto v1.1.5/go.mod h1:rA3QumHc/FZ8pAHreoekgiAbzpNsfQAosU5td4SnOrE=
github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be h1:9AeTilPcZAjCFIImctFaOjnTIavg87rW78vTPkQqLI8=
github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be/go.mod h1:ySMOLuWl6zY27l47sB3qLNK6tF2fkHG55UZxx8oIVo4=
github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 h1:0CwZNZbxp69SHPdPJAN/hZIm0C4OItdklCFmMRWYpio=
github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5/go.mod h1:wHh0iHkYZB8zMSxRWpUBQtwG5a7fFgvEO+odwuTv2gs=
github.com/cloudflare/circl v1.6.0 h1:cr5JKic4HI+LkINy2lg3W2jF8sHCVTBncJr5gIIq7qk=
github.com/cloudflare/circl v1.6.0/go.mod h1:uddAzsPgqdMAYatqJ0lsjX1oECcQLIlRpzZh3pJrofs=
github.com/cyphar/filepath-securejoin v0.4.1 h1:JyxxyPEaktOD+GAnqIqTf9A8tHyAG22rowi7HkoSU1s=
github.com/cyphar/filepath-securejoin v0.4.1/go.mod h1:Sdj7gXlvMcPZsbhwhQ33GguGLDGQL7h7bg04C/+u9jI=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1 h1:vj9j/u1bqnvCEfJOwUhtlOARqs3+rkHYY13jYWTU97c=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/elazarl/goproxy v1.4.0 h1:4GyuSbFa+s26+3rmYNSuUVsx+HgPrV1bk1jXI0l9wjM=
github.com/elazarl/goproxy v1.4.0/go.mod h1:X/5W/t+gzDyLfHW4DrMdpjqYjpXsURlBt9lpBDxZZZQ=
github.com/emirpasic/gods v1.18.1 h1:FXtiHYKDGKCW2KzwZKx0iC0PQmdlorYgdFG9jPXJ1Bc=
github.com/emirpasic/gods v1.18.1/go.mod h1:8tpGGwCnJ5H4r6BWwaV6OrWmMoPhUl5jm/FMNAnJvWQ=
github.com/gliderlabs/ssh v0.3.8 h1:a4YXD1V7xMF9g5nTkdfnja3Sxy1PVDCj1Zg4Wb8vY6c=
github.com/gliderlabs/ssh v0.3.8/go.mod h1:xYoytBv1sV0aL3CavoDuJIQNURXkkfPA/wxQ1pL1fAU=
github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 h1:+zs/tPmkDkHx3U66DAb0lQFJrpS6731Oaa12ikc+DiI=
github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376/go.mod h1:an3vInlBmSxCcxctByoQdvwPiA7DTK7jaaFDBTtu0ic=
github.com/go-git/go-billy/v5 v5.6.2 h1:6Q86EsPXMa7c3YZ3aLAQsMA0VlWmy43r6FHqa/UNbRM=
github.com/go-git/go-billy/v5 v5.6.2/go.mod h1:rcFC2rAsp/erv7CMz9GczHcuD0D32fWzH+MJAU+jaUU=
github.com/go-git/go-git-fixtures/v4 v4.3.2-0.20231010084843-55a94097c399 h1:eMje31YglSBqCdIqdhKBW8lokaMrL3uTkpGYlE2OOT4=
github.com/go-git/go-git-fixtures/v4 v4.3.2-0.20231010084843-55a94097c399/go.mod h1:1OCfN199q1Jm3HZlxleg+Dw/mwps2Wbk9frAWm+4FII=
github.com/go-git/go-git/v5 v5.13.2 h1:7O7xvsK7K+rZPKW6AQR1YyNhfywkv7B8/FsP3ki6Zv0=
github.com/go-git/go-git/v5 v5.13.2/go.mod h1:hWdW5P4YZRjmpGHwRH2v3zkWcNl6HeXaXQEMGb3NJ9A=
github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 h1:f+oWsMOmNPc8JmEHVZIycC7hBoQxHH9pNKQORJNozsQ=
github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8/go.mod h1:wcDNUvekVysuuOpQKo3191zZyTpiI6se1N1ULghS0sw=
github.com/google/go-cmp v0.6.0 h1:ofyhxvXcZhMsU5ulbFiLKl/XBFqE1GSq7atu8tAmTRI=
github.com/google/go-cmp v0.6.0/go.mod h1:17dUlkBOakJ0+DkrSSNjCkIjxS6bF9zb3elmeNGIjoY=
github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 h1:BQSFePA1RWJOlocH6Fxy8MmwDt+yVQYULKfN0RoTN8A=
github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99/go.mod h1:1lJo3i6rXxKeerYnT8Nvf0QmHCRC1n8sfWVwXF2Frvo=
github.com/kevinburke/ssh_config v1.2.0 h1:x584FjTGwHzMwvHx18PXxbBVzfnxogHaAReU4gf13a4=
github.com/kevinburke/ssh_config v1.2.0/go.mod h1:CT57kijsi8u/K/BOFA39wgDQJ9CxiF4nAY/ojJ6r6mM=
github.com/kr/pretty v0.1.0/go.mod h1:dAy3ld7l9f0ibDNOQOHHMYYIIbhfbHSm3C4ZsoJORNo=
github.com/kr/pretty v0.3.1 h1:flRD4NNwYAUpkphVc1HcthR4KEIFJ65n8Mw5qdRn3LE=
github.com/kr/pretty v0.3.1/go.mod h1:hoEshYVHaxMs3cyo3Yncou5ZscifuDolrwPKZanG3xk=
github.com/kr/pty v1.1.1/go.mod h1:pFQYn66WHrOpPYNljwOMqo10TkYh1fy3cYio2l3bCsQ=
github.com/kr/text v0.1.0/go.mod h1:4Jbv+DJW3UT/LiOwJeYQe1efqtUx/iVham/4vfdArNI=
github.com/kr/text v0.2.0 h1:5Nx0Ya0ZqY2ygV366QzturHI13Jq95ApcVaJBhpS+AY=
github.com/kr/text v0.2.0/go.mod h1:eLer722TekiGuMkidMxC/pM04lWEeraHUUmBw8l2grE=
github.com/onsi/gomega v1.34.1 h1:EUMJIKUjM8sKjYbtxQI9A4z2o+rruxnzNvpknOXie6k=
github.com/onsi/gomega v1.34.1/go.mod h1:kU1QgUvBDLXBJq618Xvm2LUX6rSAfRaFRTcdOeDLwwY=
github.com/pjbgf/sha1cd v0.3.2 h1:a9wb0bp1oC2TGwStyn0Umc/IGKQnEgF0vVaZ8QF8eo4=
github.com/pjbgf/sha1cd v0.3.2/go.mod h1:zQWigSxVmsHEZow5qaLtPYxpcKMMQpa09ixqBxuCS6A=
github.com/pkg/errors v0.9.1 h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=
github.com/pkg/errors v0.9.1/go.mod h1:bwawxfHBFNV+L2hUp1rHADufV3IMtnDRdf1r5NINEl0=
github.com/pmezard/go-difflib v1.0.0 h1:4DBwDE0NGyQoBHbLQYPwSUPoCMWR5BEzIk/f1lZbAQM=
github.com/pmezard/go-difflib v1.0.0/go.mod h1:iKH77koFhYxTK1pcRnkKkqfTogsbg7gZNVY4sRDYZ/4=
github.com/rogpeppe/go-internal v1.12.0 h1:exVL4IDcn6na9z1rAb56Vxr+CgyK3nn3O+epU5NdKM8=
github.com/rogpeppe/go-internal v1.12.0/go.mod h1:E+RYuTGaKKdloAfM02xzb0FW3Paa99yedzYV+kq4uf4=
github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3 h1:n661drycOFuPLCN3Uc8sB6B/s6Z4t2xvBgU1htSHuq8=
github.com/sergi/go-diff v1.3.2-0.20230802210424-5b0b94c5c0d3/go.mod h1:A0bzQcvG0E7Rwjx0REVgAGH58e96+X0MeOfepqsbeW4=
github.com/sirupsen/logrus v1.7.0/go.mod h1:yWOB1SBYBC5VeMP7gHvWumXLIWorT60ONWic61uBYv0=
github.com/skeema/knownhosts v1.3.1 h1:X2osQ+RAjK76shCbvhHHHVl3ZlgDm8apHEHFqRjnBY8=
github.com/skeema/knownhosts v1.3.1/go.mod h1:r7KTdC8l4uxWRyK2TpQZ/1o5HaSzh06ePQNxPwTcfiY=
github.com/stretchr/objx v0.1.0/go.mod h1:HFkY916IF+rwdDfMAkV7OtwuqBVzrE8GR6GFx+wExME=
github.com/stretchr/testify v1.2.2/go.mod h1:a8OnRcib4nhh0OaRAV+Yts87kKdq0PP7pXfy6kDkUVs=
github.com/stretchr/testify v1.4.0/go.mod h1:j7eGeouHqKxXV5pUuKE4zz7dFj8WfuZ+81PSLYec5m4=
github.com/stretchr/testify v1.10.0 h1:Xv5erBjTwe/5IxqUQTdXv5kgmIvbHo3QQyRwhJsOfJA=
github.com/stretchr/testify v1.10.0/go.mod h1:r2ic/lqez/lEtzL7wO/rwa5dbSLXVDPFyf8C91i36aY=
github.com/xanzy/ssh-agent v0.3.3 h1:+/15pJfg/RsTxqYcX6fHqOXZwwMP+2VyYWJeWM2qQFM=
github.com/xanzy/ssh-agent v0.3.3/go.mod h1:6dzNDKs0J9rVPHPhaGCukekBHKqfl+L3KghI1Bc68Uw=
golang.org/x/crypto v0.0.0-20220622213112-05595931fe9d/go.mod h1:IxCIyHEi3zRg3s0A5j5BB6A9Jmi73HwBIUl50j+osU4=
golang.org/x/crypto v0.33.0 h1:IOBPskki6Lysi0lo9qQvbxiQ+FvsCC/YWOecCHAixus=
golang.org/x/crypto v0.33.0/go.mod h1:bVdXmD7IV/4GdElGPozy6U7lWdRXA4qyRVGJV57uQ5M=
golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56 h1:2dVuKD2vS7b0QIHQbpyTISPd0LeHDbnYEryqj5Q1ug8=
golang.org/x/exp v0.0.0-20240719175910-8a7402abbf56/go.mod h1:M4RDyNAINzryxdtnbRXRL/OHtkFuWGRjvuhBJpk2IlY=
golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2/go.mod h1:9nx3DQGgdP8bBQD5qxJ1jj9UTztislL4KSBs9R2vV5Y=
golang.org/x/net v0.35.0 h1:T5GQRQb2y08kTAByq9L4/bz8cipCdA8FbRTXewonqY8=
golang.org/x/net v0.35.0/go.mod h1:EglIi67kWsHKlRzzVMUD93VMSWGFOMSZgxFjparz1Qk=
golang.org/x/sys v0.0.0-20191026070338-33540a1f6037/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20201119102817-f84b799fce68/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210423082822-04245dca01da/go.mod h1:h1NjWce9XRLGQEsW7wpKNCjG9DtNlClVuFLEZdDNbEs=
golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8/go.mod h1:oPkhp1MJrh7nUepCBck5+mAzfO9JrbApNNgaTdGDITg=
golang.org/x/sys v0.30.0 h1:QjkSwP/36a20jFYWkSue1YwXzLmsV5Gfq7Eiy72C1uc=
golang.org/x/sys v0.30.0/go.mod h1:/VUhepiaJMQUp4+oa/7Zr1D23ma6VTLIYjOOTFZPUcA=
golang.org/x/term v0.0.0-20201126162022-7de9c90e9dd1/go.mod h1:bj7SfCRtBDWHUb9snDiAeCFNEtKQo2Wmx5Cou7ajbmo=
golang.org/x/term v0.29.0 h1:L6pJp37ocefwRRtYPKSWOWzOtWSxVajvz2ldH/xi3iU=
golang.org/x/term v0.29.0/go.mod h1:6bl4lRlvVuDgSf3179VpIxBF0o10JUpXWOnI7nErv7s=
golang.org/x/text v0.3.6/go.mod h1:5Zoc/QRtKVWzQhOtBMvqHzDpF6irO9z98xDceosuGiQ=
golang.org/x/text v0.22.0 h1:bofq7m3/HAFvbF51jz3Q9wLg3jkvSPuiZu/pD1XwgtM=
golang.org/x/text v0.22.0/go.mod h1:YRoo4H8PVmsu+E3Ou7cqLVH8oXWIHVoX0jqUWALQhfY=
golang.org/x/tools v0.0.0-20180917221912-90fa682c2a6e/go.mod h1:n7NCudcB/nEzxVGmLbDWY5pfWTLqBcC2KZ6jyYvM4mQ=
gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15/go.mod h1:Co6ibVJAznAaIkqp8huTwlJQCZ016jof/cbN4VW5Yz0=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c h1:Hei/4ADfdWqJk1ZMxUNpqntNwaWcugrBjAiHlqqRiVk=
gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c/go.mod h1:JHkPIbrfpd72SG/EVd6muEfDQjcINNoR0C8j2r3qZ4Q=
gopkg.in/warnings.v0 v0.1.2 h1:wFXVbFY8DY5/xOe1ECiWdKCzZlxgshcYVNkBHstARME=
gopkg.in/warnings.v0 v0.1.2/go.mod h1:jksf8JmL6Qr/oQM2OXTHunEvvTAsrWBLb6OOjuVWRNI=
gopkg.in/yaml.v2 v2.2.2/go.mod h1:hI93XBmqTisBFMUTm0b8Fm+jr3Dg1NNxqwp+5A1VGuI=
gopkg.in/yaml.v2 v2.4.0/go.mod h1:RDklbk79AGWmwhnvt/jBztapEOGDOx6ZbXqjP6csGnQ=
gopkg.in/yaml.v3 v3.0.1 h1:fxVm/GzAzEWqLHuvctI91KS9hhNmmWOoWu0XTYJS7CA=
gopkg.in/yaml.v3 v3.0.1/go.mod h1:K4uyk7z7BCEPqu6E+C64Yfv1cQ7kz7rIZviUmN+EgEM=

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
    ldflags:
      - -s -w
      - -X main.version={{.Version}}
      - -X main.commit={{.Commit}}
      - -X main.date={{.Date}}

archives:
  - formats: [tar.gz]
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
        formats: [ 'zip' ]

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

```

## test_root/File at root A.txt
```txt

```

## test_root/File at root B.md
```md

```

## test_root/folder 01/File at folder 01 I.txt
```txt

```

## test_root/folder 01/File at folder 01 II.md
```md

```

## test_root/folder 01/File at folder 01 III.csv
```csv

```

## test_root/folder 02/File at folder 02 I.txt
```txt

```

## test_root/folder 02/File at folder 02 II.md
```md

```

## test_root/folder 02/File at folder 02 III.csv
```csv

```

## test_root/folder 02/folder 02 01/File at folder 02 01 I.txt
```txt

```

## test_root/folder 02/folder 02 01/File at folder 02 01 II.md
```md

```

## test_root/folder 02/folder 02 01/File at folder 02 01 III.csv
```csv

```

