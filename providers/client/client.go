package client

import (
	"context"
	"github.com/hashicorp-demoapp/hashicups-client-go"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
	Method     string
	URL        string
	Body       io.Reader
	Headers    *map[string]string
	Req        *http.Request
}

func (c *Client) GitHubNewClient(pat string) *http.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	c.HttpClient = oauth2.NewClient(ctx, ts)
	return c.HttpClient
}

func (c *Client) HashiCupsClient(host, username, password *string) (*hashicups.Client, error) {
	newClient, err := hashicups.NewClient(host, username, password)
	if err != nil {
		return nil, err
	}
	return newClient, nil
}

func (c *Client) NewRequest() (*http.Request, error) {
	var err error
	c.Req, err = http.NewRequest(c.Method, c.URL, c.Body)
	if err != nil {
		return nil, err
	}
	return c.Req, nil
}

func (c *Client) DoRequest() (*http.Response, error) {
	client := c.HttpClient
	res, err := client.Do(c.Req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
