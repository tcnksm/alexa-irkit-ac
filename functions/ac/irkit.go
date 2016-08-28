package main

import (
	"encoding/json"

	irkit "github.com/tcnksm/go-irkit/v1"
)

const (
	// IRKit related environmental var
	EnvIRClientKey = "IR_CLIENT_KEY"
	EnvIRDeviceID  = "IR_DEVICE_ID"
)

var (
	IRKitMsgACOn  irkit.Message
	IRKitMsgACOff irkit.Message
)

// readSigal reads IRKit signal json data from ./signals directory.
// Files should be transformed into go bindary data by go-bindata.
func readSignal(path string, msg *irkit.Message) error {
	data, err := Asset(path)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, msg)
}

func init() {
	err := readSignal("signals/aircon-on.json", &IRKitMsgACOn)
	if err != nil {
		panic(err)
	}

	err = readSignal("signals/aircon-off.json", &IRKitMsgACOff)
	if err != nil {
		panic(err)
	}
}
