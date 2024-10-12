package update

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
	"github.com/wwmoraes/go-rwfs"
	"github.com/wwmoraes/site/internal/adapters"
	"github.com/wwmoraes/site/internal/drivers"
	"github.com/wwmoraes/site/internal/entities"
	"github.com/wwmoraes/site/internal/usecases"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "refreshes data files",
		RunE:  update,
	}

	return cmd
}

func update(cmd *cobra.Command, args []string) error {
	log.Println("reading site config")

	data, err := os.ReadFile(".site.toml")
	if err != nil {
		return err
	}

	log.Println("loading site config")

	config, err := loadSiteConfig(string(data))
	if err != nil {
		return err
	}

	store := &drivers.HugoDataStore{FS: rwfs.OSDirFS("data")}

	updaters := make(usecases.DataUpdaters, 0, len(config.Data.Sources))

	for _, sourceConfig := range config.Data.Sources {
		sourceAdapter, err := createDataSourceFromConfig(sourceConfig)
		if err != nil {
			return err
		}

		updaters = append(updaters, &adapters.DataPipe{
			Source: sourceAdapter,
			Name:   sourceConfig.Target,
			Store:  store,
		})
	}

	http.DefaultClient.Timeout = time.Second

	errors := updaters.ExecuteAll(cmd.Context(), runtime.NumCPU())

	for err := range errors {
		cmd.PrintErrln(err)
	}

	return nil
}

func createDataSourceFromConfig(config entities.DataSourceConfig) (adapters.DataSource, error) {
	// TODO(refactor): replace switch-case with an open-closed implementation
	switch config.Adapter {
	case "github":
		return createGithubSourceFromConfig(config)
	case "goodreads":
		return createGoodreadsSourceFromConfig(config)
	default:
		return nil, fmt.Errorf("unknown data source adapter %s", config.Adapter)
	}
}

func loadSiteConfig(data string) (*entities.SiteConfig, error) {
	config := &entities.SiteConfig{}

	_, err := toml.Decode(data, config)

	return config, err
}

func createGithubSourceFromConfig(config entities.DataSourceConfig) (adapters.DataSource, error) {
	query, err := config.Parameters.GetString("query")
	if err != nil {
		return nil, err
	}

	count, err := config.Parameters.GetInt64("count")
	if err != nil {
		return nil, err
	}

	return &drivers.GithubSource{
		Count: int(count),
		Query: query,
	}, nil
}

func createGoodreadsSourceFromConfig(config entities.DataSourceConfig) (adapters.DataSource, error) {
	list, err := config.Parameters.GetString("list")
	if err != nil {
		return nil, err
	}

	shelf, err := config.Parameters.GetString("shelf")
	if err != nil {
		return nil, err
	}

	return &drivers.GoodreadsSource{
		List:  list,
		Shelf: shelf,
	}, nil
}
