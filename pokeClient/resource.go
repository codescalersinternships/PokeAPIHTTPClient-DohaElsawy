package pokeclient



// Resource represent the return resource resopnse
type Resource struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetResource(endpoint string, params ...int) (result Resource, err error) {
	c := NewClient(endpoint,params...)
	err = c.GetResponse(&result)

	if err != nil {
		return Resource{} , err
	}
	return result , nil
}


