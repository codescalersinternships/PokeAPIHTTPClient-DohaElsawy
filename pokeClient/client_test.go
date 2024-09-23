package pokeclient

import (
	"fmt"
	"reflect"
	"testing"
)

func TestClient(t *testing.T) {
	testcase := []struct {
		offset int
		limit  int
	}{
		{
			offset: 0,
			limit:  0,
		},
		{
			offset: 10,
			limit:  0,
		},
		{
			offset: 0,
			limit:  10,
		},
		{
			offset: 10,
			limit:  10,
		},
	}

	for _, test := range testcase {
		t.Run("initiate a valid client", func(t *testing.T) {

			expectURL := fmt.Sprintf("https://pokeapi.co/api/v2/gender?offset=%d&limit=%d", test.offset, test.limit)

			result := NewClient("/gender", test.offset, test.limit)

			if !reflect.DeepEqual(expectURL, result.Url) {
				t.Errorf("i got %v, expect %v", result.Url, expectURL)
			}
		})
	}

	for _, test := range testcase {
		t.Run("load endpoint from .env and initiate valid client", func(t *testing.T) {

			endpoint, err := LoadConfigFromENV("../testdata/.env")

			if err != nil {
				t.Errorf("the error is %v", err)
			}
			expectURL := fmt.Sprintf("https://pokeapi.co/api/v2%s?offset=%d&limit=%d", endpoint, test.offset, test.limit)

			result := NewClient(endpoint, test.offset, test.limit)

			if !reflect.DeepEqual(expectURL, result.Url) {
				t.Errorf("i got %v, expect %v", result.Url, expectURL)
			}
		})
	}

}
