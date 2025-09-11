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
