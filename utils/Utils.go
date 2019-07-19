package utils

import (
	"io/ioutil"
)

func ListFiles(dir string, filter func(string) bool) ([]string, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var filePaths []string
	for _, fileInfo := range fileInfos {
		if filter(fileInfo.Name()) {
			filePaths = append(filePaths, fileInfo.Name())
		}
	}
	return filePaths, nil
}
