package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Template struct {
	File   string
	Sheets map[string]SheetTemplate
}

type SheetTemplate struct {
	SrcStart int    `yaml:"src.start"`
	SrcEnd   int    `yaml:"src.end"`
	DstSheet string `yaml:"dst.sheet"`
	DstStart int    `yaml:"dst.start"`
	Mapping  map[int]int
}

func ReadTemplate(file string) (*Template, error) {
	template := &Template{}
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return template, fmt.Errorf("cannot find template file: %s, due to %v", file, err)
	}
	if err = yaml.Unmarshal(fileBytes, template); err != nil {
		return template, fmt.Errorf("malformed template file: %s, due to %v", file, err)
	}
	return template, nil
}
