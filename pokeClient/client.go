package pokeclient

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
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
func LoadConfigFromENV(path string) (string,[]int, error) {

	err := godotenv.Load(path)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return "", nil ,err
	}

	endpoint := os.Getenv("ENDPOINT")
	offset := os.Getenv("OFFSET")
	limit := os.Getenv("LIMIT")

	params , err := appendParams(offset,limit)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return "", nil ,err
	}

	
	logrus.Println("create client with load env file")

	return endpoint,params ,nil
}

// NewClient initalize new http client and take endpoint and offset and limit as options 
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

func appendParams(offset , limit string) ([]int , error){
	var params []int

	offsetInt , err := strconv.Atoi(offset)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return nil ,err
	}
	limitInt , err := strconv.Atoi(limit)

	if err != nil {
		logrus.Errorf("error while loading .env file. Err: %s", err)
		return nil, err
	}


	params = append(params, offsetInt)
	params = append(params, limitInt)
	return params , nil
}