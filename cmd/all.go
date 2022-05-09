package cmd

import (
	"fmt"
	"github.com/mayalaat/api-cli/internal/service"
	"sync"

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
		waitGroup := sync.WaitGroup{}
		waitGroup.Add(3)
		go func() {
			fmt.Println("FetchPokemons Begin")
			_, err := pokeService.FetchPokemons(&waitGroup)
			if err != nil {
				fmt.Print(err)
			}

			fmt.Println("FetchPokemons Done")
		}()

		go func() {
			fmt.Println("FetchStarWars Begin")
			_, err := warsService.FetchStarWars(&waitGroup)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Println("FetchStarWars Done")
		}()

		go func() {
			fmt.Println("FetchPunkBeers Begin")
			_, err := punkService.FetchPunkBeers(&waitGroup)
			if err != nil {
				fmt.Print(err)
			}
			fmt.Println("FetchPunkBeers Done")
		}()

		waitGroup.Wait()
	}
}
