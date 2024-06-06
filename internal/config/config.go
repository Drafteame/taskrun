package config

import (
	"log"
	"os"

	"github.com/Drafteame/taskrun/internal/models"
)

func LoadConfigFromPath(path string) (*models.Jobs, error) {
	fileContent, err := getFileContent(path)
	if err != nil {
		return nil, err
	}

	jobs := &models.Jobs{}
	if err := jobs.FromYAML(fileContent); err != nil {
		return nil, err
	}

	return jobs, nil
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
