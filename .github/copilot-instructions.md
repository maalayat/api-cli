# Copilot Instructions for API CLI

This is a Go CLI application that fetches data from multiple external APIs (Pokemon, Star Wars, and Punk Beer) using a clean architecture pattern with dependency injection.

## Architecture Overview

The project follows a 3-layer architecture with clear separation of concerns:

```
main.go              # Dependency injection & Cobra command wiring
cmd/                 # Cobra command handlers (pokemon.go, starWars.go, punk.go, all.go)
internal/
  models/            # API response structs with json tags
  repository/        # HTTP client implementations per API
  service/           # Business logic layer (thin wrappers)
```

**Key Pattern**: Each API domain (Pokemon, StarWars, Punk) has its own repository, service, and command triple. The `all.go` command orchestrates all three APIs sequentially.

## Development Patterns

### Repository Layer
- Direct `net/http` usage with `ioutil.ReadAll()` and `json.Unmarshal()`
- Concrete implementations like `PokemonRepositoryImp()` return interface types
- Error handling prints but doesn't halt execution
- Example: `internal/repository/pokemonRepository.go` shows the HTTP pattern

### Service Layer  
- Thin wrappers that delegate directly to repositories
- No business logic - just interface adaptation
- Constructor pattern: `PokemonServiceImp(repo PokemonRepository)`

### Command Layer
- Cobra commands use dependency injection via constructor functions like `InitPokemonCmd(service)`
- Higher-order functions pattern: `runPokemonFn(service) PokemonCobraFn`
- Direct `fmt.Println()` output - no structured logging

### Dependency Injection
In `main.go`, dependencies flow: Repository → Service → Command → RootCmd
```go
pokemonRepository := repository.PokemonRepositoryImp()
pokemonService := service.PokemonServiceImp(pokemonRepository) 
rootCmd.AddCommand(cmd.InitPokemonCmd(pokemonService))
```

## API Endpoints
- Pokemon: `https://pokeapi.co/api/v2/pokemon`
- Star Wars: `https://swapi.dev/api/films/`  
- Punk Beer: `https://api.punkapi.com/v2/beers`

## Common Commands
```bash
go run main.go pokemon    # Fetch Pokemon data
go run main.go starwars   # Fetch Star Wars films
go run main.go punk       # Fetch beer data  
go run main.go all        # Fetch all APIs sequentially
```

## Adding New APIs
1. Add model struct in `internal/models/models.go` with json tags
2. Create repository interface and implementation in `internal/repository/`
3. Create service interface and implementation in `internal/service/`
4. Create Cobra command in `cmd/` following the constructor pattern
5. Wire up in `main.go` and add to `all.go` command

## Code Conventions
- Interface definitions in separate files (`repositories.go`, `services.go`)
- Constructor functions return interface types, not concrete structs
- Error handling: print and continue (no error propagation stopping execution)
- No tests currently exist - new tests should follow Go testing conventions
