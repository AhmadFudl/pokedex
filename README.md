# POKEDEX

A pokédex in a command-line REPL around [pokeapi](https://pokeapi.co).

Thanks to caching, no API was harmed in the main loop.

## Installation

```sh
$ go install https://github.com/ahmadfudl/pokedex
```

## Usage

```sh
$ pokedex
pokedex > help
Welcome to the Pokedex!
Usage:

map: List next 20 location areas.
mapb: List previous 20 location areas.
explore: List all the Pokémon located in the given area.
        explore <area_name>
catch: Catching Pokémon adds them to the user's Pokedex.
        catch <pokémon_name>
inspect: prints the name, height, weight, stats and type(s) of the Pokemon.
        inspect <pokémon_name>
pokedex: Prints a list of all the names of the Pokemon the user has caught.
exit: Exit the Pokedex.
help: Display a help message.
Pokedex >
```
