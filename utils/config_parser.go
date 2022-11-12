package utils

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/aosasona/stripr/types"
)

func CheckConfigExists(params ...interface{}) (string, error) {
	var (
		pwd,
		configName string
		err error
	)
	configName = "stripr.json"
	if len(params) == 0 {
		pwd, err = os.Getwd()

		if err != nil {
			return "", &types.CustomError{Message: "Unable to get current working directory"}
		}
	} else {
		pwd = params[0].(string)
	}

	pwd = strings.Trim(pwd, "/")
	configPath := pwd + "/" + configName
	if fileExists, err := os.Stat(configPath); err != nil || fileExists.IsDir() {
		return "", &types.CustomError{Message: "Config file not found"}
	}
	return configPath, nil
}

func ReadConfig(root string) (map[string]interface{}, error) {
	var configRoot string
	if root != "" {
		configRoot = root
	} else {
		configRoot, _ = os.Getwd()
	}
	path, err := CheckConfigExists(configRoot)
	if err != nil {
		return nil, &types.CustomError{Message: "Config file not found"}
	}

	config, err := os.ReadFile(path)
	if err != nil {
		return nil, &types.CustomError{Message: "Unable to read config file"}
	}

	var configInJSON map[string]interface{}
	err = json.Unmarshal(config, &configInJSON)
	if err != nil {
		return nil, &types.CustomError{Message: "Corrupted config file, unable to parse"}
	}

	return configInJSON, nil
}
