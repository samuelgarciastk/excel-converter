package converter

import (
	"github.com/samuelgarciastk/excel-converter/utils"
	"log"
	"path/filepath"
)

type Converter interface {
	Convert() error
}

func BatchConvert(config utils.Config) {
	srcFileNames, err := utils.ListFiles(config.Source, utils.FileFilter())
	if err != nil {
		log.Fatalf("cannot list files in directory: %s", config.Source)
	}

	template, err := utils.ReadTemplate(config.Template)
	if err != nil {
		log.Fatalf("cannot read template file %s, due to %v", config.Template, err)
	}

	// concurrent
	fileNum := 0
	failedNum := 0
	for _, srcFileName := range srcFileNames {
		err = ConvertFile(filepath.Join(config.Source, srcFileName),
			filepath.Join(config.Destination, srcFileName),
			*template)
		fileNum++
		if err != nil {
			failedNum++
			log.Printf("cannot convert file: %s, due to %v", srcFileName, err)
		}
	}
	log.Printf("Convert %d files, %d succeeded, %d failed.", fileNum, fileNum-failedNum, failedNum)
}

func ConvertFile(srcFile, dstFile string, template utils.Template) error {
	converter := Excel{
		Source:      srcFile,
		Destination: dstFile,
		Template:    template,
	}
	return converter.Convert()
}
