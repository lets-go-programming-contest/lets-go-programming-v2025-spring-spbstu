package main

import (
	"errors"
	"fmt"

	"github.com/mdlayher/wifi"

	myWifi "example_mock/internal/wifi"
)

var (
	errGetAddressesFailed = errors.New("failed to get the addresses")
	errCliemtCreateFailed = errors.New("failed creating wifiClient")
)

func main() {
	wifiClient, err := wifi.New()
	if err != nil {
		panic(errors.Join(errCliemtCreateFailed, err))
	}

	wifiService := myWifi.New(wifiClient)

	addrs, err := wifiService.GetAddresses()
	if err != nil {
		panic(errors.Join(errGetAddressesFailed, err))
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
