package config

// This model defines the config option for the gwhich command line utility.
// Each instance of this model contains settings for a app from which
// to pull command line shortcuts from
type Application struct {
	ConfigPath  string `toml:"config_path,omitempty"`
	Description string `toml:"description,omitempty"`
	Name        string `toml:"name,omitempty"`
	Regex       string `toml:"regex,omitempty"`
}
