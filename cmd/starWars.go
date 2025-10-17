package cmd

import (
	"fmt"
	"github.com/maalayat/api-cli/internal/service"

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
		starWars, err := service.FetchStarWars()
		if err != nil {
			fmt.Print(err)
		}
		if starWars.Results != nil {
			for _, film := range *starWars.Results {
				fmt.Printf("Title: %s, EpisodeId: %d, Director: %s, ReleaseDate: %s\n", film.Title, film.EpisodeId, film.Director, film.ReleaseDate)
			}
		}
	}
}
