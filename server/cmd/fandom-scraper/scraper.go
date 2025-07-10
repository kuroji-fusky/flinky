package main

import (
	"flag"
	"fmt"
	"time"
)

// Adds a N second delay so I don't get my ass rate limited
func delayRequest(delaySec int) {
	time.Sleep(time.Duration(delaySec) * time.Second)
}

func defineBoolFlags(ptr *bool, flags []string, defaultVal bool, help string) {
	for _, f := range flags {
		flag.BoolVar(ptr, f, defaultVal, help)
	}
}

func defineIntFlags(ptr *int, flags []string, defaultVal int, help string) {
	for _, f := range flags {
		flag.IntVar(ptr, f, defaultVal, help)
	}
}

func defineStringFlags(ptr *string, flags []string, defaultVal string, help string) {
	for _, f := range flags {
		flag.StringVar(ptr, f, defaultVal, help)
	}
}

func main() {
	// Flag stuff

	var flagReqDelay int
	var flagReqTimeout int
	var flagReqRetries int
	var flagVerboseMode bool
	var flagUseDatabase string

	defineIntFlags(&flagReqDelay, []string{"delay", "d"}, 4, "Delay between requests in seconds, default is 4")
	defineIntFlags(&flagReqTimeout, []string{"timeout", "T"}, 15, "Request timeout in seconds, default is 15")
	defineIntFlags(&flagReqRetries, []string{"retries", "r"}, 3, "Number of retries for failed requests, default is 3")

	defineBoolFlags(&flagVerboseMode, []string{"verbose", "v"}, false, "Enable verbose mode for debugging")
	defineStringFlags(&flagUseDatabase, []string{"use-db", "db"}, "", "Enable database usage")

	flag.Parse()

	if flagVerboseMode {
		fmt.Println("Verbose mode enabled")
	}

	// WIP rewriting all the scraper logic from Python so there's literally nothing here yet lmao
	ReadConfig()

	delayRequest(flagReqDelay)
}
