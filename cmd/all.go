package cmd

import (
	"fmt"
	"github.com/mayalaat/api-cli/internal/service"
	"github.com/spf13/cobra"
)

func InitAllCmd(
	pokeService service.PokemonService,
	warsService service.StarWarsService,
	punkService service.PunkService) *cobra.Command {
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "A brief description of your command",
		Run:   runAllFn(pokeService, warsService, punkService),
	}
	return allCmd
}

type AllCobraFn func(cmd *cobra.Command, args []string)

func runAllFn(
	pokeService service.PokemonService,
	warsService service.StarWarsService,
	punkService service.PunkService) AllCobraFn {
	return func(cmd *cobra.Command, args []string) {
		fmt.Println("FetchPokemons Begin")
		_, err := pokeService.FetchPokemons()
		fmt.Println("FetchPokemons Done")
		if err != nil {
			fmt.Print(err)
		}

		fmt.Println("FetchStarWars Begin")
		_, err = warsService.FetchStarWars()
		fmt.Println("FetchStarWars Done")
		if err != nil {
			fmt.Print(err)
		}

		fmt.Println("FetchPunkBeers Begin")
		_, err = punkService.FetchPunkBeers()
		fmt.Println("FetchPunkBeers Done")
		if err != nil {
			fmt.Print(err)
		}

	}
}
