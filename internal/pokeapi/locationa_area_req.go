package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	var fullURL, endpoint string
	if pageURL == nil {
		endpoint = "/location"
		fullURL = baseUrl + endpoint
	} else {
		fullURL = *pageURL
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationAreasResp := LocationAreasResp{}
	err = json.Unmarshal(body, &locationAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	return locationAreasResp, nil
}
