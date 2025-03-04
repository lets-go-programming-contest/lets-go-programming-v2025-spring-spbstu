package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"task-3/internal/cbr"
	configreader "task-3/internal/configReader"
)

func main() {
	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}

func Run() error {
	config, err := configreader.ReadConfig()
	if err != nil {
		return err
	}

	valCursXml, err := cbr.ReadCbrXml(config.InputFile)
	if err != nil {
		return err
	}

	valCursJson, err := cbr.GetSortedJsonString(valCursXml)
	if err != nil {
		return err
	}

	err = writeFile(config.OutputFile, valCursJson)
	if err != nil {
		return err
	}

	fmt.Println("JSON file created successfully!")
	return nil
}

func writeFile(filePath string, data *[]byte) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(filePath), 0777)
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(*data)
	if err != nil {
		return err
	}

	return nil
}
