package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/apex/go-apex"
)

const (
	EnvAccessToken = "ACCESS_TOKEN"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		logger := log.New(os.Stderr, "", log.LstdFlags)
		logger.Printf("[INFO] Start ac function")

		accessToken := os.Getenv(EnvAccessToken)
		if len(accessToken) == 0 {
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("missing %q env var", EnvAccessToken)
		}

		var d Directive
		if err := json.Unmarshal(event, &d); err != nil {
			logger.Printf("[ERROR] Failed to unmarshal event: %s", err)
			res := ErrorResponse(DriverInternalError, "", "")
			return res, err
		}

		token, err := d.AccessToken()
		if err != nil {
			logger.Printf("[ERROR] Invalid request: %s", err)
			res := ErrorResponse(
				DriverInternalError,
				d.Header.MessageID,
				d.Header.NameSpace)
			return res, err
		}

		// Check access token is valid or not
		if token != accessToken {
			logger.Printf("[ERROR] Invalid access token")
			res := ErrorResponse(
				InvalidAccessTokenError,
				d.Header.MessageID,
				d.Header.NameSpace)
			return res, err
		}

		switch namespace := d.Header.NameSpace; namespace {
		case AlexaConnectedHomeDiscovery:
			logger.Printf("[INFO] Handle discovery request: %s", d.Header.Name)
			return handleDiscovery(&d)
		case AlexaConnectedHomeControl:
			logger.Printf("[INFO] Handle control request: %s", d.Header.Name)
			return handleControl(&d)
		default:
			logger.Printf("[ERROR] Unexpected Namespace: %s", namespace)
			res := ErrorResponse(DriverInternalError, "", "")
			return res, fmt.Errorf("unexpected namespace")
		}
	})
}
