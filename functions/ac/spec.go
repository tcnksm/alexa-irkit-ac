package main

import "fmt"

// spec.go defines Alexa Smart Home Skill request and response.
// Currently, it only defines what **need** (so apparently, API
// is not completed).
//
// See more at Alexa skill Kit documentation.
// https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/smart-home-skill-api-reference

type NameSpace string

const (
	AlexaConnectedHomeDiscovery NameSpace = "Alexa.ConnectedHome.Discovery"
	AlexaConnectedHomeControl   NameSpace = "Alexa.ConnectedHome.Control"
)

type Name string

const (
	DiscoverAppliancesRequest  Name = "DiscoverAppliancesRequest"
	DiscoverAppliancesResponse Name = "DiscoverAppliancesResponse"

	TurnOnRequest      Name = "TurnOnRequest"
	TurnOnConfirmation Name = "TurnOnConfirmation"

	TurnOffRequest      Name = "TurnOffRequest"
	TurnOffConfirmation Name = "TurnOffConfirmation"
)

const (
	DriverInternalError     = "DriverInternalError"
	InvalidAccessTokenError = "InvalidAccessTokenError"
)

const (
	PayloadVersion = "2"
)

type Directive struct {
	Header  Header                 `json:"header"`
	Payload map[string]interface{} `json:"payload"`
}

// Header has a set of expected fields that are the same across message types.
type Header struct {
	MessageID string    `json:"messageId"`
	Name      Name      `json:"name"`
	NameSpace NameSpace `json:"namespace"`
	Version   string    `json:"payloadVersion"`
}

type DiscoveredAppliance struct {
	ApplianceID         string   `json:"applianceId"`
	ManufacturerName    string   `json:"manufacturerName"`
	ModelName           string   `json:"modelName"`
	Version             string   `json:"version"`
	FriendlyName        string   `json:"friendlyName"`
	FriendlyDescription string   `json:"friendlyDescription"`
	IsReachable         bool     `json:"isReachable"`
	Actions             []string `json:"actions"`

	AditionalApplianceDetails map[string]string `json:"additionalApplianceDetails"`
}

func (d *Directive) AccessToken() (string, error) {
	payload := d.Payload
	v, ok := payload["accessToken"]
	if !ok {
		return "", fmt.Errorf("missing access token")
	}

	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("failed to type assertion")
	}

	return token, nil
}

func ErrorResponse(name Name, messageId string, namespace NameSpace) *Directive {
	return &Directive{
		Header: Header{
			MessageID: messageId,
			Name:      name,
			NameSpace: namespace,
			Version:   PayloadVersion,
		},

		Payload: map[string]interface{}{},
	}
}
