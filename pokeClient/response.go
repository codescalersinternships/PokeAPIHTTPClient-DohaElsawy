package pokeclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/sirupsen/logrus"
)

// Time takes for backoff to wait
const (
	ConstBackoffTime = 4 * time.Second
)

// ErrResponse represent error response
type ErrResponse struct {
	Err        error `json:"error"`
	StatusCode int   `json:"statuscode"`
}

// Error implement custom error types
func (e ErrResponse) Error() string {
	return fmt.Sprintf("error is %s , and http status code is %d", e.Err, e.StatusCode)
}

// AssignErrorResponse takes error and status code and return an ErrResponse type
func AssignErrorResponse(err error, statuscode int) error {
	return ErrResponse{
		Err:        err,
		StatusCode: statuscode,
	}
}

// GetResponse retrive data and error if exist
func (c *Client) GetResponse(responseType interface{}) error {

	req, err := http.NewRequest("GET", c.Url, nil)

	if err != nil {
		logrus.Errorf("from http.NewRequestWithContext function %s\n", err)
		return AssignErrorResponse(err, req.Response.StatusCode)
	}

	operation := func() (*http.Response, error) {
		res, err := c.Client.Do(req)
		return res, err
	}

	notify := func(err error, t time.Duration) {
		logrus.Printf("error: %v happened at time: %v", err, t)
	}

	b := backoff.NewConstantBackOff(ConstBackoffTime)

	res, err := backoff.RetryNotifyWithData(operation, b, notify)

	if err != nil {
		logrus.Errorf("from Client.Do function %s\n", err)
		return AssignErrorResponse(err, res.StatusCode)
	}
	defer res.Body.Close()

	// response part ------------

	if res.StatusCode != http.StatusOK {
		logrus.Errorf("Response status is not 200, it is %v", res.StatusCode)
		return AssignErrorResponse(fmt.Errorf("unsupported header format"), res.StatusCode)
	}

	decoder := json.NewDecoder(res.Body)

	err = decoder.Decode(&responseType)

	if err != nil {
		logrus.Errorf("unable to unmarchal body to json. Err = %s", err)
		return AssignErrorResponse(err, req.Response.StatusCode)
	}

	return nil
}
