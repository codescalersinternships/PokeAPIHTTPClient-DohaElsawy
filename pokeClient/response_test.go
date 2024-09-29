package pokeclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetResponse(t *testing.T) {

	t.Run("Data with json format", func(t *testing.T) {
		mockserver := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Set("Content-type", "application/json")

				w.WriteHeader(http.StatusOK)
			}),
		)

		mockserver.URL = "/pokemon"

		var resource Resource
		client := NewClient(mockserver.URL)

		err := client.GetResponse(&resource)

		if err != nil {
			t.Errorf("error is %v", err)
		}

	})

	t.Run("invalid url", func(t *testing.T) {
		var resource Resource

		client := NewClient("/notfound")

		err := client.GetResponse(&resource)

		if err == nil {
			t.Errorf("expect error %v, found nil", err)
		}

	})

}
