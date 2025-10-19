package tidy

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"strconv"
	"time"

	"github.com/cloudflare/cloudflare-go/v6"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/pages"
	"github.com/goccy/go-json"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

const deploymentsPageSize = 10

type ResultInfo struct {
	Page       uint64 `json:"page,omitempty"`
	PerPage    uint64 `json:"per_page,omitempty"`
	Count      uint64 `json:"count,omitempty"`
	TotalCount uint64 `json:"total_count,omitempty"`
	TotalPages uint64 `json:"total_pages,omitempty"`
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tidy",
		Short: "trims deployment history, keeping only aliased ones",
		Long:  "removes past deployments of the site (needed to delete the site)",
		RunE:  tidy,
	}

	return cmd
}

//nolint:cyclop,funlen // TODO refactor
func tidy(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	accountID := os.Getenv("CLOUDFLARE_ACCOUNT_ID")
	projectName := os.Getenv("CLOUDFLARE_PROJECT_NAME")

	client := cloudflare.NewClient()

	resultInfo := ResultInfo{
		Page:       1,
		PerPage:    deploymentsPageSize,
		Count:      0,
		TotalCount: 0,
		TotalPages: 1,
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.SetLimit(runtime.GOMAXPROCS(0))

	for {
		cmd.Println(time.Now().Local().Format(time.RFC3339Nano), "listing deployments...")

		deployments, err := client.Pages.Projects.Deployments.List(ctx, projectName, pages.ProjectDeploymentListParams{
			AccountID: cloudflare.String(accountID),
			// Env:       cloudflare.F(pages.ProjectDeploymentListParamsEnv("production")),
		},
			option.WithQuery("per_page", strconv.FormatUint(resultInfo.PerPage, 10)),
			option.WithQuery("page", strconv.FormatUint(resultInfo.Page, 10)),
		)
		if err != nil {
			return fmt.Errorf("failed to list preview deployments: %w", err)
		}

		if deployments == nil || len(deployments.Result) < 1 {
			cmd.Println(time.Now().Local().Format(time.RFC3339Nano), "no deployments to delete, stopping")

			break
		}

		err = json.Unmarshal([]byte(deployments.JSON.ExtraFields["result_info"].Raw()), &resultInfo)
		if err != nil {
			return fmt.Errorf("failed to decode result info: %w", err)
		}

		for _, deployment := range deployments.Result {
			// skip aliased deployments
			if slices.ContainsFunc(deployment.Aliases, func(alias string) bool {
				return alias == "https://artero.dev" || alias == "https://staging.artero.dev"
			}) {
				cmd.Println(
					time.Now().Local().Format(time.RFC3339Nano),
					"skipping deployment",
					deployment.URL,
					"with aliases",
					deployment.Aliases,
				)

				continue
			}

			group.Go(func() error {
				cmd.Println(time.Now().Local().Format(time.RFC3339Nano), "deleting deployment", deployment.URL)

				_, err := client.Pages.Projects.Deployments.Delete(
					groupCtx,
					projectName,
					deployment.ID,
					pages.ProjectDeploymentDeleteParams{
						AccountID: cloudflare.String(accountID),
					}, option.WithQuery("force", "true"))
				if err != nil {
					return fmt.Errorf("failed to delete deployment: %w", err)
				}

				return nil
			})
		}

		resultInfo.Page++

		if resultInfo.Page > resultInfo.TotalPages {
			break
		}
	}

	return group.Wait()
}
