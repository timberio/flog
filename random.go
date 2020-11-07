package main

import (
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

// RandTimeFormat returns a random time format
func RandTimeFormat() string {
	formats := []string{
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
		"2006-01-02T15:04:05.999Z",
		"02/Jan/2006:15:04:05 -0700",
		"2006/01/02 15:04:05",
		"2006-01-02T15:04:05.999999999Z",
	}

	return formats[rand.Intn(len(formats))]
}
