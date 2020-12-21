package utils

import (
	"io"
	"net/http"
	"os"
	"path"
)

var dataFolder = ""
var userHome = ""

func GetUserHome() string {
	if userHome != "" {
		return userHome
	}
	var err error
	userHome, err = os.UserHomeDir()
	HandleError(err, "Cannot get user home")
	return userHome
}

func createDataFolder(dataFolder string) {
	err := os.MkdirAll(dataFolder, 0744)
	HandleError(err, "Cannot create data folder at "+dataFolder)
}

func LoadDataFolder() string {
	if dataFolder != "" {
		return dataFolder
	}

	dataFolder = path.Join(GetUserHome(), ".cache", "tuner")

	_, err := os.Stat(dataFolder)
	if os.IsNotExist(err) {
		createDataFolder(dataFolder)
	}

	return dataFolder
}

func DownloadFile(fileURL, targetFilePath string) error {
	res, err := http.Get(fileURL)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	file, err := os.Create(targetFilePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, res.Body)
	return err
}
