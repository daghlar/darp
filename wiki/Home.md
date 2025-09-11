# DARP Wiki

Welcome to the DARP (Arch Linux Router Protocol) documentation wiki! This comprehensive guide will help you understand, install, configure, and troubleshoot DARP.

## 📚 Table of Contents

- [Installation Guide](Installation-Guide)
- [Configuration Reference](Configuration-Reference)
- [Command Reference](Command-Reference)
- [Troubleshooting](Troubleshooting)
- [Development Guide](Development-Guide)
- [FAQ](FAQ)
- [Contributing](Contributing)

## 🚀 Quick Start

DARP is a modular Cloudflare WARP client designed specifically for Arch Linux. It provides:

- **No API Required**: Works directly with WireGuard
- **High Performance**: Optimized for Arch Linux
- **Modular Architecture**: Clean, maintainable codebase
- **Advanced Monitoring**: Real-time network diagnostics

### Installation

```bash
# Clone and build
git clone https://github.com/daghlar/darp.git
cd darp
./build.sh

# Install
cd build
tar -xzf darp-*.tar.gz
cd darp-*
sudo ./install.sh
```

### Basic Usage

```bash
# Connect to WARP
sudo darp connect

# Check status
darp status

# Run tests
darp test connectivity
```

## 📖 Documentation Pages

### [Installation Guide](Installation-Guide)
Complete step-by-step installation instructions for different scenarios.

### [Configuration Reference](Configuration-Reference)
Detailed configuration options and examples.

### [Command Reference](Command-Reference)
Complete list of all available commands and options.

### [Troubleshooting](Troubleshooting)
Common issues and their solutions.

### [Development Guide](Development-Guide)
How to contribute to DARP development.

### [FAQ](FAQ)
Frequently asked questions and answers.

## 🤝 Contributing

We welcome contributions! See our [Contributing Guide](Contributing) for details.

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/daghlar/darp/issues)
- **Discussions**: [GitHub Discussions](https://github.com/daghlar/darp/discussions)

---

**Made with ❤️ for the Arch Linux community**
