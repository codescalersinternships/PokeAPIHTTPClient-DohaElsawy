package pokeclient

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Client represent a http client
type Client struct {
	Client *http.Client
	Url    string
}

// Generic http client for default values
var (
	TimeoutDefault = 2 * time.Second
	UrlDefault     = "https://pokeapi.co/api/v2"
)

// LoadConfigFromENV load endpoint from env file and return error if exist
func LoadConfigFromENV(path string) (string, error) {

	err := godotenv.Load(path)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return "", err
	}

	endpoint := os.Getenv("ENDPOINT")

	logrus.Println("create client with load env file")

	return endpoint, nil
}

// NewClient initalize new http client and take endpoint
func NewClient(endpoint string, params ...int) *Client {

	offset, limit := parseParams(params)
	client := &Client{
		Client: &http.Client{
			Timeout: TimeoutDefault,
		},
		Url: fmt.Sprintf("%s%s?offset=%d&limit=%d", UrlDefault, endpoint, offset, limit),
	}

	logrus.Printf("new client created %v\n", client)

	return client
}

func parseParams(params []int) (int, int) {
	size := len(params)

	if size == 2 {
		return params[0], params[1]
	}
	if size == 1 {
		return params[0], params[1]
	}
	return 0, 0
}
