package github

import (
	"context"
	"fmt"
	"sort"
	"time"

	github "github.com/shurcooL/githubv4"
)

type recentContributionsQuery struct {
	User struct {
		// Login                   github.String
		ContributionsCollection struct {
			CommitContributionsByRepository []struct {
				Contributions struct {
					Edges []struct {
						// Cursor github.String
						Node struct {
							OccurredAt github.DateTime
						}
					}
				} `graphql:"contributions(first: 1)"`
				Repository qlRepository
			} `graphql:"commitContributionsByRepository(maxRepositories: 100)"`
		}
	} `graphql:"user(login:$username)"`
}

type RecentContribution struct {
	Repository Repository
	OccurredAt time.Time
}

func (handler *Handler) GetRecentContributions(ctx context.Context, count int) ([]RecentContribution, error) {
	query := recentContributionsQuery{}
	vars := Variables{
		"username": github.String(handler.username),
	}

	err := handler.client.Query(ctx, &query, vars)
	if err != nil {
		return nil, err
	}

	contributions := make([]RecentContribution, 0, count)

	for _, repository := range query.User.ContributionsCollection.CommitContributionsByRepository {
		if repository.Repository.IsPrivate {
			continue
		}

		if string(repository.Repository.NameWithOwner) == fmt.Sprintf("%s/%s", handler.username, handler.username) {
			continue
		}

		contributions = append(contributions, RecentContribution{
			Repository: *repository.Repository.Unwrap(),
			OccurredAt: repository.Contributions.Edges[0].Node.OccurredAt.Time,
		})

		if len(contributions) == count {
			break
		}
	}

	sort.Slice(contributions, func(i, j int) bool {
		return contributions[i].OccurredAt.After(contributions[j].OccurredAt)
	})

	return contributions, nil
}
