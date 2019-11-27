package geonames

import "net/http"

// Client is the main entry point for the geonames library
type Client struct {
	c http.Client
}
