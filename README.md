# Readmit: AI-Powered Documentation Generator

Readmit is a powerful tool designed to streamline your documentation process by leveraging Artificial Intelligence to automatically generate various project files, including READMEs, CONTRIBUTING guides, and insightful commit messages. It consists of a Command-Line Interface (CLI) application written in Go and an AI-powered backend service built with Next.js and TypeScript. By analyzing your codebase and Git history, Readmit ensures your project documentation is always comprehensive, accurate, and up-to-date, allowing developers to focus on coding rather than manual documentation tasks.

## Features

- **Automated README.md Generation:** Generate comprehensive `README.md` files for your projects, covering essential sections like installation, usage, and project overview.
- **Automated CONTRIBUTING.md Generation:** Create clear and professional `CONTRIBUTING.md` files to guide new contributors on how to engage with your project.
- **Intelligent Commit Message Generation:** Get concise and conventional Git commit message suggestions based on your staged, unstaged, or last commit changes, adhering to standards like Conventional Commits.
- **Dockerfile Generation:** Generate optimized `Dockerfile`s tailored to your project's technology stack and dependencies.
- **Smart Codebase Analysis:** The CLI intelligently reads and processes your project files, automatically ignoring irrelevant content such as build artifacts, node modules, `.git` directories, temporary files, and various media files, ensuring only pertinent code is sent for analysis.
- **Scalable AI Integration:** Utilizes the Google GenAI service on the backend for robust and intelligent content generation.

## Stacks / Technologies

| Category    | Technology         | Link                                                                   | Description                                                                                                                                                    |
| :---------- | :----------------- | :--------------------------------------------------------------------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **CLI**     | Go                 | [https://golang.org/](https://golang.org/)                             | The primary language for the command-line interface, providing performance and concurrency.                                                                    |
|             | Cobra              | [https://cobra.dev/](https://cobra.dev/)                               | A library for creating powerful modern CLI applications in Go, used for command parsing and application structure.                                             |
|             | `net/http`         | [https://pkg.go.dev/net/http](https://pkg.go.dev/net/http)             | Go's standard library for making HTTP requests to the Readmit backend API.                                                                                     |
|             | `os/exec`          | [https://pkg.go.dev/os/exec](https://pkg.go.dev/os/exec)               | Used to execute external commands, specifically Git, for extracting diff information.                                                                          |
| **Web/API** | Next.js            | [https://nextjs.org/](https://nextjs.org/)                             | React framework used for the Readmit documentation site and hosting the AI generation API endpoints.                                                           |
|             | TypeScript         | [https://www.typescriptlang.org/](https://www.typescriptlang.org/)     | Provides static typing to JavaScript, enhancing code quality and maintainability across the web application and API.                                           |
|             | Tailwind CSS       | [https://tailwindcss.com/](https://tailwindcss.com/)                   | A utility-first CSS framework for rapidly building custom designs, used for styling the Readmit documentation website.                                         |
|             | Supabase           | [https://supabase.com/](https://supabase.com/)                         | An open-source Firebase alternative, utilized for secure file storage and generating signed URLs for codebase uploads.                                         |
|             | Google GenAI       | [https://ai.google.dev/](https://ai.google.dev/)                       | The AI service powering the content generation capabilities, providing intelligent analysis and text generation.                                               |
|             | MDX                | [https://mdxjs.com/](https://mdxjs.com/)                               | Allows writing JSX in Markdown documents, used for creating rich and interactive documentation content on the website.                                         |
|             | Radix UI/Shadcn UI | [https://www.radix-ui.com/](https://www.radix-ui.com/)                 | A set of unstyled, accessible UI components used as a foundation for the Readmit documentation website's user interface.                                       |
|             | Husky              | [https://typicode.github.io/husky/](https://typicode.github.io/husky/) | Git hooks for ensuring code quality by running scripts (like `post-process.sh` to generate `search-data/documents.json`) before committing or pushing changes. |

## Installation

To install the `readmit` CLI, ensure you have Go installed on your system (version 1.16 or higher).

1.  **Install the CLI:**

    ```bash
    go install github.com/treasureuzoma/readmit/readmit@latest
    ```

    This command downloads the latest version of the Readmit CLI and installs the `readmit` executable in your `$GOPATH/bin` directory. Ensure your `$GOPATH/bin` is included in your system's `PATH` environment variable.

2.  **Verify Installation:**
    You can verify the installation by running:
    ```bash
    readmit --help
    ```
    This should display the help message for the Readmit CLI.

## Usage

The `readmit` CLI provides a `generate` command to create various types of documentation files.

```bash
readmit generate [type]
```

### Available `type` options:

- `readme`: Generates a `README.md` file in the current directory.
- `contribution`: Generates a `CONTRIBUTING.md` file in the current directory.
- `commit`: Generates a commit message and prints it to the console (based on your Git diff).
- `dockerfile`: Generates an optimized `Dockerfile` for your project.

### Examples:

**Generate a `README.md` file:**

```bash
readmit generate readme
```

**Generate a `CONTRIBUTING.md` file:**

```bash
readmit generate contribution
```

**Generate a commit message:**

```bash
readmit generate commit
```

**Generate a `Dockerfile`:**

```bash
readmit generate dockerfile
```

## Contributing

We welcome contributions to Readmit! Whether you want to report a bug, suggest a feature, or submit code, please follow these guidelines:

1.  **Report Bugs:** If you find a bug, please open an issue on our [GitHub repository](https://github.com/treasureuzoma/readmit/issues) with a clear description of the problem and steps to reproduce it.
2.  **Suggest Features:** Have an idea for a new feature or improvement? Open an issue to discuss it with the community.
3.  **Submit Pull Requests:**
    - Fork the repository and create a new branch for your feature or bug fix.
    - Ensure your code adheres to the project's coding standards.
    - Write clear and concise commit messages.
    - Submit a pull request with a detailed description of your changes.

[![Readme was generated by Readmit](https://img.shields.io/badge/Readme%20was%20generated%20by-Readmit-brightred)](https://readmit.vercel.app)
