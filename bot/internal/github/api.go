package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
	"golang.org/x/oauth2"
)

type api struct {
	client *github.Client
}

func (api *api) getUserClient(ctx context.Context, user string) (*github.Client, error) {
	installation, _, err := api.client.Apps.FindUserInstallation(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("find user installation: %w", err)
	}

	// TODO: store client to cache

	return api.getClient(ctx, *installation.ID)
}

func (api *api) getOrganizationClient(ctx context.Context, org string) (*github.Client, error) {
	installation, _, err := api.client.Apps.FindOrganizationInstallation(ctx, org)
	if err != nil {
		return nil, fmt.Errorf("find organization installation: %w", err)
	}

	// TODO: store client to cache

	return api.getClient(ctx, *installation.ID)

}
func (api *api) getClient(ctx context.Context, installationId int64) (*github.Client, error) {
	token, _, err := api.client.Apps.CreateInstallationToken(ctx, installationId, &github.InstallationTokenOptions{})
	if err != nil {
		return nil, fmt.Errorf("create installation token: %w", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.GetToken()},
	)
	oAuthClient := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(oAuthClient)

	return client, nil
}
