# Archi-TS-CLI

## Description
Archi-TS-CLI is a Go-based command-line tool (CLI) designed to simplify project architecture management. It provides features to generate, organize, and analyze project structures.

## Features
- Automatic generation of project structures.
- Analysis and validation of existing architectures.
- Customizable configurations for various project types.

## Development Prerequisites
- **Go** (version 1.18 or higher)
- **Linux** (or use WSL on windows systems)
- **Node & npm** (LTS version)


## Development Build Instructions
Clone the repository and build the project:
```bash
git clone https://github.com/thomas-bressel/archi-ts-cli.git
cd archi-ts-cli
make install
source ~/.bashrc
```


# ArchiTS Installation for Users

## Prerequisites
- **Node.js** and **npm** must be installed on your system
- **For Windows users**: Consider using WSL (Windows Subsystem for Linux) for a smoother experience

## Installation Steps

### 1. Install ArchiTS CLI globally
```bash
npm install -g archits-cli
```

### 2. Create a shorter alias (optional but recommended)
If you're using **nvm** (Node Version Manager), create a symlink for easier usage:

```bash
ln -sf /home/[your-username]/.nvm/versions/node/v[your-node-version]/lib/node_modules/archits-cli/bin/archits /home/[your-username]/.nvm/versions/node/v[your-node-version]/bin/archi
```

**Replace:**
- `[your-username]` with your actual username
- `[your-node-version]` with your Node.js version (e.g., `18.17.0`)

### 3. Verify installation
```bash
archits --version
# or if you created the alias:
archi --version
```

## Alternative: Using alias instead of symlink
Add this line to your `~/.bashrc` or `~/.zshrc`:
```bash
alias archi='archits'
```

Then reload your shell:
```bash
source ~/.bashrc  # or source ~/.zshrc
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

To generate files from an entity name:
```bash
archi generate entity <entity-name>
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

## Code details about the development of Archi-TS-CLI:

To help my open source co-workers, please check this link :
https://github.com/thomas-bressel/archi-ts-cli/discussions


## 💖 Sponsor

If you find this project valuable and would like to support its development:

[Donate with PayPal](https://www.paypal.com/donate/?hosted_button_id=UMJGHGGV4YUHE)

Your support helps me maintain and improve this project. Thank you! 🙏


## License

This project is licensed under the [MIT License](./LICENSE).  
© 2025 Thomas Bressel
