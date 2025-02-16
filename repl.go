package main

import (
	"bufio"
	"fmt"
	"github.com/ahmadfudl/pokedex/internal/pokecache"
	"os"
	"strings"
)

type config_t struct {
	next_url string
	prev_url string
	cache    *pokecache.Cache
}

func clean_input(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func repl(config *config_t) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()

		parts := clean_input(scanner.Text())
		if len(parts) == 0 {
			continue
		}

		command, ok := commands[parts[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(config, parts[1:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
