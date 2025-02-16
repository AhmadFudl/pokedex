package main

import (
	"fmt"
	"github.com/ahmadfudl/pokedex/internal/pokeapi"
	"github.com/ahmadfudl/pokedex/internal/pokeapi/types"
	"math/rand/v2"
	"os"
)

type (
	command_t struct {
		name        string
		description string
		callback    func(*config_t, []string) error
	}
	pokedex_t map[string]types.Pokemon
)

var commands map[string]command_t
var pokedex pokedex_t

func init() {
	commands = map[string]command_t{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    _exit,
		},
		"help": {
			name:        "help",
			description: "Display a help message.",
			callback:    _help,
		},
		"map": {
			name:        "map",
			description: "List next 20 location areas.",
			callback:    _mapn,
		},
		"mapb": {
			name:        "mapb",
			description: "List previous 20 location areas.",
			callback:    _mapb,
		},
		"explore": {
			name: "explore",
			description: `List all the Pokémon located in the given area.
	explore <area_name>`,
			callback: _explore,
		},
		"catch": {
			name: "catch",
			description: `Catching Pokémon adds them to the user's Pokedex.
	catch <pokémon_name>`,
			callback: _catch,
		},
		"inspect": {
			name: "inspect",
			description: `prints the name, height, weight, stats and type(s) of the Pokemon.
	inspect <pokémon_name>`,
			callback: _inspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: `Prints a list of all the names of the Pokemon the user has caught.`,
			callback:    _pokedex,
		},
	}

	pokedex = make(map[string]types.Pokemon)
}

func _exit(config *config_t, args []string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func _help(config *config_t, args []string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

func _mapn(config *config_t, args []string) error {
	return _map(config.next_url, config)
}

func _mapb(config *config_t, args []string) error {
	if config.prev_url == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return _map(config.prev_url, config)
}

func _map(url string, config *config_t) error {
	results, err := pokeapi.Location_area_l(url, config.cache)
	if err != nil {
		return err
	}

	for _, result := range results.Results {
		fmt.Println(result.Name)
	}

	config.next_url = results.Next
	config.prev_url = results.Previous
	return nil
}

func _explore(config *config_t, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: explore <area_name>")
	}

	results, err := pokeapi.Location_area(args[0], config.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", args[0])

	for _, encounters := range results.Pokemon_encounters {
		fmt.Printf("- %s\n", encounters.Pokemon.Name)
	}

	return nil
}

func _catch(config *config_t, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: catch <pokémon_name>")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	results, err := pokeapi.Pokemon(args[0], config.cache)
	if err != nil {
		return err
	}

	delta := results.Base_experience / 10
	guess := rand.IntN(results.Base_experience)

	luck := func(v int) int {
		y := v >> 63
		return (v ^ y) - y
	}(guess - rand.IntN(results.Base_experience))

	if luck <= delta {
		fmt.Printf("%s was caught!\n", results.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		pokedex[results.Name] = results
	} else {
		fmt.Printf("%s escaped!\n", results.Name)
	}

	return nil
}

func _inspect(config *config_t, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Usage: inspect <pokémon_name>")
	}

	if pokemon, ok := pokedex[args[0]]; !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Println("Name:", pokemon.Name)
		fmt.Println("Height:", pokemon.Height)
		fmt.Println("Weight:", pokemon.Weight)

		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t-%s: %d\n", stat.Stat.Name, stat.Base_stat)
		}

		fmt.Println("Type:")
		for _, t := range pokemon.Types {
			fmt.Println("\t-", t.Type.Name)
		}
	}
	return nil
}

func _pokedex(config *config_t, args []string) error {
	_ = config
	_ = args
	fmt.Println("Your Pokedex:")
	for k := range pokedex {
		fmt.Println("\t-", k)
	}
	return nil
}
