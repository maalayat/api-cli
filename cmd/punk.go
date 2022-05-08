package cmd

import (
	"fmt"
	"github.com/mayalaat/api-cli/internal/service"

	"github.com/spf13/cobra"
)

func InitPunkCmd(service service.PunkService) *cobra.Command {
	punkCmd := &cobra.Command{
		Use:   "punk",
		Short: "A brief description of your command",
		Run:   runPunkFn(service),
	}
	return punkCmd
}

type PunkCobraFn func(cmd *cobra.Command, args []string)

func runPunkFn(service service.PunkService) PunkCobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers, err := service.FetchPunkBeers()
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(beers)

	}
}
