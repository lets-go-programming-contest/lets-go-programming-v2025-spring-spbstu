package wifi_test

import (
	"errors"
	"fmt"
	"net"
	"testing"

	wifiExternal "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/denisboborukhin/testing/internal/wifi"
)

//go:generate mockery --testonly --all --quiet --outpkg wifi_test --output .

func getInterfaces(macStrs, names []string) ([]*wifiExternal.Interface, error) {
	if len(macStrs) != len(names) {
		return nil, fmt.Errorf("length mismatch: macStrs has %d elements, names has %d elements", len(macStrs), len(names))
	}

	var ifaces []*wifiExternal.Interface
	for i, macStr := range macStrs {
		addr, err := net.ParseMAC(macStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MAC address: %v", err)
		}
		iface := &wifiExternal.Interface{
			Index:        i,
			Name:         names[i],
			HardwareAddr: addr,
		}
		ifaces = append(ifaces, iface)
	}
	return ifaces, nil
}

func getHWAddrs(macStrs []string) ([]net.HardwareAddr, error) {
	var addrs []net.HardwareAddr
	for _, s := range macStrs {
		addr, err := net.ParseMAC(s)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MAC address: %v", err)
		}
		addrs = append(addrs, addr)
	}
	return addrs, nil
}

func setupMock(mockWiFi *WiFi, macs, names []string, retErr error) error {
	mockWiFi.ExpectedCalls = nil

	ifaces, err := getInterfaces(macs, names)
	if err != nil {
		mockWiFi.On("Interfaces").Return(nil, err)
		return nil
	}

	mockWiFi.On("Interfaces").Return(ifaces, retErr)
	return nil
}

func setupService(t *testing.T) (*WiFi, *wifi.Service) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)
	return mockWiFi, &service
}

func TestGetAddresses(t *testing.T) {
	testCases := []struct {
		name        string
		addrs       []string
		names       []string
		expectedErr error
	}{
		{
			name:        "successful case with multiple interfaces",
			addrs:       []string{"11:11:11:11:11:11", "00:00:11:11:22:22"},
			names:       []string{"test1", "test2"},
			expectedErr: nil,
		},
		{
			name:        "error case",
			addrs:       nil,
			names:       nil,
			expectedErr: errors.New("error"),
		},
		{
			name:        "empty interfaces",
			addrs:       []string{},
			names:       []string{},
			expectedErr: nil,
		},
		{
			name:        "invalid MAC address",
			addrs:       []string{"invalid:mac:address"},
			names:       []string{"test1"},
			expectedErr: errors.New("failed to parse MAC address: address invalid:mac:address: invalid MAC address"),
		},
		{
			name:        "length mismatch",
			addrs:       []string{"11:11:11:11:11:11"},
			names:       []string{"test1", "test2"},
			expectedErr: errors.New("length mismatch: macStrs has 1 elements, names has 2 elements"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockWiFi, service := setupService(t)

			err := setupMock(mockWiFi, tc.addrs, tc.names, tc.expectedErr)
			require.NoError(t, err)

			resultAddrs, err := service.GetAddresses()

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error())
				assert.Nil(t, resultAddrs)
				return
			}

			require.NoError(t, err)

			expectedAddrs, err := getHWAddrs(tc.addrs)
			require.NoError(t, err)

			assert.ElementsMatch(t, expectedAddrs, resultAddrs)
		})
	}
}

func TestGetNames(t *testing.T) {
	testCases := []struct {
		name        string
		addrs       []string
		names       []string
		expectedErr error
	}{
		{
			name:        "successful case with multiple interfaces",
			addrs:       []string{"11:11:11:11:11:11", "00:00:11:11:22:22"},
			names:       []string{"test1", "test2"},
			expectedErr: nil,
		},
		{
			name:        "error case",
			addrs:       nil,
			names:       nil,
			expectedErr: errors.New("error"),
		},
		{
			name:        "empty interfaces",
			addrs:       []string{},
			names:       []string{},
			expectedErr: nil,
		},
		{
			name:        "invalid MAC address",
			addrs:       []string{"invalid:mac:address"},
			names:       []string{"test1"},
			expectedErr: errors.New("failed to parse MAC address: address invalid:mac:address: invalid MAC address"),
		},
		{
			name:        "length mismatch",
			addrs:       []string{"11:11:11:11:11:11"},
			names:       []string{"test1", "test2"},
			expectedErr: errors.New("length mismatch: macStrs has 1 elements, names has 2 elements"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			mockWiFi, service := setupService(t)

			err := setupMock(mockWiFi, tc.addrs, tc.names, tc.expectedErr)
			require.NoError(t, err)

			names, err := service.GetNames()

			if tc.expectedErr != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tc.expectedErr.Error())
				assert.Nil(t, names)
				return
			}

			require.NoError(t, err)

			if len(tc.names) == 0 {
				assert.Nil(t, names)
			} else {
				assert.ElementsMatch(t, tc.names, names)
			}
		})
	}
}
