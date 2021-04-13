package configs

import (
	"os"
	"os/user"
	"path/filepath"
)

// DefaultConfigsFile stores cli credentials and hosts.
const DefaultConfigsFile = "configs"

// DefaultDir retrieves the default directory used to store CLI configs.
func DefaultDir() (string, error) {
	var dir string
	// By default, store meta and data files in current users home directory
	u, err := user.Current()
	if err == nil {
		dir = u.HomeDir
	} else if home := os.Getenv("HOME"); home != "" {
		dir = home
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		dir = wd
	}
	dir = filepath.Join(dir, ".influxdbv2")

	return dir, nil
}
