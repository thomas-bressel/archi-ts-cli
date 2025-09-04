# archits-cli

> A powerful CLI tool for TypeScript/JavaScript project architecture management

[![npm version](https://badge.fury.io/js/archits-cli.svg)](https://badge.fury.io/js/archits-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Archi-TS-CLI is a Go-based command-line tool designed to simplify project architecture management for TypeScript and JavaScript projects. It provides features to generate, organize, and analyze project structures with clean architecture patterns.

## 🚀 Quick Start


### Global Installation (Recommended)
```bash
# Install globally
npm install -g archits-cli

# Create a symlink for easier access (if using nvm)
ln -sf /home/[your-username]/.nvm/versions/node/v[your-node-version]/lib/node_modules/archits-cli/bin/archits /home/[your-username]/.nvm/versions/node/v[your-node-version]/bin/archi

# Verify installation
archits --version

# Use anywhere
archits create
archits generate entity User
archits --help
```

## 📋 Features

- **Automatic project structure generation** - Create clean architecture projects instantly
- **Entity generation** - Generate controller, service, repository, entity, and route files
- **TypeScript & JavaScript support** - Choose your preferred language
- **Multiple architecture patterns** - Support for different project architectures
- **Express.js integration** - Optional Express setup
- **Jest configuration** - Pre-configured testing setup
- **Clean code templates** - Following best practices

## 🛠️ Usage

### Create New Project
```bash
npx archits-cli create
```

This will prompt you to choose:
- Project name
- Language (TypeScript/JavaScript)  
- Architecture type
- Express.js integration
- And more...

### Generate Entity Files
```bash
npx archits-cli generate entity Product
# Aliases: npx archits-cli g e Product
```

Generates:
- `Product.controller.ts`
- `Product.service.ts` 
- `Product.repository.ts`
- `Product.entity.ts`
- `Product.route.ts`
- Unit tests for each layer

### Other Commands
```bash
# Show version
npx archits-cli version

# Show help
npx archits-cli --help
```

## 📁 Generated Project Structure

```
my-project/
├── src/
│   ├── controllers/
│   ├── services/
│   ├── repositories/
│   ├── entities/
│   ├── routes/
│   └── index.ts
├── tests/
├── package.json
├── tsconfig.json
├── jest.config.ts
├── .env
└── .gitignore
```

## 🔧 Requirements

- Node.js >= 14
- npm or yarn

## 📖 Examples

### Creating a User Management Project
```bash
npx archits-cli create
# Follow the prompts...
cd my-project
npm install
npm run dev
```

### Adding a new Entity
```bash
cd my-project
npx archits-cli generate entity Order
# Files generated in appropriate folders
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/thomas-bressel/archi-ts-cli/blob/main/LICENSE) file for details.

## 👨‍💻 Author

**Thomas Bressel**
- GitHub: [@thomas-bressel](https://github.com/thomas-bressel)

## 💖 Support

If you find this project helpful, please consider:

[![PayPal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://www.paypal.com/donate/?hosted_button_id=UMJGHGGV4YUHE)

## 🔗 Links

- [Source Code](https://github.com/thomas-bressel/archi-ts-cli)
- [Issues](https://github.com/thomas-bressel/archi-ts-cli/issues)
- [Discussions](https://github.com/thomas-bressel/archi-ts-cli/discussions)

---

Made with ❤️ by Thomas Bressel