package client

import (
	action "github.com/iacguru/custom-provider-framework/providers/data/github"
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
	Workflows  *action.Workflow
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
