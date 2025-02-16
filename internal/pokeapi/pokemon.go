package pokeapi

import (
	"github.com/ahmadfudl/pokedex/internal/pokeapi/types"
	"github.com/ahmadfudl/pokedex/internal/pokecache"
)

func Pokemon(
	name string,
	cache *pokecache.Cache,
) (types.Pokemon, error) {

	var results types.Pokemon
	endpoint := "pokemon/" + name

	if value, ok := cache.Get(endpoint); ok {
		return value.(types.Pokemon), nil
	}

	err := get(endpoint, &results)
	if err != nil {
		return results, nil
	}

	cache.Add(endpoint, results)
	return results, err
}
