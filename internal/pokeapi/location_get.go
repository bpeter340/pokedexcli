package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationAreaName string) (Location, error) {
	url := BaseURL + "/location-area/" + locationAreaName

	if data, exists := c.cache.Get(url); exists {
		var responseData Location
		err := json.Unmarshal(data, &responseData)
		if err != nil {
			return Location{}, err
		}
		return responseData, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return Location{}, errors.New("Location not found")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	var responseData Location
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return responseData, nil
}
