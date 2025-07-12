package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/gofrs/flock"
)

type CacheContents struct {
	filePath string
	file     *os.File
	lock     *flock.Flock
}

// The base cache function, creates a new cache file or opens an existing one
func NewCache(filePath string) (*CacheContents, error) {
	lock := flock.New(filePath + ".lock")

	isLocked, err := lock.TryLock()
	if err != nil {
		return nil, err
	}
	if !isLocked {
		return nil, fmt.Errorf("could not acquire lock on %s", filePath)
	}

	cacheFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		lock.Unlock()
		return nil, err
	}

	info, err := cacheFile.Stat()
	if err != nil {
		cacheFile.Close()
		lock.Unlock()
		return nil, err
	}

	if info.Size() == 0 {
		dateHeader := "last updated " + time.Now().Format(time.RFC3339) + "\n"
		if _, err := cacheFile.WriteString(dateHeader); err != nil {
			cacheFile.Close()
			lock.Unlock()
			return nil, err
		}
	}

	return &CacheContents{
		filePath: filePath,
		file:     cacheFile,
		lock:     lock,
	}, nil
}

// Returns the contents of the cache file
func (c *CacheContents) Read() ([]string, error) {
	c.file.Seek(0, 0)

	data, err := io.ReadAll(c.file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines, nil
}

// Appends whatever new data is passed to the cache file
func (c *CacheContents) Append(data string) error {
	_, err := c.file.WriteString(data + "\n")
	return err
}

// Nuke deletes the cache file and closes it
func (c *CacheContents) Nuke() error {
	if err := c.file.Close(); err != nil {
		return err
	}
	c.lock.Unlock()
	return os.Remove(c.filePath)
}

// Safely closes the cache file
func (c *CacheContents) Close() error {
	defer c.lock.Unlock()
	return c.file.Close()
}
