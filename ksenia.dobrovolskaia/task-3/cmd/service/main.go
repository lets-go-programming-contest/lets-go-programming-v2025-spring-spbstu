package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"

  "github.com/kseniadobrovolskaia/task-3/internal/sortValue"
  "github.com/kseniadobrovolskaia/task-3/internal/config"
  "github.com/kseniadobrovolskaia/task-3/internal/xmlToJson"
)

var configPath = flag.String("config", "", "Путь до файла конфигурации .yaml. \nФормат конфигурации:\n    input-file: \"source/input_02_03_2002.xml\"\n    output-file: \"result/output_02_03_2002.json\"\n\ninput-file - путь до файла с состоянием валют с сайта ЦБР\noutput-file - путь до файла с результатом программы")


func main() {
	flag.Parse()
	if *configPath == "" {
		color.Red("--config not specified")
		os.Exit(0)
	}

	//------------------READ CONFIG FILE-------------------
  config, err := config.ReadConfigFile(*configPath)
	if err != nil {
	  color.Red(err.Error())
    os.Exit(1)
  }
	//--------------------READ XML FILE--------------------
	vals, err := xmlToJson.ReadXMLFile(config.InputFile)
	if err != nil {
	  color.Red(err.Error())
    os.Exit(1)
  }

	//-----------------------SORTING-----------------------
	sort.Sort(sortValue.ByValue(vals))

	//------------------DUMP IN JSON FILE------------------
	lenBytes, err := xmlToJson.WriteInJSONFile(config.OutputFile, vals)
	if err != nil {
		panic(err)
	}
  
  fmt.Printf("%d %s\n", lenBytes, color.GreenString("bytes written in " + config.OutputFile))
}
