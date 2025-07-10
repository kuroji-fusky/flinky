package main

import "os"

type ScraperConfig struct {
	Allowlist     []string `yaml:"allowlist"`
	IgnoreFilters []struct {
		Type  string `yaml:"type"`
		Value string `yaml:"value"`
	} `yaml:"ignore_filters"`
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
