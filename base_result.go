package lastfm

import "encoding/xml"

// Base is a base result struct.
type Base struct {
	XMLName xml.Name `xml:"lfm"`
	Status  string   `xml:"status,attr"`
	Inner   []byte   `xml:",innerxml"`
}

// APIError is a base API error struct.
type APIError struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}
