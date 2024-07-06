package infra

import (
	"fmt"
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

	ex, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	exPath := filepath.Dir(ex)

	mappingPath := filepath.Join(exPath, "pkg/mapper/internal/mapping.json")
	repositoryPath := filepath.Join(exPath, "repository")

	Config = &ConfigData{
		Version:        0.8,
		MappingPath:    mappingPath,
		RepositoryPath: repositoryPath,
	}

	return nil

}
