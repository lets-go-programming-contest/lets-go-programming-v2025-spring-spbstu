package main

import (
	"fmt"

	myWifi "github.com/denisboborukhin/testing/internal/wifi"
	"github.com/mdlayher/wifi"
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		fmt.Printf("Ошибка при создании wifiClient: %s\n", err.Error())
		return
	}

	wifiService := myWifi.New(wifiClient)

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		fmt.Printf("Ошибка при получении адресов: %s\n", err.Error())
		return
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
