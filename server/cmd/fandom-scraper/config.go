package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

type IgnoreFilterTuple struct {
	ItemType, ItemVal string
}

type ScraperConfig struct {
	Allowlist     []string            `yaml:"allowlist"`
	IgnoreFilters []IgnoreFilterTuple `yaml:"ignore_filters"`
}

func ReadConfig() {
	cfgFile, err := os.Open("scraper-config.yml")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := cfgFile.Close(); err != nil {
			panic(err)
		}
	}()
}
