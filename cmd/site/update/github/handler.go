package github

import (
	"context"
	"net/http"
	"os"
	"sync"

	github "github.com/shurcooL/githubv4"
	"github.com/wwmoraes/go-rwfs"
	"golang.org/x/oauth2"
)

type Handler struct {
	client       *github.Client
	username     string
	fileHandlers map[string]HandlerFn[any]
}

type HandlerFn[T any] func(context.Context, int) ([]T, error)

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

	handler := &Handler{
		client:   client,
		username: username,
	}

	handler.fileHandlers = map[string]HandlerFn[any]{
		"contributions.json": genericHandle(handler.GetRecentContributions),
		"repositories.json":  genericHandle(handler.GetRecentRepositories),
		"releases.json":      genericHandle(handler.GetRecentReleases),
		"pullRequests.json":  genericHandle(handler.GetRecentPullRequests),
		"starred.json":       genericHandle(handler.GetRecentStars),
		"gists.json":         genericHandle(handler.GetRecentGists),
		"sponsors.json":      genericHandle(handler.GetRecentSponsors),
	}

	return handler, nil
}

func (handler *Handler) HandleAll(ctx context.Context, count int, fsys rwfs.FS) <-chan error {
	errors := make(chan error, 10)

	go func() {
		var wg sync.WaitGroup

		for name, handle := range handler.fileHandlers {
			wg.Add(1)

			go func(name string, handle HandlerFn[any]) {
				data, err := handle(ctx, count)
				if err != nil {
					errors <- err

					return
				}

				err = writeFile(fsys, name, data)
				if err != nil {
					errors <- err
				}
			}(name, handle)
		}

		wg.Wait()

		close(errors)
	}()

	return errors
}

func genericHandle[T any](handle func(context.Context, int) ([]T, error)) HandlerFn[any] {
	return func(ctx context.Context, count int) ([]interface{}, error) {
		data, err := handle(ctx, count)
		if err != nil {
			return nil, err
		}

		generic := make([]interface{}, len(data))
		for index, entry := range data {
			generic[index] = entry
		}

		return generic, err
	}
}
