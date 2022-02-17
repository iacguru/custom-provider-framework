package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
)

func GitHubNewClient(pat string) *Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: pat},
	)
	return &Client{
		HttpClient: oauth2.NewClient(ctx, ts),
	}
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
