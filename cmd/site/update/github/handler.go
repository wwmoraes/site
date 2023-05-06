package github

import (
	"context"
	"net/http"
	"os"

	github "github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Handler struct {
	client   *github.Client
	username string
}

func NewHandler(ctx context.Context) (*Handler, error) {
	var httpClient *http.Client

	gitHubToken := os.Getenv("GITHUB_TOKEN")
	if len(gitHubToken) > 0 {
		httpClient = oauth2.NewClient(ctx, oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: gitHubToken},
		))
	}

	client := github.NewClient(httpClient)
	username, err := getUsername(ctx, client)
	if err != nil {
		return nil, err
	}

	return &Handler{client, username}, nil
}
