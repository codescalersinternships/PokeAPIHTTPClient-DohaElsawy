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

}
