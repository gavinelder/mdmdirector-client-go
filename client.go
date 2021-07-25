package mdmdirector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HostURL - Default MdmDirector URL
const HostURL string = "http://localhost:8000"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Username   string
	Password   string
}

// NewClient -
func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		Username:   *username,
		Password:   *password,
	}
	if host != nil {
		c.HostURL = *host
	}
	if c.Username == "" || c.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
