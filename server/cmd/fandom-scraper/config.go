package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

type (
	IgnoreFilterTuple struct {
		ItemType, ItemVal string
	}

	ScraperConfig struct {
		Allowlist     []string            `yaml:"allowlist"`
		IgnoreFilters []IgnoreFilterTuple `yaml:"ignore_filters"`
	}
)

func readYAMLFile(filePath string) (*ScraperConfig, error) {
	cfgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := cfgFile.Close(); err != nil {
			panic(err)
		}
	}()

	scraperCfg := ScraperConfig{}
	if err := yaml.NewDecoder(cfgFile).Decode(&scraperCfg); err != nil {
		return nil, err
	}

	return &scraperCfg, nil
}

func ReadConfig() {
	scraperCfg, err := readYAMLFile("scraper-config.yml")
	if err != nil {
		panic(err)
	}

	cfgAllowlist := scraperCfg.Allowlist
	cfgIgnoreFilters := scraperCfg.IgnoreFilters

	if len(cfgAllowlist) == 0 {
		panic("Allowlist cannot be empty")
	}

	if len(cfgIgnoreFilters) == 0 {
		panic("IgnoreFilters cannot be empty")
	}
}
