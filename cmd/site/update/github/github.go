package github

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/cobra"
)

type Variables map[string]interface{}

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "github",
		Short: "fetch GitHub user data",
		RunE:  cmdFetch,
	}

	flags := cmd.Flags()
	flags.Int("count", 10, "maximum number of entries for each data source")
	flags.String("path", "data/github/recents", "base directory where the JSON results will be stored at")

	return cmd
}

func cmdFetch(cmd *cobra.Command, args []string) error {
	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		return err
	}

	basePath, err := cmd.Flags().GetString("path")
	if err != nil {
		return err
	}

	err = os.MkdirAll(basePath, 0750)
	if err != nil {
		return err
	}

	handler, err := NewHandler(cmd.Context())
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var errors []error

	// recent contributions
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentContributions(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "contributions.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent repositories
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentRepositories(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "repositories.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent releases
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentReleases(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "releases.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent pull requests
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentPullRequests(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "pullRequests.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent starred repositories
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentStars(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "starred.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent gists
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentGists(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "gists.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// recent sponsors
	wg.Add(1)
	go func() {
		defer wg.Done()

		data, err := handler.GetRecentSponsors(cmd.Context(), count)
		if err != nil {
			errors = append(errors, err)
			return
		}

		err = writeFile(filepath.Join(basePath, "sponsors.json"), data)
		if err != nil {
			errors = append(errors, err)
			return
		}
	}()

	// wait all concurrent requests and aggregate the errors, if any
	wg.Wait()
	if len(errors) > 0 {
		return fmt.Errorf("fetch returned %d errors: %v", len(errors), errors)
	}

	return nil
}

func writeFile(path string, res any) error {
	fd, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0640)
	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = marshal(res, fd)

	return err
}

func marshal(v any, w io.Writer) (int64, error) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return 0, err
	}

	n, err := w.Write(data)

	return int64(n), err
}
