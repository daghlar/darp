# Command Reference

Complete reference for all DARP commands and options.

## Basic Usage

```bash
darp [command] [options]
```

## Global Options

| Option | Description |
|--------|-------------|
| `--config` | Path to configuration file |
| `--verbose` | Enable verbose logging |
| `--version` | Show version information |
| `--help` | Show help information |

## Commands

### connect

Establishes a connection to Cloudflare WARP.

```bash
sudo darp connect
```

**Description**: Connects to Cloudflare WARP using WireGuard. This command requires root privileges.

**Examples**:
```bash
# Basic connection
sudo darp connect

# With verbose output
sudo darp connect --verbose

# With custom config
sudo darp connect --config /path/to/config.json
```

### disconnect

Terminates the current WARP connection.

```bash
sudo darp disconnect
```

**Description**: Disconnects from Cloudflare WARP and stops the WireGuard interface.

**Examples**:
```bash
# Basic disconnect
sudo darp disconnect

# With verbose output
sudo darp disconnect --verbose
```

### status

Shows connection status and information.

```bash
darp status [options]
```

**Options**:
| Option | Description | Default |
|--------|-------------|---------|
| `--format` | Output format (table, json) | `table` |

**Examples**:
```bash
# Show status in table format
darp status

# Show status in JSON format
darp status --format json

# Show status with verbose output
darp status --verbose
```

**Output Example**:
```
┌─────────────────────────────────────────┐
│              DARP Status                │
├─────────────────────────────────────────┤
│ Status: ✅ Connected                    │
│ Interface: warp0                       │
│ IP Address: 10.0.0.1                  │
│ DNS Servers: 1.1.1.1, 1.0.0.1        │
│ Uptime: 2h 15m                        │
│ Data Sent: 1.2 GB                     │
│ Data Received: 3.4 GB                │
└─────────────────────────────────────────┘
```

### config

Manages configuration settings.

#### config show

Displays current configuration.

```bash
darp config show
```

**Examples**:
```bash
# Show current configuration
darp config show

# Show with verbose output
darp config show --verbose
```

#### config set

Sets configuration values.

```bash
darp config set <key> <value>
```

**Examples**:
```bash
# Set WARP endpoint
darp config set cloudflare.warp_endpoint "engage.cloudflareclient.com:2408"

# Set DNS servers
darp config set network.dns "[\"8.8.8.8\", \"8.8.4.4\"]"

# Set log level
darp config set logging.level "debug"
```

### test

Runs various network tests.

#### test connectivity

Tests basic network connectivity.

```bash
darp test connectivity
```

**Description**: Tests DNS resolution, internet connectivity, and WARP API availability.

**Examples**:
```bash
# Run connectivity tests
darp test connectivity

# Run with verbose output
darp test connectivity --verbose
```

**Output Example**:
```
🔍 Testing network connectivity...
  ✅ PASS DNS Resolution
  ✅ PASS Internet Connectivity
  ✅ PASS Cloudflare WARP API
  ✅ PASS WireGuard Interface

🎉 All connectivity tests passed!
```

#### test latency

Tests latency to various endpoints.

```bash
darp test latency
```

**Description**: Measures latency to Cloudflare and Google DNS servers.

**Examples**:
```bash
# Test latency
darp test latency

# Test with verbose output
darp test latency --verbose
```

**Output Example**:
```
⏱️  Testing latency to various endpoints...
  Cloudflare DNS (1.1.1.1): 12ms
  Cloudflare DNS (1.0.0.1): 15ms
  Google DNS (8.8.8.8): 25ms
  Google DNS (8.8.4.4): 28ms
```

#### test dns

Tests DNS resolution.

```bash
darp test dns
```

**Description**: Tests DNS resolution for common domains.

**Examples**:
```bash
# Test DNS resolution
darp test dns

# Test with verbose output
darp test dns --verbose
```

**Output Example**:
```
🌐 Testing DNS resolution...
  Resolving cloudflare.com... ✅ OK
  Resolving google.com... ✅ OK
  Resolving github.com... ✅ OK
  Resolving archlinux.org... ✅ OK
```

### optimize

Optimizes network settings for better performance.

```bash
sudo darp optimize
```

**Description**: Applies network optimizations including TCP congestion control and buffer tuning.

**Examples**:
```bash
# Optimize network settings
sudo darp optimize

# Optimize with verbose output
sudo darp optimize --verbose
```

**Output Example**:
```
⚡ Optimizing network settings...
  Setting TCP congestion control to BBR... ✅ Done
  Increasing network buffer sizes... ✅ Done
  Optimizing WireGuard MTU... ✅ Done
  Configuring DNS caching... ✅ Done

🎯 Network optimization completed!
```

## Service Management

DARP can also be managed as a systemd service:

### Start Service

```bash
sudo systemctl start darp
```

### Stop Service

```bash
sudo systemctl stop darp
```

### Enable Service

```bash
sudo systemctl enable darp
```

### Disable Service

```bash
sudo systemctl disable darp
```

### Check Service Status

```bash
sudo systemctl status darp
```

### View Service Logs

```bash
# View recent logs
sudo journalctl -u darp

# Follow logs in real-time
sudo journalctl -u darp -f

# View logs with timestamps
sudo journalctl -u darp --since "1 hour ago"
```

## Command Examples

### Complete Workflow

```bash
# 1. Check current status
darp status

# 2. Run connectivity tests
darp test connectivity

# 3. Connect to WARP
sudo darp connect

# 4. Check status again
darp status

# 5. Test latency
darp test latency

# 6. Optimize network
sudo darp optimize

# 7. Disconnect when done
sudo darp disconnect
```

### Troubleshooting Workflow

```bash
# 1. Check status
darp status

# 2. Test connectivity
darp test connectivity

# 3. Test DNS
darp test dns

# 4. Check configuration
darp config show

# 5. Connect with verbose output
sudo darp connect --verbose
```

### Configuration Management

```bash
# 1. View current config
darp config show

# 2. Set custom DNS
darp config set network.dns "[\"8.8.8.8\", \"8.8.4.4\"]"

# 3. Set log level
darp config set logging.level "debug"

# 4. Verify changes
darp config show
```

## Exit Codes

DARP uses standard Unix exit codes:

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | General error |
| `2` | Configuration error |
| `3` | Network error |
| `4` | Permission error |

## Environment Variables

| Variable | Description |
|----------|-------------|
| `DARP_CONFIG` | Path to configuration file |
| `DARP_LOG_LEVEL` | Log level (debug, info, warn, error) |
| `DARP_LOG_FORMAT` | Log format (json, text) |

## Next Steps

- [Troubleshooting](Troubleshooting) - Common issues and solutions
- [Configuration Reference](Configuration-Reference) - Detailed configuration options
