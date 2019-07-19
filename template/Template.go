package template

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Template struct {
	Header []string
	Sheets map[string]SheetTemplate
}

type SheetTemplate struct {
	Start   int
	End     int
	Mapping map[int]int
}

func readTemplate(file string) Template {
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Cannot find template file: %s\n", file)
		panic(err)
	}

	template := Template{}
	err = yaml.Unmarshal(fileBytes, &template)
	if err != nil {
		fmt.Printf("Malformed template file: %s\n", file)
		panic(err)
	}
	return template
}

func DetermineTemplate(file string) Template {
	return readTemplate("/home/stk/go/src/github.com/samuelgarciastk/excel-converter/conf/template1.yml")
}
