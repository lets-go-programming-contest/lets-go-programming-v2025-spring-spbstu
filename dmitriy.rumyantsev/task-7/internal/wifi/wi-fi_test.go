package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWifi "github.com/dmitriy.rumyantsev/task-7/internal/wifi"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --name=WiFi --outpkg wifi_test --output . --testonly

func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		return nil
	}
	return hwAddr
}

func mockIfaces(addrs []string, emptyName bool) []*wifi.Interface {
	var interfaces []*wifi.Interface
	for i, addr := range addrs {
		iface := &wifi.Interface{
			Index:        i,
			Name:         fmt.Sprintf("eth%d", i),
			HardwareAddr: parseMAC(addr),
		}
		if emptyName {
			iface.Name = ""
		}
		interfaces = append(interfaces, iface)
	}
	return interfaces
}

func extractMACs(ifs []*wifi.Interface) []net.HardwareAddr {
	var macs []net.HardwareAddr
	for _, iface := range ifs {
		macs = append(macs, iface.HardwareAddr)
	}
	return macs
}

func extractNames(ifs []*wifi.Interface) []string {
	var names []string
	for _, iface := range ifs {
		names = append(names, iface.Name)
	}
	return names
}

func TestWiFiService_GetAddresses(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		ifaces := mockIfaces([]string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"}, false)
		mockWifi.On("Interfaces").Return(ifaces, nil)

		service := myWifi.Service{WiFi: mockWifi}
		result, err := service.GetAddresses()

		require.NoError(t, err)
		require.Equal(t, extractMACs(ifaces), result)
	})

	t.Run("interfaces returns error", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		mockWifi.On("Interfaces").Return(nil, errors.New("fail"))

		service := myWifi.Service{WiFi: mockWifi}
		result, err := service.GetAddresses()

		require.Error(t, err)
		require.Nil(t, result)
	})

	t.Run("empty list", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		mockWifi.On("Interfaces").Return([]*wifi.Interface{}, nil)

		service := myWifi.Service{WiFi: mockWifi}
		result, err := service.GetAddresses()

		require.NoError(t, err)
		require.Empty(t, result)
	})

	t.Run("interface with nil MAC", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		ifaces := []*wifi.Interface{{Name: "eth0", HardwareAddr: nil}}
		mockWifi.On("Interfaces").Return(ifaces, nil)

		service := myWifi.Service{WiFi: mockWifi}
		result, err := service.GetAddresses()

		require.NoError(t, err)
		require.Equal(t, []net.HardwareAddr{nil}, result)
	})
}

func TestWiFiService_GetNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		ifaces := mockIfaces([]string{"00:11:22:33:44:55"}, false)
		mockWifi.On("Interfaces").Return(ifaces, nil)

		service := myWifi.Service{WiFi: mockWifi}
		names, err := service.GetNames()

		require.NoError(t, err)
		require.Equal(t, extractNames(ifaces), names)
	})

	t.Run("error from Interfaces()", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		mockWifi.On("Interfaces").Return(nil, errors.New("test"))

		service := myWifi.Service{WiFi: mockWifi}
		names, err := service.GetNames()

		require.Error(t, err)
		require.Nil(t, names)
	})

	t.Run("empty interface list", func(t *testing.T) {
		mockWifi := NewWiFi(t)
		mockWifi.On("Interfaces").Return([]*wifi.Interface{}, nil)

		service := myWifi.Service{WiFi: mockWifi}
		names, err := service.GetNames()

		require.NoError(t, err)
		require.Empty(t, names)
	})
}
