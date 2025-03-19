package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"xml2json/internal/readconfig"
	"xml2json/internal/staff"
	"xml2json/internal/xml2json"
)

func main() {
	configPath := flag.String("config", "", "path to config file")

	flag.Parse()

	if *configPath == "" {
		fmt.Println("error: the argument -config=<path2file> must be specified")
		return
	}

	config := readconfig.ReadConfig(*configPath)

	rawData, errRF := os.ReadFile(config.Src)
	if errRF != nil {
		errRF = errors.New("error during reading file")
	}

	recodeData := staff.Win2UTF(rawData)

	data := xml2json.ReadXML(recodeData)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	jsonData := xml2json.WriteJSON(data)

	errMkdir := os.MkdirAll(filepath.Dir(config.Dst), os.ModePerm)
	if errMkdir != nil {
		errMkdir = errors.New("error during creating directories")
	}

	errWF := os.WriteFile(config.Dst, jsonData, 0644)
	if errWF != nil {
		errWF = errors.New("error during writing file")
	}

	err := errors.Join(errRF, errMkdir, errWF)
	if err != nil {
		fmt.Println("errors:", err)
	}
}
