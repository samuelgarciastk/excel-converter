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
	SrcStart int    `yaml:"src_start"`
	SrcEnd   int    `yaml:"src_end"`
	DstSheet string `yaml:"dst_sheet"`
	DstStart int    `yaml:"dst_start"`
	Mapping  map[int]int
}

func ReadTemplate(file string) (*Template, error) {
	template := &Template{}

	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return template, fmt.Errorf("cannot find template file: %s, due to %v", file, err)
	}

	err = yaml.Unmarshal(fileBytes, template)
	if err != nil {
		return template, fmt.Errorf("malformed template file: %s, due to %v", file, err)
	}
	return template, nil
}
