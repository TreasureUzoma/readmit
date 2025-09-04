This project, named `readmit`, is a command-line interface (CLI) application written in Go that assists in generating various project files, such as READMEs, contribution guidelines, and commit messages. It operates by reading the local codebase, sending the content to a backend API for generation, and then writing the results back to the local file system or printing them to the console.

### Features

- **README Generation:** Automatically generates a `README.md` file for your project based on its content.
- **Contribution Guide Generation:** Creates a `CONTRIBUTION.md` file to guide contributors.
- **Commit Message Generation:** Generates concise commit messages for your changes.
- **Intelligent File Reading:** Recursively reads project files while intelligently ignoring common development artifacts, version control directories, temporary files, and other non-essential content based on a comprehensive ignore list.
- **Backend Integration:** Utilizes a backend API for processing and generating content, ensuring secure handling of project data.

### Installation

To install `readmit`, ensure you have Go installed on your system. Then, you can use `go install`:

```bash
go install
```

This will install the `readmit` executable in your `$GOPATH/bin` directory.

### Usage

The primary command for `readmit` is `generate`, which requires a `--type` flag to specify the kind of file to generate.

```bash
readmit generate [file_type]
```

#### Available `file_type` options:

- `readme`: Generates a `README.md` file in the current directory.
- `contribution`: Generates a `CONTRIBUTION.md` file in the current directory.
- `commit`: Generates a commit message and prints it to the console.

#### Examples:

**Generate a README.md:**

```bash
readmit generate readme
```

**Generate a CONTRIBUTION.md:**

```bash
readmit generate contribution
```

**Generate a commit message:**

```bash
readmit generate commit
```

### How it Works

The `readmit` application performs the following steps when a generation command is executed:

1.  **Reads Project Files:** It recursively walks through the current directory, reading the contents of relevant files. It skips directories and files that match a predefined set of ignore patterns (e.g., `node_modules`, `.git`, `dist`, `*.log`, `*.zip`, image files).
2.  **Buffers Content:** The read file contents are stored in an in-memory buffer, with each file prefixed by `=== [filename] ===`.
3.  **Obtains Signed URL:** The application calls a backend API (`http://localhost:3000/api/upload-url`) to get a pre-signed URL for uploading the buffered project content.
4.  **Uploads Content:** The in-memory buffer containing the project files is uploaded to the obtained signed URL.
5.  **Calls Generation API:** It then calls another backend API (`http://localhost:3000/api/generate`), passing the filename (from the upload) and the requested generation mode (e.g., "readme").
6.  **Receives and Writes Content:** The generated content is received from the backend.
    - For `readme` and `contribution` types, the content is written to `README.md` and `CONTRIBUTION.md` respectively in the current directory.
    - For `commit` type, the content is printed directly to the console.

### Ignoring Files

The application includes a comprehensive list of `IgnorePatterns` to exclude various files and directories from being processed. These include:

- Package manager directories (`node_modules`, `vendor`, `Pods`)
- Build artifacts (`dist`, `build`, `bin`, `.next`)
- Version control system directories (`.git`, `.svn`, `.idea`, `.vscode`)
- Environment and config files (`.env`, `.secrets`)
- Log files (`*.log`, `npm-debug.log*`)
- Dependency lock files (`go.mod`, `package-lock.json`)
- Archives and compressed files (`*.zip`, `*.tar.gz`)
- Binary executables (`*.exe`, `*.dll`, `*.class`)
- Disk images and installers (`*.iso`, `*.dmg`)
- Various media files (images, audio, video, fonts)
- Office documents (`*.doc`, `*.pdf`)
- Backup files (`*.bak`, `*~`)
- Coverage and report files (`coverage`, `.nyc_output`)
- Cloud and IDE-specific directories (`.vercel`, `.github`)
- Miscellaneous development junk (`.eslintcache`, `__pycache__`)
- Database files and dumps (`*.sqlite`, `*.dump`)

### Dependencies

The project uses the following Go modules:

- `github.com/cpuguy83/go-md2man/v2`
- `github.com/inconshreveable/mousetrap`
- `github.com/russross/blackfriday/v2`
- `github.com/spf13/cobra`
- `github.com/spf13/pflag`
- `gopkg.in/check.v1`
- `gopkg.in/yaml.v3`

### License

Copyright © 2025 TREASURE UZOMA <EMAIL ADDRESS>
Please add the specific licensing information here.

### Contact

TREASURE UZOMA <EMAIL ADDRESS>
