package github

import (
	"context"
	"time"

	github "github.com/shurcooL/githubv4"
)

type recentReleasesQuery struct {
	User struct {
		// Login                     github.String
		RepositoriesContributedTo struct {
			// TotalCount github.Int
			Edges []struct {
				Cursor github.String
				Node   struct {
					qlRepository
					Releases qlRelease `graphql:"releases(first: 10, orderBy: {field: CREATED_AT, direction: DESC})"`
				}
			}
		} `graphql:"repositoriesContributedTo(first: 100, after:$after includeUserRepositories: true, contributionTypes: COMMIT, privacy: PUBLIC)"` //nolint:lll
	} `graphql:"user(login:$username)"`
}

type RepositoryRelease struct {
	Repository `json:"Repository"` //nolint:tagliatelle
	Release
}

type Release struct {
	Name        string
	TagName     string
	PublishedAt time.Time
	URL         string
}

func (handler *Handler) GetRecentReleases(ctx context.Context, count int) ([]RepositoryRelease, error) {
	query := recentReleasesQuery{}
	repositories := make([]RepositoryRelease, 0, count)
	vars := Variables{
		"username": github.String(handler.username),
	}

	var after *github.String

	for {
		vars["after"] = after

		err := handler.client.Query(ctx, &query, vars)
		if err != nil {
			return nil, err
		}

		if len(query.User.RepositoriesContributedTo.Edges) == 0 {
			break
		}

		var repository *RepositoryRelease
		for _, v := range query.User.RepositoriesContributedTo.Edges {
			repository = &RepositoryRelease{
				Repository: *v.Node.Unwrap(),
			}

			for _, rel := range v.Node.Releases.Nodes {
				if rel.IsPrerelease || rel.IsDraft {
					continue
				}

				if v.Node.Releases.Nodes[0].TagName == "" ||
					v.Node.Releases.Nodes[0].PublishedAt.Time.IsZero() {
					continue
				}

				repository.Release = *v.Node.Releases.Nodes[0].Unwrap()

				break
			}

			if !repository.Release.PublishedAt.IsZero() {
				repositories = append(repositories, *repository)
			}

			after = github.NewString(v.Cursor)
		}
	}

	if len(repositories) > count {
		return repositories[:count], nil
	}

	return repositories, nil
}
