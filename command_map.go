package main

import (
	"errors"
	"fmt"
	"internal/pokeapi"
)

func commandMap(cfg *config) error {
	locationAreasResp, err := pokeapi.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	locationAreas := locationAreasResp.Results
	for _, location := range locationAreas {
		fmt.Println(location.Name)
	}

	cfg.previousLocationAreaURL = locationAreasResp.Previous
	cfg.nextLocationAreaURL = locationAreasResp.Next

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreasResp, err := pokeapi.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		return err
	}

	locationAreas := locationAreasResp.Results
	for _, location := range locationAreas {
		fmt.Println(location.Name)
	}

	cfg.previousLocationAreaURL = locationAreasResp.Previous
	cfg.nextLocationAreaURL = locationAreasResp.Next

	return nil
}
