package wifi

import (
	"fmt"
	"net"

	"github.com/mdlayher/wifi"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi WiFi
}

func New(wifi WiFi) Service {
	return Service{WiFi: wifi}
}

func (service Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error in Interfaces: %w", err)
	}

	capacity := 10
	length := 0
	addrs := make([]net.HardwareAddr, length, capacity)

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error in Interfaces: %w", err)
	}

	nameList := make([]string, 0)
	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}

	return nameList, nil
}
