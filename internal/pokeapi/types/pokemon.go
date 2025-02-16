package types

type (
	Pokemon struct {
		Id                       int                `json:"id"`
		Name                     string             `json:"name"`
		Base_experience          int                `json:"base_experience"`
		Height                   int                `json:"height"`
		Is_default               bool               `json:"is_default"`
		Order                    int                `json:"order"`
		Weight                   int                `json:"weight"`
		Abilities                []PokemonAbility   `json:"abilities"`
		Forms                    []NamedResource    `json:"forms"`
		Game_indices             []VersionGameIndex `json:"game_indices"`
		Held_items               []PokemonHeldItem  `json:"held_items"`
		Location_area_encounters string             `json:"location_area_encounters"`
		Moves                    []PokemonMove      `json:"moves"`
		Past_types               []PokemonTypePast  `json:"past_types"`
		Sprites                  PokemonSprites     `json:"sprites"`
		Cires                    PokemonCries       `json:"cries"`
		Species                  NamedResource      `json:"species"`
		Stats                    []PokemonStat      `json:"stats"`
		Types                    []PokemonType      `json:"types"`
	}

	PokemonAbility struct {
		Is_hidden bool          `json:"is_hidden"`
		Slot      int           `json:"slot"`
		Ability   NamedResource `json:"ability"`
	}

	PokemonHeldItem struct {
		Item            NamedResource            `json:"item"`
		Version_details []PokemonHeldItemVersion `json:"version_details"`
	}

	PokemonHeldItemVersion struct {
		Version NamedResource `json:"version"`
		Rarity  int           `json:"rarity"`
	}

	PokemonMove struct {
		Move                  NamedResource        `json:"move"`
		Version_group_details []PokemonMoveVersion `json:"version_group_details"`
	}

	PokemonMoveVersion struct {
		Move_learn_method NamedResource `json:"move_learn_method"`
		Version_group     NamedResource `json:"version_group"`
		Level_learned_at  int           `json:"level_learned_at"`
	}

	PokemonTypePast struct {
		Generation NamedResource `json:"generation"`
		Types      []PokemonType `json:"types"`
	}

	PokemonSprites struct {
		Front_default      string `json:"front_default"`
		Front_shiny        string `json:"front_shiny"`
		Front_female       string `json:"front_female"`
		Front_shiny_female string `json:"front_shiny_female"`
		Back_default       string `json:"back_default"`
		Back_shiny         string `json:"back_shiny"`
		Back_female        string `json:"back_female"`
		Back_shiny_female  string `json:"back_shiny_female"`
	}

	PokemonCries struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	}

	PokemonStat struct {
		Stat      NamedResource `json:"stat"`
		Effort    int           `json:"effort"`
		Base_stat int           `json:"base_stat"`
	}

	PokemonType struct {
		Slot int           `json:"slot"`
		Type NamedResource `json:"type"`
	}
)
