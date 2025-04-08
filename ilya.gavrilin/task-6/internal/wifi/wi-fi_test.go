package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	myWifi "example_mock/internal/wifi"

	wifiLib "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type wifiTestCase struct {
	macAddrs    []string
	ifaceNames  []string
	errExpected error
}

var wifiTests = []wifiTestCase{
	{
		macAddrs:   []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		ifaceNames: []string{"wlan0", "wlan1"},
	},
	{
		errExpected: errors.New("interface error"),
	},
}

func TestGetAddresses(t *testing.T) {
	mockWiFi := NewWiFi(t) // Generated Mock object
	service := myWifi.New(mockWiFi)

	for i, tc := range wifiTests {
		mockWiFi.ExpectedCalls = nil
		mockWiFi.
			On("Interfaces").
			Return(mockInterfaces(tc.macAddrs, tc.ifaceNames), tc.errExpected).
			Once()

		addrs, err := service.GetAddresses()
		if tc.errExpected != nil {
			require.Error(t, err, "test case %d: expected error", i)
			continue
		}
		require.NoError(t, err)
		require.Equal(t, parseMACs(tc.macAddrs), addrs, "test case %d: addresses mismatch", i)
	}
}

func TestGetNames(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := myWifi.New(mockWiFi)

	for i, tc := range wifiTests {
		mockWiFi.ExpectedCalls = nil
		mockWiFi.
			On("Interfaces").
			Return(mockInterfaces(tc.macAddrs, tc.ifaceNames), tc.errExpected).
			Once()

		names, err := service.GetNames()
		if tc.errExpected != nil {
			require.Error(t, err, "test case %d: expected error", i)
			continue
		}
		require.NoError(t, err)
		require.Equal(t, tc.ifaceNames, names, "test case %d: names mismatch", i)
	}
}

// Similar to sixth lecture
func mockInterfaces(macStrs, names []string) []*wifiLib.Interface {
	var ifaces []*wifiLib.Interface
	for i, macStr := range macStrs {
		hw, _ := net.ParseMAC(macStr)
		iface := &wifiLib.Interface{
			Index:        i + 1,
			Name:         fmt.Sprintf("eth%d", i+1),
			HardwareAddr: hw,
		}
		if i < len(names) {
			iface.Name = names[i]
		}
		ifaces = append(ifaces, iface)
	}
	return ifaces
}

func parseMACs(macStrs []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, s := range macStrs {
		hw, _ := net.ParseMAC(s)
		addrs = append(addrs, hw)
	}
	return addrs
}
