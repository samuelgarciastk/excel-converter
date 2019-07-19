package main

import (
	"fmt"
	"github.com/samuelgarciastk/excel-converter/converter"
	"github.com/samuelgarciastk/excel-converter/template"
	"github.com/samuelgarciastk/excel-converter/utils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	conf := loadConf()

	srcFileNames, err := utils.ListFiles(conf.Source, fileFilter())
	if err != nil {
		fmt.Printf("Cannot list files in directory: %s\n", conf.Source)
	}

	for _, srcFileName := range srcFileNames {
		fileConverter := genConverter(srcFileName, conf)
		if fileConverter != nil {
			fileConverter.Convert()
		}
	}
}

func fileFilter() func(string) bool {
	return func(file string) bool {
		return !strings.HasPrefix(file, "~")
	}
}

func genConverter(fileName string, conf Conf) converter.Converter {
	fileTemplate := template.DetermineTemplate(fileName)
	switch ext := filepath.Ext(fileName); ext {
	case ".xlsx":
		return &converter.Excel{
			Source:   filepath.Join(conf.Source, fileName),
			Target:   filepath.Join(conf.Target, fileName),
			Template: fileTemplate,
		}
	default:
		fmt.Printf("Extension [%s] is not supported.", ext)
		return nil
	}
}

// Config
const confPath = "conf/conf.yml"

type Conf struct {
	Source string
	Target string
}

func loadConf() Conf {
	conf := Conf{}
	bytes, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic("Cannot find configuration file.")
	}
	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		panic("Malformed configuration.")
	}
	return conf
}
