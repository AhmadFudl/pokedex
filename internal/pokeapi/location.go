package pokeapi

import (
	"fmt"
	"github.com/ahmadfudl/pokedex/internal/pokeapi/types"
	"github.com/ahmadfudl/pokedex/internal/pokecache"
	"net/url"
)

// return a location area list for pagination
func Location_area_l(
	u string,
	cache *pokecache.Cache,
) (types.NamedResourceList, error) {

	var results types.NamedResourceList

	pu, err := url.Parse(u)
	if err != nil {
		return results, err
	}

	limit := pu.Query().Get("limit")
	offset := pu.Query().Get("offset")
	endpoint := fmt.Sprintf("location-area?offset=%s&limit=%s", offset, limit)

	if value, ok := cache.Get(endpoint); ok {
		return value.(types.NamedResourceList), nil
	}

	err = get(endpoint, &results)
	if err != nil {
		return results, err
	}

	cache.Add(endpoint, results)
	return results, err
}

func Location_area(
	name string,
	cache *pokecache.Cache,
) (types.LocationArea, error) {

	var results types.LocationArea
	endpoint := "location-area/" + name

	if value, ok := cache.Get(endpoint); ok {
		return value.(types.LocationArea), nil
	}

	err := get(endpoint, &results)
	if err != nil {
		return results, nil
	}

	cache.Add(endpoint, results)
	return results, err
}
