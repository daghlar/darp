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
