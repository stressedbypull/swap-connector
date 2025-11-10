package swapi

// DTOs mirror SWAPI JSON. Use string types for fields that the API returns as strings
// so we can control parsing in the mapper.

type PersonDTO struct {
	Name    string   `json:"name"`
	Mass    string   `json:"mass"`
	Created string   `json:"created"`
	Films   []string `json:"films"`
}

type PlanetDTO struct {
	Name      string   `json:"name"`
	Residents []string `json:"residents"`
	Created   string   `json:"created"`
	Films     []string `json:"films"`
}

type SWAPIPeopleResponse struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []PersonDTO `json:"results"`
}

type SWAPIPlanetsResponse struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []PlanetDTO `json:"results"`
}
