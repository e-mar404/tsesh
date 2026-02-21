package config

type Config struct {
	Search Search `toml:"search"`
}

type Search struct {
	Paths []string `toml:"search_paths"`
	IgnorePattern string `toml:"ignore_pattern"`
	IgnoreHidden bool `toml:"ignore_hidden"`
	MaxDepth int `toml:"max_depth"`	
}
