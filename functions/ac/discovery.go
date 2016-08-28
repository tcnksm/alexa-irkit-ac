package main

import "fmt"

// handleDiscovery handles discovery appliances request.
// Currently, it only returns constanct value.
func handleDiscovery(d *Directive) (*Directive, error) {
	if name := d.Header.Name; name != DiscoverAppliancesRequest {
		res := ErrorResponse(DriverInternalError, "", "")
		return res, fmt.Errorf("unexpected name: %s", name)
	}

	// Change header name
	d.Header.Name = DiscoverAppliancesResponse

	// Construct response
	res := &Directive{
		Header: d.Header,
		Payload: map[string]interface{}{
			"discoveredAppliances": []DiscoveredAppliance{
				{
					ApplianceID:         "irkit-air-conditioner",
					ManufacturerName:    "tcnksm",
					ModelName:           "irkit-001",
					Version:             "0.0.1",
					FriendlyName:        "AC",
					FriendlyDescription: "Air Conditioner",
					IsReachable:         true,
					Actions: []string{
						"turnOn",
						"turnOff",
					},
					AditionalApplianceDetails: map[string]string{},
				},
			},
		},
	}

	return res, nil
}
