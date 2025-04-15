package wifi_test

import (
	"errors"
	"net"
	"testing"

	wifiExternal "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"

	"github.com/denisboborukhin/testing/internal/wifi"
)

//go:generate mockery --testonly --all --quiet --outpkg wifi_test --output .

type testCase struct {
	addrs       []string
	names       []string
	expectedErr error
}

var testCases = []testCase{
	{
		addrs:       []string{"11:11:11:11:11:11", "00:00:11:11:22:22"},
		names:       []string{"test1", "test2"},
		expectedErr: nil,
	},
	{
		addrs:       nil,
		names:       nil,
		expectedErr: errors.New("error"),
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
	mockWiFi.On("Interfaces").Return(getInterfaces(tc.addrs, tc.names), tc.expectedErr).Once()
}

func TestGetAddresses(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	for i, tc := range testCases {
		setupMock(mockWiFi, tc)

		resultAddrs, err := service.GetAddresses()
		if tc.expectedErr != nil {
			assert.Error(t, err, "test %d: expected error", i)
			continue
		}
		assert.NoError(t, err)
		addrs, err := getHWAddrs(tc.addrs)
		if err != nil {
			t.Fatalf("test %d: parsing address error: %v", i, tc.addrs)
		}
		assert.Equal(t, addrs, resultAddrs, "test %d: addresses not equal", i)
	}
}

func TestGetNames(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	for i, tc := range testCases {
		setupMock(mockWiFi, tc)

		names, err := service.GetNames()
		if tc.expectedErr != nil {
			assert.Error(t, err, "test %d: expected error", i)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, tc.names, names, "test %d: names not equal", i)
	}
}
