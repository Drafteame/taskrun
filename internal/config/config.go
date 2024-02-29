package config

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"gopkg.in/yaml.v3"

	"github.com/Drafteame/taskrun/internal/models"
)

func LoadConfigFromPath(path string, replace map[string]string) (models.Jobs, error) {
	fileContent, err := getFileContent(path)
	if err != nil {
		return models.Jobs{}, err
	}

	strContent, errRep := applyReplacers(string(fileContent), replace)
	if errRep != nil {
		return models.Jobs{}, errRep
	}

	config := models.Jobs{}

	if errDecode := yaml.Unmarshal([]byte(strContent), &config); errDecode != nil {
		return models.Jobs{}, errDecode
	}

	return config, nil
}

func getFileContent(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		if errClose := file.Close(); errClose != nil {
			log.Println("Error closing file:", errClose.Error())
		}
	}(file)

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)

	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

func applyReplacers(content string, replace map[string]string) (string, error) {
	for k, v := range replace {
		exp, err := regexp.Compile(fmt.Sprintf(`\$\{(\s+)?%s(\s+)?\}`, k))
		if err != nil {
			return "", err
		}

		content = exp.ReplaceAllString(content, v)
	}

	return content, nil
}
