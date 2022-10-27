package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Settings struct {
	Server string `yaml:"server"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func Parse(fp string) (s Settings, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	if err = dec.Decode(&s); err != nil {
		return
	}

	return
}
