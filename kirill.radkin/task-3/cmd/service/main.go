package main

import (
	"flag"

	je "github.com/yanelox/task-3/internal/json_encoder"
	"github.com/yanelox/task-3/internal/mytypes"
	vl_sort "github.com/yanelox/task-3/internal/valute_sorting"
	xd "github.com/yanelox/task-3/internal/xml_decoder"
	yd "github.com/yanelox/task-3/internal/yaml_decoder"
)

var InputConfigPath = flag.String("config", "default-config.yaml", "Path to config file")

func main() {
	flag.Parse()

	yamlConfig, err := yd.Decode(*InputConfigPath);
	if err != nil {
		panic(err)
	}

	ValCurs, err := xd.Decode(yamlConfig.InputFile);
	if err != nil {
		panic(err)
	}

	sort_by_value := func(v1, v2 mytypes.Valute) bool {
		return v1.Value < v2.Value
	}

	vl_sort.ValuteSort(ValCurs.Valutes, sort_by_value)

	if err := je.Encode(yamlConfig.OutputFile, ValCurs.Valutes); err != nil {
		panic(err)
	}
}
