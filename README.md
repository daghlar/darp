# DARP - Cloudflare WARP Client for Arch Linux

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Arch Linux](https://img.shields.io/badge/Arch%20Linux-Supported-blue.svg)](https://archlinux.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

**DARP** (Arch Linux Router Protocol) is a lightweight, modular Cloudflare WARP client specifically designed for Arch Linux. It provides seamless integration with Cloudflare's WARP service using WireGuard technology without requiring API keys.

## ✨ Features

- 🚀 **Lightweight**: Only 6 Go files, ~1,000 lines of clean code
- 🔧 **Modular Architecture**: Clean, maintainable codebase with separate modules
- 🛡️ **Secure**: Uses WireGuard for encrypted tunneling
- 📊 **Network Diagnostics**: Real-time connectivity testing and performance monitoring
- 🎯 **Easy Configuration**: JSON-based configuration with sensible defaults
- 📱 **CLI Interface**: Intuitive command-line interface with personalized welcome messages
- 🐧 **Arch Linux Optimized**: Specifically tuned for Arch Linux networking stack
- 🔑 **No API Required**: Works without Cloudflare API keys
- ⚡ **Quick Setup**: Get running in minutes, not hours

## 🏗️ Architecture

```
darp/
├── cmd/darp/           # Main application entry point
├── pkg/
│   ├── cli/            # Command-line interface
│   ├── config/         # Configuration management
│   ├── network/        # Network diagnostics and optimization
│   └── warp/           # Cloudflare WARP integration
├── internal/           # Internal utilities (reserved for future use)
├── build/              # Build output directory
├── build.sh            # Build automation script
├── Makefile            # Development commands
├── go.mod              # Go module definition
├── LICENSE             # MIT License
└── README.md           # This file
```

## 🚀 Quick Start

### Prerequisites

- Arch Linux (or compatible distribution)
- Go 1.21+ (for building from source)
- WireGuard tools (`sudo pacman -S wireguard-tools`)
- Root privileges (for network configuration)

### Installation

#### Option 1: Pre-built Package (Recommended)

```bash
# Download the latest release
curl -L -O https://github.com/daghlar/darp/releases/latest/download/darp-latest.tar.gz

# Extract and install
tar -xzf darp-latest.tar.gz
cd darp-*
sudo ./install.sh
```

#### Option 2: Build from Source

```bash
# Clone the repository
git clone https://github.com/daghlar/darp.git
cd darp

# Make build script executable
chmod +x build.sh

# Build the project
./build.sh

# Install the built package
cd build
tar -xzf darp-*.tar.gz
cd darp-*
sudo ./install.sh
```

### Configuration

1. **Edit the configuration file:**
   ```bash
   sudo nano /etc/darp/config.json
   ```

2. **No API credentials needed!** DARP works with Cloudflare WARP without any API keys:
   ```json
   {
     "cloudflare": {
       "warp_endpoint": "engage.cloudflareclient.com:2408"
     }
   }
   ```

3. **Start the service:**
   ```bash
   sudo systemctl start darp
   ```

## 📖 Usage

### Command Line Interface

```bash
# Show welcome message and available commands
darp

# Connect to WARP
sudo darp connect

# Disconnect from WARP
sudo darp disconnect

# Check connection status
darp status

# Show status in JSON format
darp status --format json

# Run network tests
darp test connectivity
darp test latency
darp test dns

# Optimize network settings
sudo darp optimize

# Show configuration
darp config show

# Set configuration values
darp config set cloudflare.warp_endpoint "engage.cloudflareclient.com:2408"
```

### Service Management

```bash
# Start the service
sudo systemctl start darp

# Stop the service
sudo systemctl stop darp

# Enable auto-start on boot
sudo systemctl enable darp

# Check service status
sudo systemctl status darp

# View logs
sudo journalctl -u darp -f
```

## ⚙️ Configuration

### Configuration File Location

- **System-wide**: `/etc/darp/config.json`
- **User-specific**: `~/.config/darp/config.json`

### Configuration Options

```json
{
  "cloudflare": {
    "warp_endpoint": "engage.cloudflareclient.com:2408"
  },
  "network": {
    "interface": "warp0",
    "dns": ["1.1.1.1", "1.0.0.1"],
    "mtu": 1280,
    "timeout": 30
  },
  "logging": {
    "level": "info",
    "format": "json", 
    "output": "stdout"
  }
}
```

## 🔧 Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/daghlar/darp.git
cd darp

# Install dependencies
go mod download

# Run tests
go test ./...

# Build the project
./build.sh

# Run the application
sudo ./build/darp connect
```

### Project Structure

- **`cmd/darp/`**: Main application entry point
- **`pkg/cli/`**: Command-line interface implementation
- **`pkg/config/`**: Configuration management and validation
- **`pkg/network/`**: Network diagnostics and optimization
- **`pkg/warp/`**: Cloudflare WARP API integration and WireGuard management
- **`internal/`**: Internal utilities and helpers

### Adding New Features

1. Create a new module in the appropriate `pkg/` directory
2. Add CLI commands in `pkg/cli/cli.go`
3. Update configuration schema in `pkg/config/config.go`
4. Add tests for your new functionality
5. Update documentation

## 🧪 Testing

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run specific test package
go test ./pkg/warp/...
```

## 📊 Performance

DARP is optimized for high performance on Arch Linux:

- **BBR Congestion Control**: Automatic TCP optimization via `sysctl`
- **Buffer Tuning**: Optimized network buffer sizes (128MB)
- **WireGuard Integration**: Direct WireGuard config generation
- **Memory Efficient**: Minimal memory footprint (~10MB)
- **Fast Startup**: Quick connection establishment
- **No API Overhead**: Direct WireGuard connection without API calls
- **Network Diagnostics**: Built-in latency testing and connectivity checks

## 🛠️ Troubleshooting

### Common Issues

1. **WireGuard not found**
   ```bash
   sudo pacman -S wireguard-tools
   ```

2. **Permission denied**
   ```bash
   sudo darp connect
   ```

3. **Configuration errors**
   ```bash
   darp config show
   # Check your configuration file
   ```

4. **Network connectivity issues**
   ```bash
   darp test connectivity
   darp test dns
   ```

### Debug Mode

```bash
# Enable verbose logging
darp --verbose connect

# Check system logs
sudo journalctl -u darp -f
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](https://github.com/daghlar/darp/blob/main/wiki/Contributing.md) for details.

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Cloudflare](https://cloudflare.com) for the WARP service
- [WireGuard](https://wireguard.com) for the VPN technology
- [Cobra](https://github.com/spf13/cobra) for the CLI framework
- [Arch Linux](https://archlinux.org) community

## 📞 Support

- **Issues**: [GitHub Issues](https://github.com/daghlar/darp/issues)
- **Discussions**: [GitHub Discussions](https://github.com/daghlar/darp/discussions)
- **Documentation**: [Wiki](https://github.com/daghlar/darp/wiki)

---

**Made with ❤️ for the Arch Linux community**