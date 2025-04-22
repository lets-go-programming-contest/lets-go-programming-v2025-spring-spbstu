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

func (s Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := s.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("GetAddresses: %w", err)
	}

	addrs := make([]net.HardwareAddr, 0, len(interfaces))

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (s Service) GetNames() ([]string, error) {
	interfaces, err := s.WiFi.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("GetNames: %w", err)
	}

	nameList := make([]string, 0, len(interfaces))

	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}

	return nameList, nil
}
