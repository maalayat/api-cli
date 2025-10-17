package cmd

import (
	"fmt"
	"github.com/maalayat/api-cli/internal/service"

	"github.com/spf13/cobra"
)

func InitBreweryCmd(service service.BreweryService) *cobra.Command {
	breweryCmd := &cobra.Command{
		Use:   "brewery",
		Short: "A brief description of your command",
		Run:   runBrewerFn(service),
	}
	return breweryCmd
}

type BreweryCobraFn func(cmd *cobra.Command, args []string)

func runBrewerFn(service service.BreweryService) BreweryCobraFn {
	return func(cmd *cobra.Command, args []string) {
		breweries, err := service.FetchBreweries()
		if err != nil {
			fmt.Print(err)
		}
		if breweries != nil {
			for _, brewery := range breweries {
				fmt.Printf("Name: %s, Brewery Type: %s, State: %s, Country: %s\n",
					brewery.Name, brewery.BreweryType, brewery.State, brewery.Country)
			}
		}

	}
}
