# API CLI

A Go CLI application that fetches data from multiple external APIs (Pokemon, Star Wars, and Punk Beer) using a clean architecture pattern with dependency injection.

## Features

- **Pokemon API**: Fetch Pokemon data from PokeAPI
- **Star Wars API**: Retrieve Star Wars films information
- **Punk Beer API**: Get craft beer data
- **All APIs**: Fetch data from all APIs sequentially in one command

## Architecture

The project follows a 3-layer clean architecture with clear separation of concerns:

```
main.go              # Dependency injection & Cobra command wiring
cmd/                 # Cobra command handlers (pokemon.go, starWars.go, punk.go, all.go)
internal/
  models/            # API response structs with json tags
  repository/        # HTTP client implementations per API
  service/           # Business logic layer (thin wrappers)
```

### Key Patterns

- **Domain Separation**: Each API (Pokemon, StarWars, Punk) has its own repository, service, and command triple
- **Dependency Injection**: Clean dependency flow from Repository → Service → Command → RootCmd
- **Interface-based Design**: All implementations return interface types for flexibility
- **Error Handling**: Non-blocking error handling that prints errors but continues execution

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Internet connection (for API calls)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd api-cli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o api-cli main.go
```

## Usage

### Available Commands

```bash
# Fetch Pokemon data
go run main.go pokemon

# Fetch Star Wars films
go run main.go starwars

# Fetch beer data
go run main.go punk

# Fetch all APIs sequentially
go run main.go all

# Show help
go run main.go --help
```

### Using the built binary

```bash
./api-cli pokemon
./api-cli starwars
./api-cli punk
./api-cli all
```

## API Endpoints

The application fetches data from the following public APIs:

- **Pokemon**: `https://pokeapi.co/api/v2/pokemon`
- **Star Wars**: `https://swapi.dev/api/films/`
- **Punk Beer**: `https://api.punkapi.com/v2/beers`

## Project Structure

```
├── main.go                          # Entry point with dependency injection
├── go.mod                           # Go module definition
├── go.sum                           # Go module checksums
├── cmd/                             # Command layer (Cobra commands)
│   ├── all.go                       # Command to fetch all APIs
│   ├── pokemon.go                   # Pokemon API command
│   ├── punk.go                      # Punk Beer API command
│   └── starWars.go                  # Star Wars API command
└── internal/                        # Internal application code
    ├── models/                      # Data models
    │   └── models.go                # API response structs
    ├── repository/                  # Data access layer
    │   ├── repositories.go          # Repository interfaces
    │   ├── pokemonRepository.go     # Pokemon API client
    │   ├── punkRepository.go        # Punk Beer API client
    │   └── starWarsRepository.go    # Star Wars API client
    └── service/                     # Business logic layer
        ├── services.go              # Service interfaces
        ├── pokemonService.go        # Pokemon business logic
        ├── punkService.go           # Punk Beer business logic
        └── starWarsService.go       # Star Wars business logic
```

## Development

### Adding a New API

To add support for a new API, follow these steps:

1. **Add Model**: Define the response struct in `internal/models/models.go` with appropriate JSON tags
2. **Create Repository**: Add interface in `internal/repository/repositories.go` and implementation in a new file
3. **Create Service**: Add interface in `internal/service/services.go` and implementation in a new file
4. **Create Command**: Add a new Cobra command in `cmd/` following the constructor pattern
5. **Wire Dependencies**: Update `main.go` to inject dependencies and add the command
6. **Update All Command**: Add the new API call to `cmd/all.go`

### Code Conventions

- **Interface Definitions**: Kept in separate files (`repositories.go`, `services.go`)
- **Constructor Pattern**: Functions return interface types, not concrete structs
- **Error Handling**: Print errors and continue execution (non-blocking)
- **HTTP Clients**: Direct `net/http` usage with `ioutil.ReadAll()` and `json.Unmarshal()`
- **Output**: Direct `fmt.Println()` for simplicity (no structured logging)

### Example: Repository Implementation

```go
func PokemonRepositoryImp() PokemonRepository {
    return &pokemonRepositoryImp{}
}

func (p *pokemonRepositoryImp) GetPokemon() []models.Pokemon {
    resp, err := http.Get("https://pokeapi.co/api/v2/pokemon")
    if err != nil {
        fmt.Println("Error fetching Pokemon:", err)
        return nil
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return nil
    }
    
    var result models.PokemonResponse
    json.Unmarshal(body, &result)
    return result.Results
}
```

### Dependency Injection Flow

```go
// In main.go
pokemonRepository := repository.PokemonRepositoryImp()
pokemonService := service.PokemonServiceImp(pokemonRepository)
rootCmd.AddCommand(cmd.InitPokemonCmd(pokemonService))
```

## Testing

Currently, no tests exist in the project. When adding tests, follow Go testing conventions:

- Unit tests for each layer (repository, service, command)
- Integration tests for API endpoints
- Mock interfaces for testing in isolation

## Contributing

1. Fork the repository
2. Create a feature branch
3. Follow the existing patterns and conventions
4. Add tests for new functionality
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).
