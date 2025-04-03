package wifi_test

import (
	"errors"
	"net"
	"testing"

	wifiExternal "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	"example_mock/internal/wifi"
)

//go:generate mockery --testonly --all --quiet --outpkg wifi_test --output .

type testCase struct {
	addrs       []string
	names       []string
	errExpected error
}

var testCases = []testCase{
	{
		addrs:       []string{"12:34:56:78:9a:bc", "de:f0:12:34:56:78"},
		names:       []string{"name1", "name2"},
		errExpected: nil,
	},
	{
		addrs:       nil,
		names:       nil,
		errExpected: errors.New("error"),
	},
}

func getInterfaces(macStrs, names []string) []*wifiExternal.Interface {
	var ifaces []*wifiExternal.Interface
	for i, macStr := range macStrs {
		addr, _ := net.ParseMAC(macStr)
		iface := &wifiExternal.Interface{
			Index:        i,
			Name:         names[i],
			HardwareAddr: addr,
		}
		ifaces = append(ifaces, iface)
	}
	return ifaces
}

func getHWAddrs(macStrs []string) ([]net.HardwareAddr, error) {
	var addrs []net.HardwareAddr
	for _, s := range macStrs {
		addr, err := net.ParseMAC(s)
		if err != nil {
			return nil, err
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

func setupMock(mockWiFi *WiFi, tc testCase) {
	mockWiFi.ExpectedCalls = nil
	mockWiFi.
		On("Interfaces").
		Return(getInterfaces(tc.addrs, tc.names), tc.errExpected).
		Once()
}

func TestGetNames(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	for i, tc := range testCases {
		setupMock(mockWiFi, tc)
		mockWiFi.ExpectedCalls = nil
		mockWiFi.
			On("Interfaces").
			Return(getInterfaces(tc.addrs, tc.names), tc.errExpected).
			Once()

		names, err := service.GetNames()
		if tc.errExpected != nil {
			require.Error(t, err, "test %d: expected error", i)
			continue
		}
		require.NoError(t, err)
		require.Equal(t, tc.names, names, "test %d: names not equal", i)
	}
}

func TestGetAddresses(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	for i, tc := range testCases {
		setupMock(mockWiFi, tc)

		resultAddrs, err := service.GetAddresses()
		if tc.errExpected != nil {
			require.Error(t, err, "test %d: expected error", i)
			continue
		}
		require.NoError(t, err)
		addrs, err := getHWAddrs(tc.addrs)
		if err != nil {
			t.Fatalf("test %d: parsing address error: %v", i, tc.addrs)
		}
		require.Equal(t, addrs, resultAddrs, "test %d: addresses not equal", i)
	}
}
