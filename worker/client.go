package main

import (
	"net/http"
	"time"
)

type Client struct {
	BaseUrl    string
	HTTPClient *http.Client
}

func NewClient() *Client {
	c := &Client{
		BaseUrl: "https://restcountries.com/v3.1",
		HTTPClient: &http.Client{
			Timeout: 5 * time.Minute,
		},
	}
	return c
}
