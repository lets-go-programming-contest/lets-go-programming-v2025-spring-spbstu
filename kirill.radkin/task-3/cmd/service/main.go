package main

import (
	"flag"
	"os"

	jsonencoder "github.com/yanelox/task-3/internal/json_encoder"
	vl "github.com/yanelox/task-3/internal/valute"
	vl_sort "github.com/yanelox/task-3/internal/valute_sorting"
	xd "github.com/yanelox/task-3/internal/xml_decoder"
	yd "github.com/yanelox/task-3/internal/yaml_decoder"
)

var InputConfigPath = flag.String("config", "default-config.yaml", "Path to config file")

func main() {
	if len(os.Args) != 1 && len(os.Args) != 3 {
		panic("Incorrect arguments")
	}

	flag.Parse()

	var InputConfig yd.YamlConfig

	err := yd.Decode(*InputConfigPath, &InputConfig)
	if err != nil {
		panic(err)
	}

	input_file := InputConfig.InputFile

	var ValCurs vl.ValCurs
	err = xd.Decode(input_file, &ValCurs)
	if err != nil {
		panic(err)
	}

	sort_by_value := func(v1, v2 vl.Valute) bool {
		return v1.Value < v2.Value
	}

	vl_sort.ValuteSort(ValCurs.Valutes, sort_by_value)

	jsonencoder.Encode(InputConfig.OutputFile, ValCurs.Valutes)
}
