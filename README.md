# Archi-TS-CLI

## Description
Archi-TS-CLI is a Go-based command-line tool (CLI) designed to simplify project architecture management. It provides features to generate, organize, and analyze project structures.

## Features
- Automatic generation of project structures.
- Analysis and validation of existing architectures.
- Customizable configurations for various project types.

## Prerequisites
- **Go** (version 1.18 or higher)
- **Linux** (or use WSL on windows systems)

## Installation
Clone the repository and build the project:
```bash
git clone https://github.com/thomas-bressel/archi-ts-cli.git
cd archi-ts-cli
make install
source ~/.bashrc
```

## Usage
Run the following command to display help:
```bash
archi --help
```

### Example
To generate a project structure:
```bash
archi create
```

## Contribution
Contributions are welcome! Please follow these steps:
1. Fork the project.
2. Create a branch for your changes: `git checkout -b feature/your-feature-name`.  (feature soon to come !)
3. Submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for details.

## Author
Developed by [Thomas Bressel].

## Dependencies
- **Go Modules** for dependency management.
- **cobra** for command-line interface.
- **fmt** for formatted I/O.
- **os** for operating system functionality.
- **path/filepath** for file path manipulation.
- **prompui** for user prompts.
- **color** for colored output : go get github.com/fatih/color


____________________________________

## features descriptions :

**Project Structure**
- `main.go`: Entry point of the CLI application.
- `root.go`: A very important file in whitch createCmd variable is init.

**Prompts commands**
- `prompts/prompts.go`: the creation of a prompts starts here in a named function.
- `cmd/create.go`: to use your prompt functions, have a look into the runCreate function.
