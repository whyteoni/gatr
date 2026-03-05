package main

import (
	"github.com/whyteoni/gatr/internal/config"
	"github.com/whyteoni/gatr/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type CliCommand struct {
	Name     string
	Desc     string
	Callback func(state, []string) error
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
