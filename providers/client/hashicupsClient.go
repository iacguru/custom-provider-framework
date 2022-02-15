package client

import "github.com/hashicorp-demoapp/hashicups-client-go"

func (c *Client) HashiCupsClient(host, username, password *string) (*hashicups.Client, error) {
	newClient, err := hashicups.NewClient(host, username, password)
	if err != nil {
		return nil, err
	}
	return newClient, nil
}
