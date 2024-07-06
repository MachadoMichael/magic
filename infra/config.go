package infra

import (
	"os"
	"path/filepath"
)

type ConfigData struct {
	Version        float32
	MappingPath    string
	RepositoryPath string
}

var Config *ConfigData

func Init() error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}

	mappingPath := filepath.Join(rootPath, "pkg/mapper/internal/mapping.json")
	repositoryPath := filepath.Join(rootPath, "repository")

	Config = &ConfigData{
		Version:        0.7,
		MappingPath:    mappingPath,
		RepositoryPath: repositoryPath,
	}

	return nil

}
