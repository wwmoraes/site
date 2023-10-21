package github

import (
	"context"

	github "github.com/shurcooL/githubv4"
)

type recentRepositoriesQuery struct {
	User struct {
		// Login        github.String
		Repositories struct {
			// TotalCount github.Int
			Edges []struct {
				// Cursor github.String
				Node qlRepository
			}
		} `graphql:"repositories(first: $count, privacy: PUBLIC, isFork: $isFork, ownerAffiliations: OWNER, orderBy: {field: CREATED_AT, direction: DESC})"` //nolint:lll
	} `graphql:"user(login:$username)"`
}

type Repository struct {
	NameWithOwner string
	URL           string
	Description   string
	Stargazers    int
}

func (handler *Handler) GetRecentRepositories(ctx context.Context, count int) ([]Repository, error) {
	query := recentRepositoriesQuery{}
	vars := Variables{
		"username": github.String(handler.username),
		"count":    github.Int(count + 1),
		"isFork":   github.Boolean(false),
	}

	err := handler.client.Query(ctx, &query, vars)
	if err != nil {
		return nil, err
	}

	repositories := make([]Repository, len(query.User.Repositories.Edges))
	for index, repository := range query.User.Repositories.Edges {
		repositories[index] = *repository.Node.Unwrap()
	}

	if len(repositories) > count {
		return repositories[:count], nil
	}

	return repositories, nil
}
