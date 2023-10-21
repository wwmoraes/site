package github

import (
	"context"
	"fmt"
	"time"

	github "github.com/shurcooL/githubv4"
)

type recentPullRequestsQuery struct {
	User struct {
		Login        github.String
		PullRequests struct {
			// TotalCount github.Int
			Edges []struct {
				// Cursor github.String
				Node qlPullRequest
			}
		} `graphql:"pullRequests(first: $count, orderBy: {field: CREATED_AT, direction: DESC})"`
	} `graphql:"user(login:$username)"`
}

type PullRequest struct {
	Repository `json:"Repository"` //nolint:tagliatelle

	URL       string
	Title     string
	State     string
	CreatedAt time.Time
}

func (handler *Handler) GetRecentPullRequests(ctx context.Context, count int) ([]PullRequest, error) {
	query := recentPullRequestsQuery{}
	vars := Variables{
		"username": github.String(handler.username),
		"count":    github.Int(count + 1),
	}

	err := handler.client.Query(ctx, &query, vars)
	if err != nil {
		return nil, err
	}

	pullRequests := make([]PullRequest, 0, count)

	for _, v := range query.User.PullRequests.Edges {
		if string(v.Node.Repository.NameWithOwner) == fmt.Sprintf("%s/%s", handler.username, handler.username) {
			continue
		}

		if v.Node.Repository.IsPrivate {
			continue
		}

		pullRequests = append(pullRequests, *v.Node.Unwrap())
		if len(pullRequests) == count {
			break
		}
	}

	return pullRequests, nil
}
