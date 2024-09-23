package pokeclient

import "fmt"

type Pockemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	IsDefault      bool   `json:"is_default"`
	Order          int    `json:"order"`
	Weight         int    `json:"weight"`
	Abilities      []struct {
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
		Ability  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
}

func GetPokemon(id int) (result Pockemon, err error) {

	endpoint := fmt.Sprintf("/pokemon/%d", id)

	c := NewClient(endpoint)

	err = c.GetResponse(&result)

	if err != nil {
		return Pockemon{}, err
	}
	return result, nil
}
