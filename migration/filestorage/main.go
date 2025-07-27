package filestorage

import (
	"encoding/json"
	"os"
)

type Modulos struct {
	Name   string   `json:"name"`
	Folder string   `json:"folder"`
	Videos []string `json:"videos"`
}

// "../assets/media/mod.json"
func GetJsonConfig(filePath string) map[string]Modulos {

	mod := map[string]Modulos{}

	f, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(f, &mod)
	if err != nil {
		panic(err)
	}

	return mod
}
