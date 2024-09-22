package client

import (
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Client represent a http client
type Client struct {
	Client *http.Client
	Endpoint    string
}

// Generic http client for default values
var (
	TimeoutDefault  = 2 * time.Second
	UrlDefault      = "https://pokeapi.co/api/v2"
)


// LoadConfigFromENV load endpoint from env file and return error if exist
func (c *Client) LoadConfigFromENV(path string) error {

	err := godotenv.Load(path)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return err
	}

	endpoint := os.Getenv("ENDPOINT")
	
	tempCLient := NewClient(endpoint)

	c.Client = tempCLient.Client
	c.Endpoint = tempCLient.Endpoint

	logrus.Println("create client with load env file")

	return nil
}

// NewClient initalize new http client and take option url string
func NewClient(endpoint string) *Client {

	client := &Client{
		Client: &http.Client{
			Timeout: TimeoutDefault,
		},
		Endpoint: UrlDefault+endpoint,
	}

	logrus.Printf("new client created %v\n", client)

	return client
}
