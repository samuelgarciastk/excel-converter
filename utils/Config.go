package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Source      string `yaml:"source.dir"`
	Destination string `yaml:"destination.dir"`
	Template    string `yaml:"template.file"`
	ServerPort  int    `yaml:"server.port"`
}

func Load(path string) (*Config, error) {
	conf := &Config{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, fmt.Errorf("cannot find configuration file: %s, due to %v", path, err)
	}
	if err = yaml.Unmarshal(bytes, conf); err != nil {
		return conf, fmt.Errorf("malformed configuration file: %s, due to %v", path, err)
	}
	return conf, nil
}
