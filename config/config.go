package config

import (
	"errors"
	"os"
	"path/filepath"

	lua "github.com/yuin/gopher-lua"
)

var ErrEmptyConfigFile = errors.New("config.lua is empty")
var ErrNoSearchPaths = errors.New("config.lua does not have field search_paths")

type Config struct {
	SearchPaths []string
}

func Load() (*Config, error) {
	home, _ := os.UserHomeDir()
	cfgPath := filepath.Join(home, ".config", "tsesh", "config.lua")

	L := lua.NewState()
	defer L.Close()

	if err := L.DoFile(cfgPath); err != nil {
		return nil, err
	}

	cfg := &Config{}
	ret := L.Get(-1)
	if ret == lua.LNil {
		return nil, ErrEmptyConfigFile
	}

	if tbl, ok := ret.(*lua.LTable); ok {
		paths := L.GetField(tbl, "search_paths")
		if paths == lua.LNil {
			return nil, ErrNoSearchPaths
		}

		L.ForEach(paths.(*lua.LTable), func(_, v lua.LValue) {
			cfg.SearchPaths = append(cfg.SearchPaths, v.String())
		})
	}
	L.Pop(1)

	return cfg, nil
}
