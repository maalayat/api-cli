package cmd

import (
	"fmt"
	"github.com/mayalaat/api-cli/internal/service"

	"github.com/spf13/cobra"
)

func InitPokemonCmd(pokemonService service.PokemonService) *cobra.Command {
	pokemonCmd := &cobra.Command{
		Use:   "pokemon",
		Short: "Get pokemons from API",
		Run:   runPokemonFn(pokemonService),
	}

	return pokemonCmd
}

type PokemonCobraFn func(cmd *cobra.Command, args []string)

func runPokemonFn(pokemonService service.PokemonService) PokemonCobraFn {
	return func(cmd *cobra.Command, args []string) {
		pokemons, err := pokemonService.FetchPokemons()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(pokemons)

	}
}
