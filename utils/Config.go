package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Source      string
	Destination string
	Template    string
}

func Load(path string) (*Config, error) {
	conf := &Config{}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return conf, fmt.Errorf("cannot find configuration file: %s, due to %v", path, err)
	}
	err = yaml.Unmarshal(bytes, conf)
	if err != nil {
		return conf, fmt.Errorf("malformed configuration file: %s, due to %v", path, err)
	}
	return conf, nil
}
