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
- **Node & npm** (LTS version)


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


## How does it work ??? 


**Launch commands**
- `init()`: is the first function executed before main().
It adds the content of `createCmd` to the content of `rootCmd` by the cobra method `AddCommand()`
To add more commands to le CLI, its the place where you must had them.

- `main()`: is the second function executed after `init()`.
It lauches the `Execute()` cobra function.

- `Execute()`: is a methode from cobra library.
It reads what is after the instruction `archi` and run the right function.
(exemple : Use = command name, RunE = function to execute)



**Create workflow**
- `runCreate()`: Execute a list of function that manage prompts, stock the results, and execute the result.

*Prompts :*
1 - `PromptProjectName()`
2 - `PromptArchitecture()`
3 - `PromptLanguage()`
4 - `PromptExpress()`

It stores the result of each prompt in an object `config` to use it into the last function to create the project `GenerateProject()`

*Execution :*
1 - `RunNpmInit()`
2 - `InstallDependencies()`
3 - `InstallPnpm()`

**Project Generation**
- `GenerateProject()`: is the main function that manage the whole creation of folders `createDirectoryStructure()` and base files `generateBaseFiles()`

- `createDirectoryStrucure()`: is calling the template function choose by the user and create them `CreateDirectory()`

- `generateBaseFile()`: is calling the template function to create base file with  `WriteFile()`


## 💖 Sponsor

If you find this project valuable and would like to support its development:

[Donate with PayPal](https://www.paypal.com/donate/?hosted_button_id=UMJGHGGV4YUHE)

Your support helps me maintain and improve this project. Thank you! 🙏