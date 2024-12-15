package pokeapi

type RespLocationPokemon struct {
	ID                int                 `json:"id"`
	Location          Location            `json:"location"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}
type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}
