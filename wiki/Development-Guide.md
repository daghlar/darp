# Development Guide

This guide helps developers contribute to DARP development.

## Getting Started

### Prerequisites

- **Go 1.21+**: Required for building DARP
- **Git**: For version control
- **WireGuard tools**: For testing
- **Arch Linux**: Recommended development environment

### Setting Up Development Environment

```bash
# Clone the repository
git clone https://github.com/daghlar/darp.git
cd darp

# Install dependencies
go mod download

# Build the project
make build

# Run tests
make test
```

## Project Structure

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
└── README.md           # Project documentation
```

## Development Workflow

### 1. Fork and Clone

```bash
# Fork the repository on GitHub
# Then clone your fork
git clone https://github.com/yourusername/darp.git
cd darp

# Add upstream remote
git remote add upstream https://github.com/daghlar/darp.git
```

### 2. Create Feature Branch

```bash
# Create and switch to feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description
```

### 3. Make Changes

```bash
# Make your changes
# Test your changes
make test

# Build and test
make build
./build/darp --version
```

### 4. Commit Changes

```bash
# Add changes
git add .

# Commit with descriptive message
git commit -m "Add feature: brief description

- Detailed explanation of changes
- Any additional notes"
```

### 5. Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
```

## Code Style

### Go Code Style

Follow standard Go conventions:

```go
// Package comment
package cli

// Function comment
func NewCLI() *CLI {
    // Implementation
}
```

### Naming Conventions

- **Packages**: lowercase, single word
- **Functions**: PascalCase for public, camelCase for private
- **Variables**: camelCase
- **Constants**: PascalCase or UPPER_CASE

### Comments

- **No comments in code**: DARP follows a no-comment policy
- **README and docs**: Comprehensive documentation
- **Commit messages**: Clear and descriptive

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run specific package tests
go test ./pkg/warp/...
```

### Writing Tests

Create test files with `_test.go` suffix:

```go
package warp

import (
    "testing"
)

func TestNewClient(t *testing.T) {
    client := NewClient()
    if client == nil {
        t.Error("NewClient() returned nil")
    }
}
```

### Test Coverage

Aim for high test coverage:

```bash
# Check coverage
go test ./... -cover

# Generate coverage report
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Building

### Development Build

```bash
# Quick build
make build

# Build with debug info
go build -ldflags "-X main.version=dev -X main.build=$(date -u '+%Y-%m-%d_%H:%M:%S')" -o build/darp ./cmd/darp
```

### Production Build

```bash
# Full build with all targets
make build-all

# Or use build script
./build.sh
```

### Cross-Platform Build

```bash
# Build for different architectures
GOOS=linux GOARCH=amd64 go build -o darp-linux-amd64 ./cmd/darp
GOOS=linux GOARCH=arm64 go build -o darp-linux-arm64 ./cmd/darp
```

## Adding Features

### 1. Plan the Feature

- **Identify requirements**: What does the feature need to do?
- **Design the API**: How will users interact with it?
- **Consider impact**: How does it affect existing code?

### 2. Implement the Feature

- **Create new files** if needed
- **Add to existing packages** if appropriate
- **Update CLI commands** if user-facing
- **Add configuration options** if needed

### 3. Update Documentation

- **README.md**: Update if user-facing
- **Wiki pages**: Add relevant documentation
- **Code comments**: None (policy)

### 4. Add Tests

- **Unit tests** for new functions
- **Integration tests** for new features
- **CLI tests** for new commands

### Example: Adding a New Command

```go
// In pkg/cli/cli.go
func (c *CLI) newCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "newcommand",
        Short: "Description of new command",
        Long:  "Detailed description of new command",
        RunE: func(cmd *cobra.Command, args []string) error {
            return c.handleNewCommand()
        },
    }
}

func (c *CLI) handleNewCommand() error {
    // Implementation
    return nil
}
```

## Configuration Changes

### Adding New Configuration Options

1. **Update config struct**:
```go
type Config struct {
    Cloudflare CloudflareConfig `json:"cloudflare"`
    Network    NetworkConfig    `json:"network"`
    Logging    LoggingConfig    `json:"logging"`
    NewSection NewSectionConfig `json:"new_section"` // New field
}
```

2. **Update default config**:
```go
func DefaultConfig() *Config {
    return &Config{
        // ... existing fields
        NewSection: NewSectionConfig{
            NewOption: "default_value",
        },
    }
}
```

3. **Update validation**:
```go
func (c *Config) Validate() error {
    // ... existing validation
    if c.NewSection.NewOption == "" {
        return fmt.Errorf("new option is required")
    }
    return nil
}
```

## CLI Changes

### Adding New Commands

1. **Create command function**:
```go
func (c *CLI) newCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "newcommand",
        Short: "Description",
        RunE: func(cmd *cobra.Command, args []string) error {
            return c.handleNewCommand()
        },
    }
}
```

2. **Add to setupCommands**:
```go
func (c *CLI) setupCommands() {
    // ... existing commands
    c.rootCmd.AddCommand(c.newCommand())
}
```

3. **Implement handler**:
```go
func (c *CLI) handleNewCommand() error {
    // Implementation
    return nil
}
```

## Debugging

### Enable Debug Logging

```bash
# Set debug level
darp config set logging.level "debug"

# Run with verbose output
darp --verbose connect
```

### Common Debug Commands

```bash
# Check configuration
darp config show

# Test connectivity
darp test connectivity

# Check logs
sudo journalctl -u darp -f
```

### Debugging Tips

1. **Use verbose mode**: `--verbose` flag
2. **Check logs**: systemd journal
3. **Test components**: Individual package tests
4. **Use debugger**: Delve or VS Code debugger

## Performance Optimization

### Profiling

```bash
# CPU profiling
go test -cpuprofile=cpu.prof ./...

# Memory profiling
go test -memprofile=mem.prof ./...

# Analyze profiles
go tool pprof cpu.prof
go tool pprof mem.prof
```

### Benchmarking

```bash
# Run benchmarks
go test -bench=. ./...

# Run specific benchmark
go test -bench=BenchmarkFunctionName ./pkg/package
```

## Contributing Guidelines

### Pull Request Process

1. **Fork the repository**
2. **Create feature branch**
3. **Make changes**
4. **Add tests**
5. **Update documentation**
6. **Submit pull request**

### Pull Request Template

```markdown
## Description
Brief description of changes

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
- [ ] Tests pass
- [ ] Manual testing completed
- [ ] No regressions

## Checklist
- [ ] Code follows style guidelines
- [ ] Self-review completed
- [ ] Documentation updated
- [ ] Tests added/updated
```

### Code Review

- **Review all changes** before merging
- **Test the changes** locally
- **Check documentation** updates
- **Verify tests** pass

## Release Process

### Versioning

DARP uses semantic versioning (MAJOR.MINOR.PATCH):

- **MAJOR**: Breaking changes
- **MINOR**: New features
- **PATCH**: Bug fixes

### Creating Releases

1. **Update version** in `cmd/darp/main.go`
2. **Update CHANGELOG.md**
3. **Create git tag**:
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
4. **Create GitHub release**

## Resources

### Documentation

- [Go Documentation](https://golang.org/doc/)
- [Cobra CLI Library](https://github.com/spf13/cobra)
- [WireGuard Documentation](https://www.wireguard.com/)

### Tools

- **Go**: Programming language
- **Cobra**: CLI framework
- **WireGuard**: VPN protocol
- **systemd**: Service management

### Community

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: Questions and discussions
- **Pull Requests**: Code contributions

## Getting Help

### Development Questions

1. **Check documentation**: This guide and other wiki pages
2. **Search issues**: Look for similar problems
3. **Ask in discussions**: Create a new discussion
4. **Review code**: Look at existing implementations

### Useful Commands

```bash
# Development workflow
make clean && make build && make test

# Check code style
gofmt -d .

# Run specific tests
go test -run TestFunctionName ./pkg/package

# Build and test
./build.sh && ./build/darp --version
```

## Next Steps

- [Installation Guide](Installation-Guide) - Set up development environment
- [Configuration Reference](Configuration-Reference) - Understand configuration
- [Command Reference](Command-Reference) - Learn available commands
- [Troubleshooting](Troubleshooting) - Debug issues
