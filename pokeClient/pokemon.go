package pokeclient

import (
	"fmt"

	"github.com/sirupsen/logrus"
)


// Pockemon represent pokemon structure
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


// GetPokemon returns pokemon object and error if exist and takes pokemon id 
func GetPokemon(id int) (result Pockemon, err error) {

	endpoint := fmt.Sprintf("/pokemon/%d", id)

	c := NewClient(endpoint)

	err = c.GetResponse(&result)

	if err != nil {

		logrus.Printf("error in getting pokemon, error %v\n", err)

		return Pockemon{}, err
	}

	logrus.Printf("successfully getting pokemon")

	return result, nil
}
