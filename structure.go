package main

type Config struct {
	Address   string `xml:"Address"`
	WUSHost   string `xml:"WUSHost"`
	DebugMode bool   `xml:"DebugMode"`
}
