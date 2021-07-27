package mdmdirector

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gavinelder/mdmdirector-client-go/types"
)

// GetSharedProfiles are profiles that go on every device.
func (c *Client) GetSharedProfiles() (*[]types.SharedProfile, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/profile", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	sharedProfiles := []types.SharedProfile{}
	err = json.Unmarshal(body, &sharedProfiles)
	if err != nil {
		return nil, err
	}

	return &sharedProfiles, nil
}

// GetDeviceProfiles are profiles that are assigned to a device either by UDID or SerialNumber.
func (c *Client) GetDeviceProfiles(deviceID string) (*[]types.DeviceProfile, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/profile/%s", c.HostURL, deviceID), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	deviceProfiles := []types.DeviceProfile{}
	err = json.Unmarshal(body, &deviceProfiles)
	if err != nil {
		return nil, err
	}

	return &deviceProfiles, nil
}
