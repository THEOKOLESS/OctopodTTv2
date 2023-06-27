package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type TranslatedCountryName struct {
	Fra struct {
		Official string `json:"official"`
	} `json:"fra"`
}

type CountryData struct {
	Name       TranslatedCountryName `json:"translations"`
	Flag       string                `json:"flag"`
	Population int                   `json:"population"`
}

func GetCountries(c *Client) ([]CountryData, error) {
	b, err := c.SendRequest("GET", "/all")
	if err != nil {
		fmt.Println("Error while connecting to API:", err)
		return nil, err
	}
	var countries []CountryData
	err = json.Unmarshal(b, &countries)
	if err != nil {
		log.Fatalf("JSON Unmarshal error: %v", err)
		return nil, err
	}

	return countries, nil
}
