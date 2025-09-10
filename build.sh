#!/bin/bash

# DARP Build Script for Arch Linux
# Cloudflare WARP client build automation

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Build configuration
APP_NAME="darp"
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(date -u '+%Y-%m-%d_%H:%M:%S')
GO_VERSION=$(go version | awk '{print $3}')

# Output directory
BUILD_DIR="build"
BINARY_NAME="${APP_NAME}-${VERSION}"

echo -e "${BLUE}🚀 Building DARP - Cloudflare WARP Client${NC}"
echo -e "${BLUE}==========================================${NC}"
echo -e "Version: ${GREEN}${VERSION}${NC}"
echo -e "Build Time: ${GREEN}${BUILD_TIME}${NC}"
echo -e "Go Version: ${GREEN}${GO_VERSION}${NC}"
echo ""

# Clean previous builds
echo -e "${YELLOW}🧹 Cleaning previous builds...${NC}"
rm -rf ${BUILD_DIR}
mkdir -p ${BUILD_DIR}

# Check dependencies
echo -e "${YELLOW}📦 Checking dependencies...${NC}"
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go first.${NC}"
    echo -e "   sudo pacman -S go"
    exit 1
fi

if ! command -v git &> /dev/null; then
    echo -e "${RED}❌ Git is not installed. Please install Git first.${NC}"
    echo -e "   sudo pacman -S git"
    exit 1
fi

# Download dependencies
echo -e "${YELLOW}📥 Downloading dependencies...${NC}"
go mod download
go mod tidy

# Run tests
echo -e "${YELLOW}🧪 Running tests...${NC}"
go test ./... -v

# Build for different architectures
echo -e "${YELLOW}🔨 Building binaries...${NC}"

# AMD64 (x86_64)
echo -e "  Building for ${GREEN}linux/amd64${NC}..."
GOOS=linux GOARCH=amd64 go build \
    -ldflags "-X main.version=${VERSION} -X main.build=${BUILD_TIME}" \
    -o ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 \
    ./cmd/darp

# ARM64
echo -e "  Building for ${GREEN}linux/arm64${NC}..."
GOOS=linux GOARCH=arm64 go build \
    -ldflags "-X main.version=${VERSION} -X main.build=${BUILD_TIME}" \
    -o ${BUILD_DIR}/${BINARY_NAME}-linux-arm64 \
    ./cmd/darp

# Create symlinks for easy access
ln -sf ${BINARY_NAME}-linux-amd64 ${BUILD_DIR}/darp
ln -sf ${BINARY_NAME}-linux-amd64 ${BUILD_DIR}/darp-latest

# Create installation package
echo -e "${YELLOW}📦 Creating installation package...${NC}"
PACKAGE_DIR="${BUILD_DIR}/darp-${VERSION}"
mkdir -p ${PACKAGE_DIR}

# Copy binary
cp ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 ${PACKAGE_DIR}/darp
chmod +x ${PACKAGE_DIR}/darp

# Copy configuration template
mkdir -p ${PACKAGE_DIR}/config
cat > ${PACKAGE_DIR}/config/config.json << EOF
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

# Copy documentation
cp README.md ${PACKAGE_DIR}/
cp LICENSE ${PACKAGE_DIR}/ 2>/dev/null || echo "No LICENSE file found"

# Create systemd service file
cat > ${PACKAGE_DIR}/darp.service << EOF
[Unit]
Description=DARP Cloudflare WARP Client
After=network.target

[Service]
Type=simple
User=root
ExecStart=/usr/local/bin/darp connect
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
EOF

# Create install script
cat > ${PACKAGE_DIR}/install.sh << 'EOF'
#!/bin/bash

# DARP Installation Script for Arch Linux

set -e

echo "🚀 Installing DARP - Cloudflare WARP Client"

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "❌ Please run as root (use sudo)"
    exit 1
fi

# Install WireGuard if not present
if ! command -v wg &> /dev/null; then
    echo "📦 Installing WireGuard..."
    pacman -S --noconfirm wireguard-tools
fi

# Copy binary
echo "📁 Installing binary..."
cp darp /usr/local/bin/
chmod +x /usr/local/bin/darp

# Create config directory
echo "📁 Creating configuration directory..."
mkdir -p /etc/darp
cp config/config.json /etc/darp/

# Install systemd service
echo "🔧 Installing systemd service..."
cp darp.service /etc/systemd/system/
systemctl daemon-reload
systemctl enable darp

echo "✅ Installation completed!"
echo ""
echo "Next steps:"
echo "1. Edit /etc/darp/config.json with your Cloudflare credentials"
echo "2. Start the service: sudo systemctl start darp"
echo "3. Check status: sudo systemctl status darp"
echo "4. Or run manually: sudo darp connect"
EOF

chmod +x ${PACKAGE_DIR}/install.sh

# Create uninstall script
cat > ${PACKAGE_DIR}/uninstall.sh << 'EOF'
#!/bin/bash

# DARP Uninstallation Script

set -e

echo "🗑️  Uninstalling DARP - Cloudflare WARP Client"

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    echo "❌ Please run as root (use sudo)"
    exit 1
fi

# Stop and disable service
echo "🛑 Stopping service..."
systemctl stop darp 2>/dev/null || true
systemctl disable darp 2>/dev/null || true

# Remove files
echo "🗑️  Removing files..."
rm -f /usr/local/bin/darp
rm -f /etc/systemd/system/darp.service
rm -rf /etc/darp

# Reload systemd
systemctl daemon-reload

echo "✅ Uninstallation completed!"
EOF

chmod +x ${PACKAGE_DIR}/uninstall.sh

# Create tarball
echo -e "${YELLOW}📦 Creating distribution package...${NC}"
cd ${BUILD_DIR}
tar -czf ${BINARY_NAME}.tar.gz darp-${VERSION}/
cd ..

# Show build results
echo ""
echo -e "${GREEN}✅ Build completed successfully!${NC}"
echo -e "${GREEN}================================${NC}"
echo -e "Build directory: ${BLUE}${BUILD_DIR}/${NC}"
echo -e "Package: ${BLUE}${BUILD_DIR}/${BINARY_NAME}.tar.gz${NC}"
echo -e "Binary size: ${BLUE}$(du -h ${BUILD_DIR}/${BINARY_NAME}-linux-amd64 | cut -f1)${NC}"
echo ""
echo -e "${YELLOW}📋 Installation instructions:${NC}"
echo -e "1. Extract: ${BLUE}tar -xzf ${BUILD_DIR}/${BINARY_NAME}.tar.gz${NC}"
echo -e "2. Install: ${BLUE}cd darp-${VERSION} && sudo ./install.sh${NC}"
echo -e "3. Configure: ${BLUE}sudo nano /etc/darp/config.json${NC}"
echo -e "4. Start: ${BLUE}sudo darp connect${NC}"
echo ""
echo -e "${GREEN}🎉 Ready to deploy!${NC}"
