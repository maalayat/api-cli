package cmd

import (
	"fmt"
	"github.com/maalayat/api-cli/internal/service"
	"github.com/spf13/cobra"
)

func InitAllCmd(
	pokeService service.PokemonService,
	warsService service.StarWarsService,
	breweryService service.BreweryService) *cobra.Command {
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "A brief description of your command",
		Run:   runAllFn(pokeService, warsService, breweryService),
	}
	return allCmd
}

type AllCobraFn func(cmd *cobra.Command, args []string)

func runAllFn(
	pokeService service.PokemonService,
	warsService service.StarWarsService,
	breweryService service.BreweryService) AllCobraFn {
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

		fmt.Println("FetchBreweries Begin")
		_, err = breweryService.FetchBreweries()
		fmt.Println("FetchBreweries Done")
		if err != nil {
			fmt.Print(err)
		}

	}
}
