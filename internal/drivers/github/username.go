package github

import (
	"context"

	github "github.com/shurcooL/githubv4"
)

type viewerQuery struct {
	Viewer struct {
		Login github.String
	}
}

func getUsername(ctx context.Context, client *github.Client) (string, error) {
	var viewer viewerQuery

	err := client.Query(ctx, &viewer, nil)
	if err != nil {
		return "", err
	}

	return string(viewer.Viewer.Login), nil
}
