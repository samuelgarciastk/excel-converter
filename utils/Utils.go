package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func FileFilter() func(string) bool {
	return func(fileName string) bool {
		return !strings.HasPrefix(fileName, "~") &&
			!strings.HasPrefix(fileName, ".")
	}
}

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

func CopyFile(src, dst string) error {
	srcFileInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !srcFileInfo.Mode().IsRegular() {
		return fmt.Errorf("cannot copy non-regular source file %s (%q)", srcFileInfo.Name(), srcFileInfo.Mode().String())
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		openErr := srcFile.Close()
		if err == nil {
			err = openErr
		}
	}()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer func() {
		createErr := dstFile.Close()
		if err == nil {
			err = createErr
		}
	}()

	if _, err = io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return dstFile.Sync()
}
