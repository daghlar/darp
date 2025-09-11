# Frequently Asked Questions (FAQ)

Common questions and answers about DARP.

## General Questions

### What is DARP?

DARP (Arch Linux Router Protocol) is a lightweight, modular Cloudflare WARP client specifically designed for Arch Linux. It provides seamless integration with Cloudflare's WARP service using WireGuard technology without requiring API keys.

### Why use DARP instead of the official Cloudflare WARP client?

- **No API required**: Works without Cloudflare API keys
- **Lightweight**: Only 6 Go files, ~1,000 lines of code
- **Arch Linux optimized**: Specifically tuned for Arch Linux
- **Modular architecture**: Clean, maintainable codebase
- **Network diagnostics**: Real-time connectivity testing and performance monitoring
- **Personalized CLI**: Welcome messages with user recognition
- **Open source**: Full control over the code

### Is DARP free to use?

Yes, DARP is completely free and open source under the MIT License.

### What operating systems are supported?

DARP is primarily designed for Arch Linux, but it should work on any Linux distribution that supports WireGuard.

## Installation Questions

### Do I need to install WireGuard separately?

Yes, DARP requires WireGuard tools to be installed:

```bash
sudo pacman -S wireguard-tools
```

### Can I install DARP without root privileges?

No, DARP requires root privileges for network configuration. However, you can run most commands without root:

```bash
# These work without root
darp status
darp test connectivity
darp config show

# These require root
sudo darp connect
sudo darp disconnect
sudo darp optimize
```

### How do I update DARP?

```bash
# Pull latest changes
git pull origin main

# Rebuild
./build.sh

# Reinstall
cd build
tar -xzf darp-*.tar.gz
cd darp-*
sudo ./install.sh
```

## Configuration Questions

### Do I need Cloudflare API keys?

No! DARP works without any API keys. It connects directly to Cloudflare's public WireGuard endpoints.

### How do I change DNS servers?

```bash
# Set custom DNS servers
darp config set network.dns "[\"8.8.8.8\", \"8.8.4.4\"]"

# Or edit config file directly
sudo nano /etc/darp/config.json
```

### Can I use DARP with other VPNs?

DARP is specifically designed for Cloudflare WARP. It cannot be used with other VPN services.

### How do I configure logging?

```bash
# Set log level
darp config set logging.level "debug"

# Set log format
darp config set logging.format "text"

# Set log output
darp config set logging.output "/var/log/darp.log"
```

## Usage Questions

### How do I get started with DARP?

```bash
# Show welcome message and available commands
darp
```

### How do I connect to WARP?

```bash
sudo darp connect
```

### How do I disconnect from WARP?

```bash
sudo darp disconnect
```

### How do I check if I'm connected?

```bash
darp status
```

### Can I run DARP as a service?

Yes, DARP can be run as a systemd service:

```bash
# Start service
sudo systemctl start darp

# Enable auto-start
sudo systemctl enable darp

# Check status
sudo systemctl status darp
```

## Troubleshooting Questions

### Why can't I connect?

Common causes and solutions:

1. **WireGuard not installed**:
   ```bash
   sudo pacman -S wireguard-tools
   ```

2. **Permission denied**:
   ```bash
   sudo darp connect
   ```

3. **Network issues**:
   ```bash
   darp test connectivity
   ```

### Why is my connection slow?

Try these optimizations:

```bash
# Optimize network settings
sudo darp optimize

# Test latency
darp test latency

# Check MTU settings
darp config show | grep mtu
```

### Why does DARP stop working?

Check these:

1. **Service status**:
   ```bash
   sudo systemctl status darp
   ```

2. **Logs**:
   ```bash
   sudo journalctl -u darp --since "1 hour ago"
   ```

3. **Network connectivity**:
   ```bash
   darp test connectivity
   ```

### How do I debug issues?

Enable verbose logging:

```bash
# Run with debug output
sudo darp connect --verbose

# Set debug log level
darp config set logging.level "debug"

# Check logs
sudo journalctl -u darp -f
```

## Technical Questions

### How does DARP work?

DARP uses WireGuard to create a secure tunnel to Cloudflare's WARP servers. It:

1. Generates WireGuard keys
2. Connects to Cloudflare's public endpoints
3. Routes traffic through the tunnel
4. Provides DNS resolution through Cloudflare's DNS

### What ports does DARP use?

DARP uses:
- **UDP 2408**: Cloudflare WARP endpoint
- **UDP 51820**: WireGuard (if using custom port)

### Is my traffic encrypted?

Yes, all traffic is encrypted using WireGuard's encryption before being sent to Cloudflare's servers.

### Does DARP log my traffic?

No, DARP does not log your traffic. It only logs connection status and errors.

### Can I see what DARP is doing?

Yes, you can monitor DARP:

```bash
# Check status
darp status

# Monitor logs
sudo journalctl -u darp -f

# Check network interfaces
ip link show warp0
```

## Performance Questions

### How much bandwidth does DARP use?

DARP adds minimal overhead. The actual bandwidth usage depends on your internet traffic.

### Does DARP affect gaming?

DARP may add some latency due to the VPN tunnel. For gaming, you might want to:

1. Test latency: `darp test latency`
2. Optimize settings: `sudo darp optimize`
3. Consider disconnecting for gaming if latency is too high

### Can I use DARP on a server?

Yes, DARP can be used on servers. Make sure to:

1. Configure it as a service
2. Set appropriate logging levels
3. Monitor resource usage

## Security Questions

### Is DARP secure?

Yes, DARP uses WireGuard, which is considered very secure. It:

- Uses modern cryptography
- Has a small attack surface
- Is regularly audited

### Does DARP collect data?

No, DARP does not collect any personal data. It only connects to Cloudflare's WARP service.

### Can I audit DARP's code?

Yes, DARP is open source. You can:

1. View the source code on GitHub
2. Build from source
3. Modify the code as needed

## Development Questions

### How do I contribute to DARP?

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

### Can I modify DARP?

Yes, DARP is open source under the MIT License. You can modify it as needed.

### How do I report bugs?

1. Check the [Troubleshooting](Troubleshooting) page
2. Search existing [issues](https://github.com/daghlar/darp/issues)
3. Create a new issue with detailed information

### How do I request features?

1. Check existing [discussions](https://github.com/daghlar/darp/discussions)
2. Create a new discussion
3. Provide detailed use case and requirements

## Still Have Questions?

If you can't find the answer to your question:

1. **Check the documentation**: Browse other wiki pages
2. **Search issues**: Look for similar problems
3. **Ask in discussions**: Create a new discussion
4. **Report bugs**: Open an issue if you found a bug

### Useful Links

- [Installation Guide](Installation-Guide)
- [Configuration Reference](Configuration-Reference)
- [Command Reference](Command-Reference)
- [Troubleshooting](Troubleshooting)
- [GitHub Issues](https://github.com/daghlar/darp/issues)
- [GitHub Discussions](https://github.com/daghlar/darp/discussions)
