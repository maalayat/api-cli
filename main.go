/*
Copyright © 2022 Alejandro Ayala <alejandro.ayala@solmedia.ec>
*/
package main

import (
	"github.com/maalayat/api-cli/cmd"
	"github.com/maalayat/api-cli/internal/repository"
	"github.com/maalayat/api-cli/internal/service"
	"github.com/spf13/cobra"
)

func main() {
	pokemonRepository := repository.PokemonRepositoryImp()
	pokemonService := service.PokemonServiceImp(pokemonRepository)

	starWarsRepository := repository.StarWarsRepositoryImp()
	starWarsService := service.StarWarsServiceImp(starWarsRepository)

	breweryRepository := repository.BreweryRepositoryImp()
	breweryService := service.BreweryServiceImp(breweryRepository)

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.InitPokemonCmd(pokemonService))
	rootCmd.AddCommand(cmd.InitStarWarsCmd(starWarsService))
	rootCmd.AddCommand(cmd.InitBreweryCmd(breweryService))
	rootCmd.AddCommand(cmd.InitAllCmd(pokemonService, starWarsService, breweryService))
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
