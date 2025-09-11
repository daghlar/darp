# Contributing to DARP

We welcome contributions to DARP! This guide explains how to contribute effectively.

## How to Contribute

### Types of Contributions

- **Bug fixes**: Fix issues and improve stability
- **New features**: Add functionality and capabilities
- **Documentation**: Improve guides and documentation
- **Testing**: Add tests and improve coverage
- **Performance**: Optimize code and improve speed
- **UI/UX**: Improve command-line interface

## Getting Started

### 1. Fork the Repository

1. Go to [DARP on GitHub](https://github.com/daghlar/darp)
2. Click "Fork" button
3. Clone your fork:
   ```bash
   git clone https://github.com/yourusername/darp.git
   cd darp
   ```

### 2. Set Up Development Environment

```bash
# Add upstream remote
git remote add upstream https://github.com/daghlar/darp.git

# Install dependencies
go mod download

# Build the project
make build

# Run tests
make test
```

### 3. Create a Branch

```bash
# Create feature branch
git checkout -b feature/your-feature-name

# Or for bug fixes
git checkout -b fix/issue-description
```

## Development Process

### 1. Make Changes

- **Write clean code**: Follow Go conventions
- **No comments**: DARP follows a no-comment policy
- **Test your changes**: Ensure everything works
- **Update documentation**: If needed

### 2. Test Your Changes

```bash
# Run all tests
make test

# Run specific tests
go test ./pkg/warp/...

# Build and test
make build
./build/darp --version
```

### 3. Commit Changes

```bash
# Add changes
git add .

# Commit with descriptive message
git commit -m "Add feature: brief description

- Detailed explanation of changes
- Any additional notes"
```

### 4. Push and Create PR

```bash
# Push to your fork
git push origin feature/your-feature-name

# Create pull request on GitHub
```

## Code Style Guidelines

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

### Code Organization

- **Keep functions small**: Single responsibility
- **Use meaningful names**: Clear and descriptive
- **Group related code**: Logical organization
- **No comments**: Code should be self-explanatory

## Testing Guidelines

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

### Test Types

- **Unit tests**: Test individual functions
- **Integration tests**: Test component interactions
- **CLI tests**: Test command-line interface
- **End-to-end tests**: Test complete workflows

## Documentation Guidelines

### README Updates

- **Keep it current**: Update when adding features
- **Be clear**: Use simple, clear language
- **Include examples**: Show how to use features
- **Update installation**: If process changes

### Wiki Updates

- **Update existing pages**: When relevant
- **Add new pages**: For new features
- **Keep organized**: Use consistent structure
- **Test examples**: Ensure all code examples work

### Code Documentation

- **No comments**: DARP follows no-comment policy
- **Self-documenting code**: Use clear names and structure
- **README and docs**: Comprehensive external documentation

## Pull Request Guidelines

### Before Submitting

1. **Test thoroughly**: Ensure all tests pass
2. **Check style**: Follow code style guidelines
3. **Update docs**: If needed
4. **Self-review**: Review your own changes

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

### Review Process

1. **Automated checks**: CI/CD runs tests
2. **Code review**: Maintainers review changes
3. **Feedback**: Address any feedback
4. **Merge**: Once approved

## Issue Guidelines

### Reporting Bugs

When reporting bugs, include:

1. **System information**:
   ```bash
   uname -a
   pacman -Q wireguard-tools
   darp --version
   ```

2. **Steps to reproduce**: Clear, numbered steps
3. **Expected behavior**: What should happen
4. **Actual behavior**: What actually happens
5. **Logs**: Relevant error messages

### Requesting Features

When requesting features:

1. **Use case**: Why is this feature needed?
2. **Proposed solution**: How should it work?
3. **Alternatives**: Other ways to solve the problem
4. **Additional context**: Any other relevant information

## Development Workflow

### Daily Workflow

```bash
# Start of day
git checkout main
git pull upstream main

# Work on feature
git checkout -b feature/new-feature
# ... make changes ...
make test
make build

# End of day
git add .
git commit -m "WIP: feature description"
git push origin feature/new-feature
```

### Weekly Workflow

```bash
# Sync with upstream
git checkout main
git pull upstream main
git push origin main

# Update feature branches
git checkout feature/your-feature
git rebase main
```

## Release Process

### Versioning

DARP uses semantic versioning (MAJOR.MINOR.PATCH):

- **MAJOR**: Breaking changes
- **MINOR**: New features
- **PATCH**: Bug fixes

### Release Checklist

- [ ] All tests pass
- [ ] Documentation updated
- [ ] Version bumped
- [ ] CHANGELOG updated
- [ ] Release notes written
- [ ] Tag created
- [ ] GitHub release created

## Community Guidelines

### Code of Conduct

- **Be respectful**: Treat everyone with respect
- **Be constructive**: Provide helpful feedback
- **Be patient**: Remember that everyone is learning
- **Be inclusive**: Welcome contributors of all backgrounds

### Communication

- **GitHub Issues**: Bug reports and feature requests
- **GitHub Discussions**: Questions and general discussion
- **Pull Requests**: Code review and discussion
- **Commit Messages**: Clear and descriptive

## Getting Help

### Development Questions

1. **Check documentation**: This guide and other wiki pages
2. **Search issues**: Look for similar problems
3. **Ask in discussions**: Create a new discussion
4. **Review code**: Look at existing implementations

### Useful Resources

- [Go Documentation](https://golang.org/doc/)
- [Cobra CLI Library](https://github.com/spf13/cobra)
- [WireGuard Documentation](https://www.wireguard.com/)
- [Git Documentation](https://git-scm.com/doc)

## Recognition

### Contributors

Contributors are recognized in:

- **README.md**: Major contributors
- **GitHub**: Commit history and contributions
- **Releases**: Release notes and changelog

### Types of Contributions

- **Code**: Bug fixes, features, improvements
- **Documentation**: Guides, examples, tutorials
- **Testing**: Test cases, bug reports
- **Community**: Helping others, discussions

## Next Steps

1. **Fork the repository**
2. **Set up development environment**
3. **Pick an issue** or create a new one
4. **Make your changes**
5. **Submit a pull request**

## Useful Commands

```bash
# Development
make clean && make build && make test

# Git workflow
git status
git diff
git log --oneline

# Testing
go test ./...
go test -v ./pkg/warp
go test -cover ./...

# Building
make build
./build/darp --version
```

## Questions?

If you have questions about contributing:

1. **Check this guide**: Most questions are answered here
2. **Search discussions**: Look for similar questions
3. **Create discussion**: Ask your question
4. **Open issue**: If you found a bug

Thank you for contributing to DARP! 🎉
