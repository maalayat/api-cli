# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go CLI application that fetches data from multiple APIs (Pokemon, Star Wars, and Punk API for beers). It uses the Cobra library for command-line interface functionality and follows a clean architecture pattern with repository and service layers.

## Architecture

The codebase follows a layered architecture:

- `main.go` - Entry point that wires up dependencies and initializes Cobra commands
- `cmd/` - Cobra command definitions for each API endpoint plus an "all" command
- `internal/models/` - Data models for API responses (Pokemon, StarWars, Punk structs)
- `internal/repository/` - Repository interfaces and implementations for HTTP API calls
- `internal/service/` - Service interfaces and implementations that orchestrate repository calls

The dependency injection pattern is used in main.go where repositories are injected into services, and services are injected into commands.

## Common Development Commands

```bash
# Build the application
go build

# Run the application
go run main.go

# Run with specific commands
go run main.go pokemon
go run main.go starwars  
go run main.go punk
go run main.go all

# Test the application
go test ./...

# Format code
go fmt ./...

# Lint/vet code
go vet ./...

# Clean build artifacts
go clean
```

## Key Implementation Details

- Each API (Pokemon, Star Wars, Punk) has its own repository, service, and command
- The "all" command runs all three API fetches sequentially
- HTTP clients are implemented directly in repositories using `net/http`
- JSON unmarshaling is handled in repository layer
- Error handling prints errors but doesn't stop execution
- No tests are currently implemented in the codebase

## Dependencies

- `github.com/spf13/cobra` - CLI framework
- Standard Go libraries for HTTP and JSON processing