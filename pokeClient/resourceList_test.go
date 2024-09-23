package pokeclient

import (
	"reflect"
	"testing"
)

func assertIsEqual(t *testing.T, type1, type2 any) {
	t.Helper()

	if !reflect.DeepEqual(type1, type2) {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

func assertNotEqual(t *testing.T, type1, type2 any) {
	t.Helper()

	if type1 == type2 {
		t.Errorf("i expect %v, found %v", type1, type2)
	}
}

func TestResource(t *testing.T) {
	testcase := []struct {
		offset         int
		limit          int
		endpoint       string
		expectName     string
		expectedLength int
		err            ErrResponse
	}{
		{
			offset:         10,
			limit:          5,
			endpoint:       "/pokemon",
			expectName:     "metapod",
			expectedLength: 5,
		},
		{
			offset:         0,
			limit:          0,
			endpoint:       "/pokemon",
			expectName:     "bulbasaur",
			expectedLength: 20,
		},
		{
			offset:         0,
			limit:          -1,
			endpoint:       "/ability",
			expectName:     "stench",
			expectedLength: 366,
		},
		{
			offset:         0,
			limit:          -1,
			endpoint:       "/ability",
			expectName:     "stench",
			expectedLength: 366,
		},
	}

	for _, test := range testcase {
		t.Run("valid list resource", func(t *testing.T) {

			resource, err := GetResource(test.endpoint, test.offset, test.limit)

			assertIsEqual(t, test.expectedLength, len(resource.Results))
			assertIsEqual(t, test.expectName, resource.Results[0].Name)
			assertIsEqual(t, nil, err)

			assertNotEqual(t, 0, resource.Count)
		})
	}


	t.Run("valid list resource load from .env", func(t *testing.T) {

		endpoint , params , err := LoadConfigFromENV("../testdata/.env")
		
		assertIsEqual(t,nil,err)

		resource, err := GetResource(endpoint, params[0], params[1])

		assertIsEqual(t,params[1] , len(resource.Results))
		assertIsEqual(t,"charmeleon" , resource.Results[0].Name)
		assertIsEqual(t, nil, err)

		assertNotEqual(t, 0, resource.Count)
	})

}
