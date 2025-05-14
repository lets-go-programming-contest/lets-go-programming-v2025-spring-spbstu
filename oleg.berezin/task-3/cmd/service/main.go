package main

import (
	"errors"
	"flag"
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
		panic("error: the argument -config=<path2file> must be specified")
	}

	config := readconfig.ReadConfig(*configPath)

	rawData, errRF := os.ReadFile(config.Src)
	if errRF != nil {
		panic(errRF)
	}

	recodeData := staff.Win2UTF(rawData)

	data, errRX := xml2json.ReadXML(recodeData)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	jsonData, errWJ := xml2json.WriteJSON(data)

	errMkdir := os.MkdirAll(filepath.Dir(config.Dst), os.ModePerm)
	if errMkdir != nil {
		panic(errMkdir)
	}

	errWF := os.WriteFile(config.Dst, jsonData, 0644)
	if errWF != nil {
		panic(errWF)
	}

	err := errors.Join(errRX, errWJ)
	if err != nil {
		panic(err)
	}
}
