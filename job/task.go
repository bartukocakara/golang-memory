package job

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang-memory/config"
	"golang-memory/handler"
	"golang-memory/logger"
)

const (
	JsonFileNameFormat = "%d_golang-memoryData.json"
)

func CheckExistingData() {
	files, err := ioutil.ReadDir(config.EXPORT_FILE_PATH)
	if err != nil {
		logger.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(config.EXPORT_FILE_PATH, file.Name())
		f, err := os.ReadFile(path)
		if err != nil {
			logger.Fatal.Println(err)
		}

		err = json.Unmarshal(f, &handler.InMemDB)
		if err != nil {
			logger.Fatal.Println(err)
		}
	}
}

func removeFileInDirectory() {
	files, err := ioutil.ReadDir(config.EXPORT_FILE_PATH)
	if err != nil {
		logger.Info.Println(err)
		return
	}

	for _, file := range files {
		path := filepath.Join(config.EXPORT_FILE_PATH, file.Name())
		err := os.Remove(path)
		if err != nil {
			logger.Error.Println(err)
		}
	}
}

func saveData(ticker *time.Ticker, quit chan struct{}) {
	for {
		select {
		case <-ticker.C:
			if _, err := os.Stat(config.EXPORT_FILE_PATH); os.IsNotExist(err) {
				err := os.Mkdir(config.EXPORT_FILE_PATH, os.ModePerm)
				if err != nil {
					logger.Info.Println(err)
				}
			} else {
				removeFileInDirectory()
			}

			dataBytes, err := json.Marshal(handler.InMemDB)
			if err != nil {
				logger.Fatal.Println(err)
			}

			fileName := filepath.Join(config.EXPORT_FILE_PATH, fmt.Sprintf(JsonFileNameFormat, time.Now().Unix()))
			err = ioutil.WriteFile(fileName, dataBytes, 0777)
			if err != nil {
				logger.Fatal.Println(err)
			} else {
				logger.Info.Printf("All data added to file: %s", fileName)
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func Startjob() {
	duration, err := strconv.Atoi(config.RECORD_FREQ)
	if err != nil {
		duration = 10
	}
	ticker := time.NewTicker(time.Duration(duration) * time.Minute)
	quit := make(chan struct{})
	go saveData(ticker, quit)
}
