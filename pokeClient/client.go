package pokeclient

import (
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Client represent a http client
type Client struct {
	Client   *http.Client
	Url string
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

	return endpoint , nil
}

// NewClient initalize new http client and take endpoint
func NewClient(endpoint string) *Client {

	client := &Client{
		Client: &http.Client{
			Timeout: TimeoutDefault,
		},
		Url: UrlDefault + endpoint,
	}

	logrus.Printf("new client created %v\n", client)

	return client
}
