package client

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
)

func (c *Client) GitHubNewClient(pat string) *http.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	c.HttpClient = oauth2.NewClient(ctx, ts)
	return c.HttpClient
}

func (c *Client) GetGitWorkflows() error {
	c.NewRequest()
	res, err := c.DoRequest()
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &c.Workflows)
	defer res.Body.Close()
	if err != nil {
		fmt.Println("Error during do unmarshel:")
		return err
	}
	return nil
}
