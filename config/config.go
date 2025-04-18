package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/devrian/golb/env"
	"gopkg.in/yaml.v2"
)

func ReadModuleConfig(cfg interface{}, path string, module string) error {
	environ := env.Get()
	getFormatFile := filePath(path)

	switch getFormatFile {
	case ".json":
		fname := path + "/" + module + "." + environ + ".json"
		jsonFile, err := os.ReadFile(fname)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonFile, cfg)
	default:
		fname := path + "/" + module + "." + environ + ".yaml"
		yamlFile, err := os.ReadFile(fname)
		if err != nil {
			return err
		}
		return yaml.Unmarshal(yamlFile, cfg)
	}

}

func filePath(root string) string {
	var file string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		file = filepath.Ext(info.Name())
		return nil
	})
	return file
}
