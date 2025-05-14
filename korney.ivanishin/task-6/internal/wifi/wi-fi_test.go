package wifi_test

import (
	"errors"
	myWifi "example_mock/internal/wifi"
	"fmt"
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate ~/go/bin/mockery --all --testonly --quiet --outpkg wifi_test --output .

type rowTestGetAddresses struct {
        addrs       []string
        errExpected error
}

var testTableGetAdresses = []rowTestGetAddresses{
        {
                addrs:       []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
                errExpected: nil,
        },
        {
                addrs:       nil,
                errExpected: errors.New("ExpectedError"),
        },
}

func TestGetAddresses(t *testing.T) {
        mockWifi := NewWiFi(t)
        wifiService := myWifi.Service{WiFi: mockWifi}

        for i, row := range testTableGetAdresses {
                mockWifi.On("Interfaces").Unset()
                mockWifi.On("Interfaces").Return(mockIfacesAddrs(row.addrs), row.errExpected)

                actualAddrs, err := wifiService.GetAddresses()

                if row.errExpected != nil {
                        require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %s, actual error: %s", i, row.errExpected.Error(), err.Error())
                        continue
                }

                require.NoError(t, err, "row: %d, error must be nil", i)
                require.Equal(t, parseMACs(row.addrs), actualAddrs, "row: %d, expected addrs: %s, actual addrs: %s", i, parseMACs(row.addrs), actualAddrs)
        }
}

func parseMACs(macStr []string) []net.HardwareAddr {
        var addrs []net.HardwareAddr
        for _, addr := range macStr {
                addrs = append(addrs, parseMAC(addr))
        }

        return addrs
}

func parseMAC(macStr string) net.HardwareAddr {
        hwAddr, err := net.ParseMAC(macStr)
        if err != nil {
                return nil
        }

        return hwAddr
}

func mockIfacesAddrs(addrs []string) []*wifi.Interface {
        var interfaces []*wifi.Interface

        for i, addrStr := range addrs {
                hwAddr := parseMAC(addrStr)
                if hwAddr == nil {
                        continue
                }

                iface := &wifi.Interface{
                        Index: i + 1,
                        Name: fmt.Sprintf("eth%d", i+1),
                        HardwareAddr: hwAddr,
                        PHY: 1,
                        Device: 1,
                        Type: wifi.InterfaceTypeAPVLAN,
                        Frequency: 0,
                }

                interfaces = append(interfaces, iface)
        }

        return interfaces
}

type rowTestGetNames struct {
        names       []string
        errExpected error
}

var testTableGetNames = []rowTestGetNames{
        {
                names:       []string{"name1", "name2"},
                errExpected: nil,
        },
        {
                names:       nil,
                errExpected: errors.New("ExpectedError"),
        },
}

func TestGetNames(t *testing.T) {
        mockWifi := NewWiFi(t)
        wifiService := myWifi.Service{WiFi: mockWifi}

        for i, row := range testTableGetNames {
                mockWifi.On("Interfaces").Unset()
                mockWifi.On("Interfaces").Return(mockIfacesNames(row.names), row.errExpected)

                actualNames, err := wifiService.GetNames()

                if row.errExpected != nil {
                        require.ErrorIs(t, err, row.errExpected, "row: %d, expected error: %s, actual error: %s", i, row.errExpected.Error(), err.Error())
                        continue
                }

                require.NoError(t, err, "row: %d, error must be nil", i)
                require.Equal(t, row.names, actualNames, "row: %d, expected names: %s, actual names: %s", i, row.names, actualNames)
        }
}

func mockIfacesNames(names []string) []*wifi.Interface {
        var interfaces []*wifi.Interface

        for i, nameStr := range names {
                iface := &wifi.Interface{
                        Index: i + 1,
                        Name: nameStr,
                        HardwareAddr: net.HardwareAddr{},
                        PHY: 1,
                        Device: 1,
                        Type: wifi.InterfaceTypeAPVLAN,
                        Frequency: 0,
                }

                interfaces = append(interfaces, iface)
        }

        return interfaces
}

