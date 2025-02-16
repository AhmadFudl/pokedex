package types

type (
	VersionEncounterDetail struct {
		Version           NamedResource `json:"version"`
		Max_chance        int           `json:"max_chance"`
		Encounter_details []Encounter   `json:"encounter_details"`
	}

	VersionGameIndex struct {
		Game_index int           `json:"game_index"`
		Version    NamedResource `json:"version"`
	}

	Name struct {
		Name     string        `json:"name"`
		Language NamedResource `json:"language"`
	}
)
