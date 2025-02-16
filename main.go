package main

import (
	"github.com/ahmadfudl/pokedex/internal/pokecache"
	"time"
)

func main() {
	var config = &config_t{
		cache: pokecache.NewCache(time.Hour),
	}
	repl(config)
}
