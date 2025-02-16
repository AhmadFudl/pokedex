package types

type (
	NamedResourceList struct {
		Count    int             `json:"count"`
		Next     string          `json:"next"`
		Previous string          `json:"previous"`
		Results  []NamedResource `json:"results"`
	}

	NamedResource struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
)
