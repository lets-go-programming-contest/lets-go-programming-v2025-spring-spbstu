package main

import (
	"flag"
	"os"
	"sort"
	"xml2json/internal/readconfig"
	"xml2json/internal/staff"
	"xml2json/internal/xml2json"
)

func main() {
	configPath := flag.String("config", "", "path to config file")

	flag.Parse()

	if *configPath == "" {
		panic("Error: the argument -config=<path2file> must be specified")
	}

	config := readconfig.ReadConfig(*configPath)

	rawData, err := os.ReadFile(config.Src)
	if err != nil {
		panic("Error during reading file")
	}

	recodeData := staff.Win2UTF(rawData)

	data := xml2json.ReadXML(recodeData)

	sort.Slice(data, func(i, j int) bool {
		return data[i].Value > data[j].Value
	})

	jsonData := xml2json.WriteJSON(data)

	os.WriteFile(config.Dst, jsonData, 0644)
}
