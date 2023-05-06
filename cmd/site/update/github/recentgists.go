package github

import (
	"context"
	"fmt"
	"time"

	github "github.com/shurcooL/githubv4"
)

type gistsQuery struct {
	User struct {
		Login github.String
		Gists struct {
			// TotalCount github.Int
			Edges []struct {
				Cursor github.String
				Node   qlGist
			}
		} `graphql:"gists(first: $count, after: $after, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"user(login:$username)"`
}

type Gist struct {
	Name        string
	Description string
	URL         string
	CreatedAt   time.Time
}

func (handler *Handler) GetRecentGists(ctx context.Context, count int) ([]Gist, error) {
	query := gistsQuery{}
	vars := Variables{
		"username": github.String(handler.username),
		"count":    github.Int(count),
	}

	var after *github.String
	gists := make([]Gist, 0, count)

	for {
		vars["after"] = after
		err := handler.client.Query(ctx, &query, vars)
		if err != nil {
			return nil, fmt.Errorf("failed to query recent gists: %w", err)
		}

		if len(query.User.Gists.Edges) == 0 {
			break
		}

		for _, v := range query.User.Gists.Edges {
			if !v.Node.IsPublic {
				continue
			}

			gists = append(gists, *v.Node.Unwrap())
			after = github.NewString(v.Cursor)
		}

		if len(gists) > count {
			break
		}
	}

	if len(gists) > count {
		return gists[:count], nil
	}

	return gists, nil
}
