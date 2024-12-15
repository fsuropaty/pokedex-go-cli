package pokeapi

type RespPokemon struct {
	Abilities      []Abilities `json:"abilities"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Stats          []Stats     `json:"stats"`
	Types          []Types     `json:"types"`
	Weight         int         `json:"weight"`
}

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Abilities struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Types struct {
	Slot int  `json:"slot"`
	Type Type `json:"type"`
}
