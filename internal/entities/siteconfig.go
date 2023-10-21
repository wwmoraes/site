package entities

type SiteConfig struct {
	Data DataConfig `toml:"data"`
}

type DataConfig struct {
	Sources []DataSourceConfig `toml:"sources"`
}

type DataSourceConfig struct {
	Adapter    string     `toml:"adapter"`
	Target     string     `toml:"target"`
	Parameters Parameters `toml:"parameters"`
}
