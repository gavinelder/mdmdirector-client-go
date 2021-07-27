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
