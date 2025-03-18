package main

import (
	"flag"
	"fmt"

	"task-3/internal/config"
	"task-3/internal/currency"
	"task-3/internal/jsonwriter"
	"task-3/internal/xmlparser"
)

func main() {
	// Обработка флага -config
	configPath := flag.String("config", "", "Путь к конфигурационному файлу")
	flag.Parse()

	if *configPath == "" {
		panic("Не указан путь к конфигурационному файлу!")
	}
	fmt.Println(*configPath)

	// Загрузка конфига
	cfg := config.LoadConfig(*configPath)

	// Парсинг XML, Сортировка, Запись в JSON
	currencies := xmlparser.ParseXML(cfg.InputFile)
	currency.SortCurrencies(currencies)
	jsonwriter.WriteJSON(currencies, cfg.OutputFile)

	fmt.Println("Успешно! Результат сохранён в", cfg.OutputFile)
}
