package main_test

import (
	"os"
	"testing"
)

func TestReadConfig(t *testing.T) {
	tempFile, err := os.CreateTemp("", "scraper-config.temp.yml")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	defer os.Remove(tempFile.Name())
}
