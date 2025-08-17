# ygggo_env

[![Go Version](https://img.shields.io/badge/go-1.24.5-blue.svg)](https://golang.org)
[![Test Coverage](https://img.shields.io/badge/coverage-93.1%25-brightgreen.svg)](https://github.com/yggai/ygggo_env)
[![License](https://img.shields.io/badge/license-PolyForm%20Noncommercial-orange.svg)](LICENSE)

A powerful and type-safe Go library for environment variable management with automatic `.env` file loading and type conversion.

## Features

- 🔍 **Automatic .env file discovery** - Searches from current directory upwards
- 🛡️ **Type-safe getters** - Built-in type conversion with default values
- 📝 **Multiple data formats** - Support for JSON arrays, maps, and comma-separated values
- 🧪 **Comprehensive testing** - 93.1% test coverage with TDD approach
- 🚀 **Zero dependencies** - Pure Go implementation
- 💪 **Robust error handling** - Graceful fallbacks to default values

## Supported Types

- **String** - `GetStr(key, defaultValue)`
- **Integer** - `GetInt(key, defaultValue)`
- **Float** - `GetFloat(key, defaultValue)`
- **Boolean** - `GetBool(key, defaultValue)` (supports: true/false, 1/0, yes/no, on/off)
- **Array** - `GetArr(key, defaultValue)` (supports: comma-separated and JSON arrays)
- **Map** - `GetMap(key, defaultValue)` (supports: JSON objects)

## Installation

```bash
go get github.com/yggai/ygggo_env
```

## Quick Start

### 1. Create a `.env` file

```env
# Database configuration
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=secret
DB_DATABASE=myapp

# Application settings
DEBUG=true
TIMEOUT=30.5
SERVERS=server1,server2,server3
CONFIG={"host": "localhost", "port": 8080, "ssl": true}
FEATURES=["feature1", "feature2", "feature3"]
```

### 2. Load and use environment variables

```go
package main

import (
    "fmt"
    gge "github.com/yggai/ygggo_env"
)

func main() {
    // Automatically find and load .env file
    gge.LoadEnv()

    // Get different types of environment variables
    host := gge.GetStr("DB_HOST", "localhost")
    port := gge.GetInt("DB_PORT", 3306)
    debug := gge.GetBool("DEBUG", false)
    timeout := gge.GetFloat("TIMEOUT", 10.0)

    // Get arrays and maps
    servers := gge.GetArr("SERVERS", []string{"default_server"})
    config := gge.GetMap("CONFIG", map[string]interface{}{"default": "value"})

    fmt.Printf("Connecting to %s:%d\n", host, port)
    fmt.Printf("Debug mode: %t\n", debug)
    fmt.Printf("Servers: %v\n", servers)
    fmt.Printf("Config: %v\n", config)
}
```

## API Reference

### LoadEnv()

Automatically searches for `.env` files starting from the current directory and moving upwards. Loads and parses the first `.env` file found.

```go
err := gge.LoadEnv()
if err != nil {
    log.Fatal(err)
}
```

### Type-Safe Getters

#### GetStr(key, defaultValue)

Gets a string environment variable with a default value.

```go
host := gge.GetStr("DB_HOST", "localhost")
username := gge.GetStr("DB_USERNAME", "root")
```

#### GetInt(key, defaultValue)

Gets an integer environment variable with type conversion and default value.

```go
port := gge.GetInt("DB_PORT", 3306)
maxConnections := gge.GetInt("MAX_CONNECTIONS", 100)
```

#### GetFloat(key, defaultValue)

Gets a float64 environment variable with type conversion and default value.

```go
timeout := gge.GetFloat("TIMEOUT", 30.0)
rate := gge.GetFloat("RATE_LIMIT", 1.5)
```

#### GetBool(key, defaultValue)

Gets a boolean environment variable. Supports multiple boolean representations:
- `true`, `false`
- `1`, `0`
- `yes`, `no`
- `on`, `off`

Case-insensitive.

```go
debug := gge.GetBool("DEBUG", false)
enableSSL := gge.GetBool("ENABLE_SSL", true)
```

#### GetArr(key, defaultValue)

Gets an array of strings. Supports two formats:

**Comma-separated:**
```env
SERVERS=server1,server2,server3
```

**JSON array:**
```env
FEATURES=["feature1", "feature2", "feature3"]
```

```go
servers := gge.GetArr("SERVERS", []string{"localhost"})
features := gge.GetArr("FEATURES", []string{"default"})
```

#### GetMap(key, defaultValue)

Gets a map from JSON-formatted environment variable.

```env
CONFIG={"host": "localhost", "port": 8080, "ssl": true}
```

```go
config := gge.GetMap("CONFIG", map[string]interface{}{
    "host": "localhost",
    "port": 8080,
})
```

## Examples

The `examples/` directory contains complete working examples:

- **`c01_load_env/`** - Basic environment loading
- **`c02_get_env/`** - Type-safe getters usage
- **`c03_all_types/`** - Comprehensive example with all data types

Run examples:

```bash
# Basic loading
go run ./examples/c01_load_env/main.go

# Type-safe getters
go run ./examples/c02_get_env/main.go

# All types demonstration
go run ./examples/c03_all_types/main.go
```

## .env File Format

The library supports standard `.env` file format:

```env
# Comments are supported
KEY=value

# Strings (no quotes needed)
DB_HOST=localhost
DB_NAME=myapp

# Numbers
DB_PORT=3306
TIMEOUT=30.5

# Booleans (multiple formats supported)
DEBUG=true
ENABLE_CACHE=1
USE_SSL=yes

# Arrays (comma-separated)
SERVERS=server1,server2,server3

# Arrays (JSON format)
FEATURES=["auth", "cache", "logging"]

# Objects (JSON format)
CONFIG={"host": "localhost", "port": 8080}
```

## Testing

The library is thoroughly tested with 93.1% code coverage:

```bash
# Run tests
go test -v

# Run tests with coverage
go test -v -cover

# Generate coverage report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Error Handling

The library follows a graceful error handling approach:

- **File not found**: Returns `nil` (no error) - missing `.env` files are acceptable
- **Parse errors**: Invalid lines are skipped with detailed error messages
- **Type conversion errors**: Returns default values instead of panicking
- **JSON parse errors**: Returns default values for malformed JSON

## Performance

- **Zero allocations** for simple string operations
- **Minimal memory footprint** - no external dependencies
- **Fast file parsing** with efficient string operations
- **Cached environment access** using Go's built-in `os.Getenv()`

## Contributing

This is a personal research project. While code contributions are not accepted, you are welcome to:

- 🐛 **Report issues** - Please open an issue for bugs or feature requests
- 💡 **Suggest improvements** - Share your ideas through issues
- 📖 **Improve documentation** - Documentation improvements are welcome

## License

This project is licensed under the **PolyForm Noncommercial License 1.0.0**.

See the [LICENSE](LICENSE) file for details.

## Author

**源滚滚** (Yuan Gungun)
📧 Email: 1156956636@qq.com

---

⭐ If you find this project useful, please consider giving it a star!

## Changelog

### v1.0.0
- Initial release
- Automatic `.env` file discovery
- Type-safe getters for all basic types
- Support for arrays and maps
- Comprehensive test coverage (93.1%)
- Zero external dependencies
