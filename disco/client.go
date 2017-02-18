// Package disco controls interaction with the rqlite Discovery service
package disco

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DiscoResponse represents the response returned by a Discovery Service.
type DiscoResponse struct {
	CreatedAt string   `json:"created_at"`
	DiscoID   string   `json:"disco_id"`
	Nodes     []string `json:"nodes"`
}

// Client provides a Discovery Service client.
type Client struct {
	url string
}

// New returns an initialized Discovery Service client.
func New(url string) *Client {
	return &Client{
		url: url,
	}
}

// URL returns the Discovery Service URL used by this client.
func (c *Client) URL() string {
	return c.url
}

// Register attempts to register with the Discovery Service, using the given
// address.
func (c *Client) Register(id, addr string) (*DiscoResponse, error) {
	m := map[string]string{
		"addr": addr,
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.url, id)
	resp, err := http.Post(url, "application-type/json", bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	disco := &DiscoResponse{}
	if err := json.Unmarshal(b, disco); err != nil {
		return nil, err
	}

	return disco, nil
}
