package cmd

import (
	"fmt"
	"github.com/mayalaat/api-cli/internal/service"

	"github.com/spf13/cobra"
)

func InitStarWarsCmd(service service.StarWarsService) *cobra.Command {
	starWarsCmd := &cobra.Command{
		Use:   "starWars",
		Short: "A brief description of your command",
		Run:   runStarWarsFn(service),
	}
	return starWarsCmd
}

type StarWarsCobraFn func(cmd *cobra.Command, args []string)

func runStarWarsFn(service service.StarWarsService) StarWarsCobraFn {
	return func(cmd *cobra.Command, args []string) {
		starWars, err := service.FetchStarWars(nil)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(starWars)
	}
}
