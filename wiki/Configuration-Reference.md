# Configuration Reference

This page provides detailed information about DARP configuration options.

## Configuration Files

DARP supports multiple configuration file locations:

- **System-wide**: `/etc/darp/config.json` (requires root)
- **User-specific**: `~/.config/darp/config.json` (user home directory)

User-specific configuration takes precedence over system-wide configuration.

## Configuration Schema

### Complete Configuration Example

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

## Configuration Sections

### Cloudflare Section

Controls Cloudflare WARP connection settings.

```json
{
  "cloudflare": {
    "warp_endpoint": "engage.cloudflareclient.com:2408"
  }
}
```

#### Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `warp_endpoint` | string | `engage.cloudflareclient.com:2408` | Cloudflare WARP server endpoint |

**Note**: No API keys are required! DARP works directly with Cloudflare's public WireGuard endpoints.

### Network Section

Controls network interface and routing settings.

```json
{
  "network": {
    "interface": "warp0",
    "dns": ["1.1.1.1", "1.0.0.1"],
    "mtu": 1280,
    "timeout": 30
  }
}
```

#### Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `interface` | string | `warp0` | WireGuard interface name |
| `dns` | array | `["1.1.1.1", "1.0.0.1"]` | DNS servers to use |
| `mtu` | integer | `1280` | Maximum Transmission Unit |
| `timeout` | integer | `30` | Connection timeout in seconds |

#### DNS Servers

DARP uses Cloudflare's DNS servers by default:

- **Primary**: `1.1.1.1` - Cloudflare's main DNS
- **Secondary**: `1.0.0.1` - Cloudflare's backup DNS

You can change these to any DNS servers you prefer:

```json
{
  "network": {
    "dns": ["8.8.8.8", "8.8.4.4"]
  }
}
```

#### MTU Settings

The MTU (Maximum Transmission Unit) determines the maximum packet size:

- **Default**: `1280` - Optimized for most networks
- **Higher values**: May improve performance but can cause issues on some networks
- **Lower values**: More compatible but may reduce performance

### Logging Section

Controls logging behavior and output.

```json
{
  "logging": {
    "level": "info",
    "format": "json",
    "output": "stdout"
  }
}
```

#### Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `level` | string | `info` | Log level (debug, info, warn, error) |
| `format` | string | `json` | Log format (json, text) |
| `output` | string | `stdout` | Output destination (stdout, stderr, file) |

#### Log Levels

- **debug**: Detailed information for debugging
- **info**: General information about operations
- **warn**: Warning messages
- **error**: Error messages only

#### Log Formats

- **json**: Structured JSON format (machine-readable)
- **text**: Human-readable text format

## Configuration Management

### Viewing Configuration

```bash
# Show current configuration
darp config show

# Show in JSON format
darp config show --format json
```

### Setting Configuration Values

```bash
# Set a configuration value
darp config set cloudflare.warp_endpoint "engage.cloudflareclient.com:2408"

# Set DNS servers
darp config set network.dns "[\"8.8.8.8\", \"8.8.4.4\"]"
```

### Configuration Validation

DARP automatically validates configuration on startup:

```bash
# Test configuration
darp config show

# If there are errors, they will be displayed
```

## Advanced Configuration

### Custom WireGuard Settings

While DARP manages WireGuard configuration automatically, you can influence some settings:

```json
{
  "network": {
    "interface": "custom-warp",
    "mtu": 1420,
    "timeout": 60
  }
}
```

### Multiple DNS Servers

You can specify multiple DNS servers for redundancy:

```json
{
  "network": {
    "dns": [
      "1.1.1.1",
      "1.0.0.1",
      "8.8.8.8",
      "8.8.4.4"
    ]
  }
}
```

### Logging to File

To log to a file instead of stdout:

```json
{
  "logging": {
    "level": "debug",
    "format": "text",
    "output": "/var/log/darp.log"
  }
}
```

## Configuration Examples

### Minimal Configuration

```json
{
  "cloudflare": {
    "warp_endpoint": "engage.cloudflareclient.com:2408"
  }
}
```

### High Performance Configuration

```json
{
  "cloudflare": {
    "warp_endpoint": "engage.cloudflareclient.com:2408"
  },
  "network": {
    "interface": "warp0",
    "dns": ["1.1.1.1", "1.0.0.1"],
    "mtu": 1420,
    "timeout": 60
  },
  "logging": {
    "level": "warn",
    "format": "text",
    "output": "stdout"
  }
}
```

### Debug Configuration

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
    "level": "debug",
    "format": "json",
    "output": "/var/log/darp-debug.log"
  }
}
```

## Troubleshooting Configuration

### Common Issues

#### Invalid JSON

```bash
# Check JSON syntax
cat /etc/darp/config.json | jq .

# Or use online JSON validator
```

#### Missing Required Fields

DARP will show specific error messages for missing required fields:

```bash
# Example error
Configuration validation failed: WARP endpoint must be configured
```

#### Permission Issues

```bash
# Check file permissions
ls -la /etc/darp/config.json

# Fix permissions if needed
sudo chmod 644 /etc/darp/config.json
```

## Best Practices

1. **Start with defaults**: Use the default configuration first
2. **Test changes**: Make small changes and test each one
3. **Backup config**: Keep a backup of working configurations
4. **Use user config**: Use `~/.config/darp/config.json` for personal settings
5. **Monitor logs**: Check logs when making configuration changes

## Next Steps

- [Command Reference](Command-Reference) - Learn all available commands
- [Troubleshooting](Troubleshooting) - Common issues and solutions
