package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PaginatedResponse, error) {
	url := BaseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.cache.Get(url); exists {
		var locationResponse PaginatedResponse
		err := json.Unmarshal(data, &locationResponse)
		if err != nil {
			return PaginatedResponse{}, err
		}
		return locationResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PaginatedResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PaginatedResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PaginatedResponse{}, err
	}

	var locationsResponse PaginatedResponse

	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return PaginatedResponse{}, err
	}

	c.cache.Add(url, data)

	return locationsResponse, nil
}
