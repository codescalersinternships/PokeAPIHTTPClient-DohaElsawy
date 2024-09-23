package pokeclient

import "github.com/sirupsen/logrus"

// Resource represent the return resource list resopnse
type Resource struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


// GetResource return resource list and error if found and takes endpoint and (offset, limit) as options 
func GetResource(endpoint string, params ...int) (result Resource, err error) {
	
	c := NewClient(endpoint, params...)

	err = c.GetResponse(&result)

	if err != nil {

		logrus.Printf("error in getting resource list, error %v\n", err)

		return Resource{}, err
	}

	logrus.Printf("successfully getting resource")

	return result, nil
}
