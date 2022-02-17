package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent {
		err = json.Unmarshal(body, &c.Workflows)
		if err != nil {
			fmt.Println("Error during do unmarshel:")
			return err
		}
	} else {
		return fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return nil
}
