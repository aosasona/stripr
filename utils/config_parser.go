package utils

import (
	"encoding/json"
	"os"
	"strings"
)

func CheckConfigExists(params ...interface{}) (string, bool) {
	var (
		pwd,
		configName string
		err error
	)
	configName = "stripr.json"

	if len(params) == 0 {
		pwd, err = os.Getwd()

		if err != nil {
			Terminate(err)
		}
	} else {
		pwd = params[0].(string)
	}

	pwd = strings.Trim(pwd, "/")
	configPath := pwd + "/" + configName
	if fileExists, err := os.Stat(configPath); err != nil || fileExists.IsDir() {
		return "", false
	}
	return configPath, true
}

func ReadConfig(root string) map[string]interface{} {
	var configRoot string
	if root != "" {
		configRoot = root
	} else {
		configRoot, _ = os.Getwd()
	}
	path, exists := CheckConfigExists(configRoot)
	if !exists {
		return nil
	}

	config, err := os.ReadFile(path)
	if err != nil {
		Terminate(err)
	}

	var configInJSON map[string]interface{}
	err = json.Unmarshal(config, &configInJSON)
	if err != nil {
		Terminate(err)
	}

	return configInJSON
}
