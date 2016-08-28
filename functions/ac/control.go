package main

import (
	"context"
	"fmt"
	"os"

	irkit "github.com/tcnksm/go-irkit/v1"
)

// handleControl handles TurnON/TurnOFF requests
// It sends AC on/off signal via IRKit internet HTTP API
func handleControl(d *Directive) (*Directive, error) {

	irClientKey := os.Getenv(EnvIRClientKey)
	if len(irClientKey) == 0 {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("missing %q env var", EnvIRClientKey)
	}

	irDeviceId := os.Getenv(EnvIRDeviceID)
	if len(irDeviceId) == 0 {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("missing %q env var", EnvIRDeviceID)
	}

	// Construct IRKit internet client
	client := irkit.DefaultInternetClient()
	switch name := d.Header.Name; name {
	case TurnOnRequest:
		err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgACOn)
		if err != nil {
			res := ErrorResponse(DriverInternalError, "", "")
			return res, err
		}

		// Change header name
		d.Header.Name = TurnOnConfirmation

		return &Directive{
			Header:  d.Header,
			Payload: map[string]interface{}{},
		}, nil
	case TurnOffRequest:
		err := client.SendMessages(context.TODO(), irClientKey, irDeviceId, &IRKitMsgACOff)
		if err != nil {
			res := ErrorResponse(DriverInternalError, "", "")
			return res, err
		}

		// Change header name
		d.Header.Name = TurnOffConfirmation

		return &Directive{
			Header:  d.Header,
			Payload: map[string]interface{}{},
		}, nil
	default:
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("unexpected name: %s", name)
	}
}
