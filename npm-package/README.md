# ArchiTS-CLI

> A powerful CLI tool for TypeScript project architecture management
> 🌐 [Visit www.archits.dev](https://www.archits.dev)

[![npm version](https://badge.fury.io/js/archits-cli.svg)](https://badge.fury.io/js/archits-cli) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) [![npm downloads](https://img.shields.io/npm/dm/archits-cli.svg)](https://www.npmjs.com/package/archits-cli) [![GitHub stars](https://img.shields.io/github/stars/thomas-bressel/archi-ts-cli.svg)](https://github.com/thomas-bressel/archi-ts-cli/stargazers)

---

## 🚀 Quick Start

### Global Installation

```bash
### Global Installation (Recommended)

# Install globally
npm install -g archits-cli

# Check your Node.js version (must be >= 14)
node -v

# Create a symlink for easier access (if using nvm)
ln -sf /home/[your-username]/.nvm/versions/node/v[your-node-version]/lib/node_modules/archits-cli/bin/archits /home/[your-username]/.nvm/versions/node/v[your-node-version]/bin/archi

# Verify installation
archi version

```

---

## 📋 Features

* Automatic TypeScript project structure
* Entity generation (controller, service, repository, entity, route)
* Multiple architecture patterns
* Optional Express.js setup
* Pre-configured Jest testing
* Clean code templates with best practices

---

## 🛠️ Usage

### Create New Project

```bash
archi create
```

Prompts:

* Project name
* Architecture type
* Express.js integration
* And more...

### Generate Entity Files

```bash
archi generate entity Product
```

Generates:

* `Product.controller.ts`
* `Product.service.ts`
* `Product.repository.ts`
* `Product.entity.ts`
* `Product.route.ts`
* Unit tests for each layer

### Other Commands

```bash
archi version
archi --help
```

---

## ⚡ Quick Commands

| Command                        | Description                     |
| ------------------------------ | ------------------------------- |
| `archi create`                 | Create a new TypeScript project |
| `archi generate entity <Name>` | Generate entity files           |
| `archi version`                | Show version                    |
| `archi --help`                 | Show help                       |

---

## 🤔 Why ArchiTS?

* Save hours setting up TypeScript project architecture
* Keep code organized and scalable
* Integrated TypeScript best practices
* Perfect for Express.js + TypeScript
* Pre-configured Jest testing

---

## 🤝 Contributing

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

---

## 📝 License

MIT License - see [LICENSE](https://github.com/thomas-bressel/archi-ts-cli/blob/main/LICENSE)

---

## 👨‍💻 Author

**Thomas Bressel**

* GitHub: [@thomas-bressel](https://github.com/thomas-bressel)

---

## 💖 Support

[![PayPal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge\&logo=paypal\&logoColor=white)](https://www.paypal.com/donate/?hosted_button_id=UMJGHGGV4YUHE)

Made with ❤️ by Thomas Bressel
🌐 [www.archits.dev](https://www.archits.dev)
