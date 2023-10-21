package drivers

import (
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"github.com/wwmoraes/site/internal/drivers/github"
)

type GithubSource struct {
	Count int
	Query string
}

func (source *GithubSource) Fetch(ctx context.Context) ([]byte, error) {
	handler, err := github.NewHandler(ctx)
	if err != nil {
		return nil, err
	}

	// TODO(refactor): replace switch-case with an open-closed implementation
	switch source.Query {
	case "contributions":
		return handleData(handler.GetRecentContributions(ctx, source.Count))
	case "gists":
		return handleData(handler.GetRecentGists(ctx, source.Count))
	case "repositories":
		return handleData(handler.GetRecentRepositories(ctx, source.Count))
	case "releases":
		return handleData(handler.GetRecentReleases(ctx, source.Count))
	case "pullRequests":
		return handleData(handler.GetRecentPullRequests(ctx, source.Count))
	case "stars":
		return handleData(handler.GetRecentStars(ctx, source.Count))
	case "sponsors":
		return handleData(handler.GetRecentSponsors(ctx, source.Count))
	default:
		return nil, fmt.Errorf("unknown github query %s", source.Query)
	}
}

func handleData[T any](data []T, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(data, "", "  ")
}
