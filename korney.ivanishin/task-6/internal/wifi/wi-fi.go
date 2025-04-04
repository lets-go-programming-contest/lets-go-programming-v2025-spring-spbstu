package wifi

import (
	"errors"
	"net"

	"github.com/mdlayher/wifi"
)

type WiFi interface {
	Interfaces() ([]*wifi.Interface, error)
}

type Service struct {
	WiFi WiFi
}

var (
	errGetAddressesFailed = errors.New("failed getting adresses")
	errGetNamesFailed     = errors.New("failed getting names")
)

func New(wifi WiFi) Service {
	return Service{WiFi: wifi}
}

func (service Service) GetAddresses() ([]net.HardwareAddr, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, errors.Join(errGetAddressesFailed, err)
	}

	addrs := make([]net.HardwareAddr, 0)

	for _, iface := range interfaces {
		addrs = append(addrs, iface.HardwareAddr)
	}

	return addrs, nil
}

func (service Service) GetNames() ([]string, error) {
	interfaces, err := service.WiFi.Interfaces()
	if err != nil {
		return nil, errors.Join(errGetNamesFailed, err)
	}

	nameList := make([]string, 0)

	for _, iface := range interfaces {
		nameList = append(nameList, iface.Name)
	}

	return nameList, nil
}
