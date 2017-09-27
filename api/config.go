package main

import (
	"errors"
	"os"
)

// Config map containing  configuration settings
func verifyConfig() error {
	if os.Getenv("PORT") == "" {
		os.Setenv("PORT", ":5000")
	}

	if os.Getenv("TRACK") == "" {
		return errors.New("Missing String to Track on twitter")
	}

	if os.Getenv("CONSUMER-KEY") == "" || os.Getenv("CONSUMER-SECRET") == "" {
		return errors.New("Consumer key/secret required")
	}
	if os.Getenv("ACCESS-TOKEN") == "" || os.Getenv("ACCESS-SECRET") == "" {
		return errors.New("Access token/secret required")
	}

	return nil
}
