package pokeclient

import (
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	t.Run("initiate a valid client",func(t *testing.T) {

		expectURL := "https://pokeapi.co/api/v2/gender"

		result := NewClient("/gender")

		if !reflect.DeepEqual(expectURL,result.Url) {
			t.Errorf("%v",result)
		}
	})

	t.Run("load endpoint from .env and initiate valid client",func(t *testing.T) {

		expectURL := "https://pokeapi.co/api/v2/characters"

		endpoint , err := LoadConfigFromENV("../.env")

		if err != nil {
			t.Errorf("the error is %v", err)
		}

		result := NewClient(endpoint)


		if !reflect.DeepEqual(expectURL,result.Url) {
			t.Errorf("%v",result)
		}
	})

}
