package github

import (
	"context"
	"sort"
	"time"

	github "github.com/shurcooL/githubv4"
)

type recentStarsQuery struct {
	User struct {
		Login github.String
		Stars struct {
			TotalCount github.Int
			Edges      []struct {
				Cursor    github.String
				StarredAt github.DateTime
				Node      qlRepository
			}
		} `graphql:"starredRepositories(first: $count, after:$after, orderBy: {field: STARRED_AT, direction: DESC})"`
	} `graphql:"user(login:$username)"`
}

type StarredRepository struct {
	Repository
	StarredAt time.Time
}

func (handler *Handler) GetRecentStars(ctx context.Context, count int) ([]StarredRepository, error) {
	query := recentStarsQuery{}
	vars := Variables{
		"username": github.String(handler.username),
		"count":    github.Int(count + 1),
	}

	var after *github.String
	repositories := make([]StarredRepository, 0, count)

	for {
		vars["after"] = after
		err := handler.client.Query(ctx, &query, vars)
		if err != nil {
			return nil, err
		}

		if len(query.User.Stars.Edges) == 0 {
			break
		}

		for _, v := range query.User.Stars.Edges {
			if v.Node.IsPrivate {
				continue
			}

			repositories = append(repositories, StarredRepository{
				Repository: *v.Node.Unwrap(),
				StarredAt:  v.StarredAt.Time,
			})
		}

		if len(repositories) > count {
			break
		}
	}

	sort.Slice(repositories, func(i, j int) bool {
		return repositories[i].StarredAt.After(repositories[j].StarredAt)
	})

	if len(repositories) > count {
		return repositories[:count], nil
	}

	return repositories, nil
}
