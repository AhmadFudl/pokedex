package types

type (
	EncounterMethodRate struct {
		Encounter_method NamedResource            `json:"encounter_method"`
		Version_details  []EncounterVersionDetail `json:"version_details"`
	}

	EncounterVersionDetail struct {
		Rate    int           `json:"rate"`
		Version NamedResource `json:"version"`
	}

	Encounter struct {
		Min_level        int             `json:"min_level"`
		Max_level        int             `json:"max_level"`
		Condition_values []NamedResource `json:"condition_values"`
		Chance           int             `json:"chance"`
		Method           NamedResource   `json:"method"`
	}
)
