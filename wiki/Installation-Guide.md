# Installation Guide

This guide covers different installation methods for DARP on Arch Linux and compatible distributions.

## Prerequisites

Before installing DARP, ensure you have:

- **Arch Linux** (or compatible distribution)
- **Go 1.21+** (for building from source)
- **WireGuard tools** (`wireguard-tools` package)
- **Root privileges** (for network configuration)

## Installation Methods

### Method 1: Pre-built Package (Recommended)

This is the easiest and fastest way to install DARP.

#### Step 1: Download

```bash
# Download the latest release
curl -L -O https://github.com/daghlar/darp/releases/latest/download/darp-latest.tar.gz
```

#### Step 2: Extract and Install

```bash
# Extract the package
tar -xzf darp-latest.tar.gz
cd darp-*

# Run the installer
sudo ./install.sh
```

#### Step 3: Verify Installation

```bash
# Check if DARP is installed
darp --version

# Check if the service is available
systemctl status darp
```

### Method 2: Build from Source

For developers or users who want the latest features.

#### Step 1: Clone Repository

```bash
git clone https://github.com/daghlar/darp.git
cd darp
```

#### Step 2: Install Dependencies

```bash
# Install Go (if not already installed)
sudo pacman -S go

# Install WireGuard tools
sudo pacman -S wireguard-tools

# Download Go dependencies
go mod download
```

#### Step 3: Build

```bash
# Make build script executable
chmod +x build.sh

# Build the project
./build.sh
```

#### Step 4: Install

```bash
# Navigate to build directory
cd build

# Extract the built package
tar -xzf darp-*.tar.gz
cd darp-*

# Install
sudo ./install.sh
```

### Method 3: Manual Installation

For advanced users who want full control.

#### Step 1: Build Binary

```bash
# Build the binary
go build -o darp ./cmd/darp

# Make it executable
chmod +x darp
```

#### Step 2: Install Manually

```bash
# Copy binary
sudo cp darp /usr/local/bin/

# Create config directory
sudo mkdir -p /etc/darp

# Create default config
sudo tee /etc/darp/config.json > /dev/null << 'EOF'
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
EOF
```

## Post-Installation

### Configuration

After installation, you may want to customize the configuration:

```bash
# Edit configuration
sudo nano /etc/darp/config.json

# Or for user-specific config
nano ~/.config/darp/config.json
```

### Service Management

```bash
# Start the service
sudo systemctl start darp

# Enable auto-start on boot
sudo systemctl enable darp

# Check status
sudo systemctl status darp
```

### First Connection

```bash
# Connect to WARP
sudo darp connect

# Check status
darp status

# Run connectivity tests
darp test connectivity
```

## Uninstallation

### Remove Package Installation

```bash
# If installed via package
cd /path/to/darp-package
sudo ./uninstall.sh
```

### Manual Removal

```bash
# Stop and disable service
sudo systemctl stop darp
sudo systemctl disable darp

# Remove files
sudo rm -f /usr/local/bin/darp
sudo rm -f /etc/systemd/system/darp.service
sudo rm -rf /etc/darp

# Reload systemd
sudo systemctl daemon-reload
```

## Troubleshooting Installation

### Common Issues

#### WireGuard Not Found

```bash
# Install WireGuard tools
sudo pacman -S wireguard-tools

# Verify installation
wg --version
```

#### Permission Denied

```bash
# Ensure you're using sudo for installation
sudo ./install.sh

# Check file permissions
ls -la darp
```

#### Go Not Found

```bash
# Install Go
sudo pacman -S go

# Verify installation
go version
```

#### Build Failures

```bash
# Clean and rebuild
make clean
make build

# Or use build script
./build.sh
```

## Next Steps

After successful installation:

1. [Configure DARP](Configuration-Reference)
2. [Learn the commands](Command-Reference)
3. [Test your setup](Troubleshooting)

## Support

If you encounter issues during installation:

- Check the [Troubleshooting](Troubleshooting) page
- Open an [issue](https://github.com/daghlar/darp/issues)
- Join our [discussions](https://github.com/daghlar/darp/discussions)
