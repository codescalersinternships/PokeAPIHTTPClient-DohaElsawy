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
		err ErrResponse
	}{
		{
			offset:         10,
			limit:          5,
			endpoint:       "/pokemon",
			expectName:     "metapod",
			expectedLength: 5,
			err: ErrResponse{},
		},
		{
			offset:         0,
			limit:          0,
			endpoint:       "/pokemon",
			expectName:     "bulbasaur",
			expectedLength: 20,
			err: ErrResponse{},
		},
		{
			offset:         0,
			limit:          -1,
			endpoint:       "/ability",
			expectName:     "stench",
			expectedLength: 366,
			err: ErrResponse{},
		},
		{
			offset:         0,
			limit:          -1,
			endpoint:       "/ability",
			expectName:     "stench",
			expectedLength: 366,
			err: ErrResponse{},
		},
		
	}

	for _, test := range testcase {
		t.Run("valid pokemon list resource", func(t *testing.T) {

			var recource Resource

			c := NewClient(test.endpoint, test.offset, test.limit)

			err := c.GetResponse(&recource)

			assertIsEqual(t, test.expectedLength, len(recource.Results))
			assertIsEqual(t, recource.Results[0].Name, test.expectName)
			assertIsEqual(t, nil, err)

			assertNotEqual(t, 0, recource.Count)
		})
	}
}


