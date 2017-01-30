package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/pkg/errors"
)

func load() (*toml.TomlTree, error) {
	tree, err := toml.LoadFile(filepath.Join(os.Getenv("HOME"), ".config/tiger/config.toml"))
	if strings.Contains(err.Error(), "no such file or directory") {
		// Do not fail if no configuration present, we will use default values
		err = nil
	}
	return tree, err
}

func Get(key string, defaultValue interface{}) (interface{}, error) {
	tree, err := load()
	if tree == nil {
		return defaultValue, err
	}
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting %q config", key)
	}
	if tree.Has(key) {
		return tree.Get(key), nil
	}
	return defaultValue, nil
}

func GetAsString(key, defaultValue string) (string, error) {
	value, err := Get(key, defaultValue)
	return value.(string), err
}
