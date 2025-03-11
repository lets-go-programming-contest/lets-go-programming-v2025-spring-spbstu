package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"task-3/internal/cbr"
	"task-3/internal/config"
	"task-3/internal/shortcurrency"
)

func main() {
	configPath := flag.String("config" /*default=*/, "config.yaml" /* description */, "path to the config.yaml")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintf(w, "usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Sprintf("error loading config: %v", err))
	}

	cbrdata, err := cbr.ParseCbrXML(cfg.InputFile)
	if err != nil {
		panic(fmt.Sprintf("error parsing xml: %v", err))
	}

	currencies := cbrdata.Currencies

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	err = shortcurrency.WriteJSON(cfg.OutputFile, currencies)
	if err != nil {
		panic(fmt.Sprintf("error writing json: %v", err))
	}
}
