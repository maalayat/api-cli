/*
Copyright © 2022 Alejandro Ayala <alejandro.ayala@solmedia.ec>

*/
package main

import (
	"github.com/mayalaat/api-cli/cmd"
	"github.com/mayalaat/api-cli/internal/repository"
	"github.com/mayalaat/api-cli/internal/service"
	"github.com/spf13/cobra"
)

func main() {
	pokemonRepository := repository.PokemonRepositoryImp()
	pokemonService := service.PokemonServiceImp(pokemonRepository)

	starWarsRepository := repository.StarWarsRepositoryImp()
	starWarsService := service.StarWarsServiceImp(starWarsRepository)

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmd.InitPokemonCmd(pokemonService))
	rootCmd.AddCommand(cmd.InitStarWarsCmd(starWarsService))
	err := rootCmd.Execute()
	if err != nil {
		return
	}
}
