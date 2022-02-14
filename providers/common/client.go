package common

import (
	"context"
	"github.com/google/go-github/github"
	"github.com/hashicorp-demoapp/hashicups-client-go"
	"golang.org/x/oauth2"
)

type Client struct{}

func (c *Client) GitHubNewClient(pat *string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: *pat},
	)
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc)
}

func (c *Client) HashiCupsClient(host, username, password *string) (*hashicups.Client, error) {
	newClient, err := hashicups.NewClient(host, username, password)
	if err != nil {
		return nil, err
	}
	return newClient, nil
}
