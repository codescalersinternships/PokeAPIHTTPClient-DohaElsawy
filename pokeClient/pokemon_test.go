package pokeclient

import (
	"reflect"
	"testing"
)

func isEqual(t *testing.T, type1, type2 any) {
	t.Helper()

	if !reflect.DeepEqual(type1, type2) {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

func TestPokemon(t *testing.T) {
	testcase := []struct {
		id                    int
		expectAbilityName     string
		expectedPokemonHeight int
		err                   ErrResponse
	}{
		{
			id:                    30,
			expectAbilityName:     "poison-point",
			expectedPokemonHeight: 8,
			err:                   ErrResponse{},
		},
		{
			id:                    1,
			expectAbilityName:     "overgrow",
			expectedPokemonHeight: 7,
			err:                   ErrResponse{},
		},
	}

	for _, test := range testcase {
		t.Run("valid pokemon", func(t *testing.T) {

			pokemon, err := GetPokemon(test.id)

			isEqual(t, test.expectedPokemonHeight, pokemon.Height)
			isEqual(t, test.expectAbilityName, pokemon.Abilities[0].Ability.Name)
			isEqual(t, nil, err)

		})
	}

	t.Run("invalid pokemon", func(t *testing.T) {

		_, err := GetPokemon(-1)

		if err == nil {
			t.Errorf("should have error got nil")
		}
	})
}
