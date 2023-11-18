package hugo

import (
	"os"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/config/allconfig"
	"github.com/gohugoio/hugo/deps"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
)

type Sites struct {
	*hugolib.HugoSites
	ContentFactory hugolib.ContentFactory
}

type Config struct {
	CacheDir    string
	ConfigDir   string
	ConfigFile  string
	Environment string
	PublishDir  string
	ThemesDir   string
	WorkingDir  string
}

type ConfigOption interface {
	Apply(config *Config) error
}

type ConfigOptionFn func(config *Config) error

func (fn ConfigOptionFn) Apply(config *Config) error {
	return fn(config)
}

func New(options ...ConfigOption) (*Sites, error) {
	conf := Config{
		ConfigDir:   "config",
		Environment: "development",
		PublishDir:  "public",
		ThemesDir:   "themes",
	}

	var err error
	for _, option := range options {
		err = option.Apply(&conf)
		if err != nil {
			return nil, err
		}
	}

	if conf.WorkingDir == "" || conf.WorkingDir == "." {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		conf.WorkingDir = pwd
	}

	pwdFs := hugofs.NewDefault(config.BaseConfig{
		WorkingDir: conf.WorkingDir,
		PublishDir: conf.PublishDir,
		ThemesDir:  conf.ThemesDir,
		CacheDir:   conf.CacheDir,
	})

	configs, err := allconfig.LoadConfig(allconfig.ConfigSourceDescriptor{
		Fs:          pwdFs.Source,
		Environment: conf.Environment,
		ConfigDir:   conf.ConfigDir,
		Filename:    conf.ConfigFile,
	})
	if err != nil {
		return nil, err
	}

	sites, err := hugolib.NewHugoSites(deps.DepsCfg{
		Fs:      pwdFs,
		Configs: configs,
	})
	if err != nil {
		return nil, err
	}

	return &Sites{
		HugoSites:      sites,
		ContentFactory: hugolib.NewContentFactory(sites),
	}, nil
}
