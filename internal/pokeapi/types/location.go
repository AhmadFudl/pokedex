package types

type (
	LocationArea struct {
		Id                     int                   `json:"id"`
		Name                   string                `json:"name"`
		Game_index             int                   `json:"game_index"`
		Encounter_method_rates []EncounterMethodRate `json:"encounter_method_rates"`
		Location               NamedResource         `json:"location"`
		Names                  []Name                `json:"names"`
		Pokemon_encounters     []PokemonEncounter    `json:"pokemon_encounters"`
	}

	PokemonEncounter struct {
		Pokemon         NamedResource            `json:"pokemon"`
		Version_details []VersionEncounterDetail `json:"version_details"`
	}
)
