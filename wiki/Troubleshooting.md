# Troubleshooting

This guide helps you diagnose and resolve common issues with DARP.

## Quick Diagnostics

### Check System Status

```bash
# Check if DARP is running
darp status

# Check systemd service
sudo systemctl status darp

# Check WireGuard installation
wg --version
```

### Run Diagnostic Tests

```bash
# Test connectivity
darp test connectivity

# Test DNS resolution
darp test dns

# Test latency
darp test latency
```

## Common Issues

### Connection Issues

#### "WireGuard tools not found"

**Problem**: DARP cannot find WireGuard tools.

**Solution**:
```bash
# Install WireGuard tools
sudo pacman -S wireguard-tools

# Verify installation
wg --version
wg-quick --version
```

#### "Permission denied"

**Problem**: DARP requires root privileges for network operations.

**Solution**:
```bash
# Use sudo for network operations
sudo darp connect
sudo darp disconnect
sudo darp optimize
```

#### "Failed to start WireGuard interface"

**Problem**: WireGuard interface cannot be started.

**Possible Causes**:
- Interface already exists
- Permission issues
- Network configuration conflicts

**Solutions**:
```bash
# Check existing interfaces
ip link show

# Remove existing interface if needed
sudo wg-quick down darp

# Try connecting again
sudo darp connect
```

#### "DNS resolution failed"

**Problem**: Cannot resolve domain names.

**Solutions**:
```bash
# Test DNS manually
nslookup cloudflare.com

# Check DNS configuration
darp config show

# Test with different DNS servers
darp config set network.dns "[\"8.8.8.8\", \"8.8.4.4\"]"
```

### Configuration Issues

#### "Configuration validation failed"

**Problem**: Configuration file has errors.

**Solutions**:
```bash
# Check configuration syntax
darp config show

# Validate JSON syntax
cat /etc/darp/config.json | jq .

# Reset to default configuration
sudo rm /etc/darp/config.json
sudo darp connect  # This will create a new default config
```

#### "WARP endpoint must be configured"

**Problem**: Missing required configuration.

**Solution**:
```bash
# Set the WARP endpoint
darp config set cloudflare.warp_endpoint "engage.cloudflareclient.com:2408"
```

### Performance Issues

#### Slow Connection

**Problem**: Connection is slow or unstable.

**Solutions**:
```bash
# Optimize network settings
sudo darp optimize

# Test latency
darp test latency

# Check MTU settings
darp config show | grep mtu
```

#### High CPU Usage

**Problem**: DARP is using too much CPU.

**Solutions**:
```bash
# Check log level
darp config set logging.level "warn"

# Check for errors in logs
sudo journalctl -u darp --since "1 hour ago"
```

### Service Issues

#### Service won't start

**Problem**: systemd service fails to start.

**Solutions**:
```bash
# Check service status
sudo systemctl status darp

# Check service logs
sudo journalctl -u darp

# Restart service
sudo systemctl restart darp
```

#### Service stops unexpectedly

**Problem**: Service stops running after some time.

**Solutions**:
```bash
# Check service logs
sudo journalctl -u darp --since "1 hour ago"

# Check system resources
free -h
df -h

# Restart service
sudo systemctl restart darp
```

## Advanced Troubleshooting

### Debug Mode

Enable verbose logging for detailed information:

```bash
# Run with debug output
sudo darp connect --verbose

# Set debug log level
darp config set logging.level "debug"

# Check logs
sudo journalctl -u darp -f
```

### Network Diagnostics

#### Check Network Interfaces

```bash
# List all interfaces
ip link show

# Check WireGuard interface
ip link show warp0

# Check routing table
ip route show
```

#### Test Network Connectivity

```bash
# Test basic connectivity
ping -c 4 1.1.1.1

# Test DNS resolution
nslookup cloudflare.com

# Test specific port
telnet engage.cloudflareclient.com 2408
```

#### Check Firewall

```bash
# Check iptables rules
sudo iptables -L

# Check if firewall is blocking
sudo ufw status
```

### Log Analysis

#### View Logs

```bash
# View recent logs
sudo journalctl -u darp --since "1 hour ago"

# Follow logs in real-time
sudo journalctl -u darp -f

# View logs with timestamps
sudo journalctl -u darp --since "2024-01-01" --until "2024-01-02"
```

#### Common Log Messages

| Message | Meaning | Solution |
|---------|---------|----------|
| "WireGuard tools not found" | WireGuard not installed | Install wireguard-tools |
| "Permission denied" | Need root privileges | Use sudo |
| "Interface already exists" | Interface conflict | Remove existing interface |
| "DNS resolution failed" | DNS issues | Check DNS configuration |
| "Connection timeout" | Network timeout | Check network connectivity |

## System Requirements

### Minimum Requirements

- **OS**: Arch Linux or compatible
- **RAM**: 64 MB
- **Disk**: 10 MB
- **Network**: Internet connection

### Recommended Requirements

- **OS**: Arch Linux
- **RAM**: 128 MB
- **Disk**: 50 MB
- **Network**: Stable internet connection

## Performance Tuning

### Network Optimization

```bash
# Apply optimizations
sudo darp optimize

# Check current settings
sysctl net.core.default_qdisc
sysctl net.ipv4.tcp_congestion_control
```

### Resource Monitoring

```bash
# Monitor CPU usage
top -p $(pgrep darp)

# Monitor memory usage
ps aux | grep darp

# Monitor network usage
iftop -i warp0
```

## Getting Help

### Before Asking for Help

1. **Check this guide** for your specific issue
2. **Run diagnostics** using the commands above
3. **Check logs** for error messages
4. **Try basic troubleshooting** steps

### When Reporting Issues

Include the following information:

1. **System information**:
   ```bash
   uname -a
   pacman -Q wireguard-tools
   darp --version
   ```

2. **Configuration**:
   ```bash
   darp config show
   ```

3. **Logs**:
   ```bash
   sudo journalctl -u darp --since "1 hour ago"
   ```

4. **Error messages**: Exact error messages you're seeing

### Support Channels

- **GitHub Issues**: [Report bugs](https://github.com/daghlar/darp/issues)
- **GitHub Discussions**: [Ask questions](https://github.com/daghlar/darp/discussions)
- **Documentation**: Check other wiki pages

## Prevention

### Regular Maintenance

```bash
# Check status regularly
darp status

# Run tests periodically
darp test connectivity

# Update DARP when new versions are available
git pull origin main
./build.sh
```

### Best Practices

1. **Keep WireGuard updated**: `sudo pacman -Syu wireguard-tools`
2. **Monitor logs**: Check logs regularly for issues
3. **Backup configuration**: Keep backups of working configurations
4. **Test changes**: Test configuration changes before applying to production

## Next Steps

- [Command Reference](Command-Reference) - Learn all available commands
- [Configuration Reference](Configuration-Reference) - Detailed configuration options
- [Development Guide](Development-Guide) - Contribute to DARP development
