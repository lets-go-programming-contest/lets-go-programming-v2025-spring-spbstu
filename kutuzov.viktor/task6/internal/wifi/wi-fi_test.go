package wifi_test

import (
	"errors"
	"example_mock/internal/wifi"
	"fmt"
	"net"
	"testing"

	wifiExternal "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

//go:generate mockery --testonly --all --quiet --outpkg wifi_test --output .

type testCase struct {
	name        string
	macStrings  []string // Input MAC addresses as strings
	names       []string // Expected interface names
	expectError error    // Expected error
}

type TestSuite struct {
	suite.Suite
	testCases []testCase
}

func TestWifiSuite(t *testing.T) {
	suite.Run(t, &TestSuite{
		testCases: []testCase{
			{
				name:        "happy path",
				macStrings:  []string{"11:22:33:44:55:66", "aa:bb:cc:dd:ee:ff"},
				names:       []string{"name1", "name2"},
				expectError: nil,
			},
			{
				name:        "error case",
				macStrings:  nil,
				names:       nil,
				expectError: errors.New("mock error"),
			},
		},
	})
}

func createInterfaces(macStrings, names []string) ([]*wifiExternal.Interface, error) {
	if len(macStrings) != len(names) {
		return nil, fmt.Errorf("input length mismatch: %d MACs vs %d names",
			len(macStrings), len(names))
	}

	ifaces := make([]*wifiExternal.Interface, 0, len(macStrings))
	for i, macStr := range macStrings {
		mac, err := net.ParseMAC(macStr)
		if err != nil {
			return nil, fmt.Errorf("invalid MAC %q: %w", macStr, err)
		}

		ifaces = append(ifaces, &wifiExternal.Interface{
			Index:        i,
			Name:         names[i],
			HardwareAddr: mac,
		})
	}
	return ifaces, nil
}

func (ts *TestSuite) setupMock(mock *WiFi, tc testCase) {
	mock.ExpectedCalls = nil //Cleans earlier expectations calls for mock to avoid leakage between test

	ifaces, _ := createInterfaces(tc.macStrings, tc.names)
	mock.On("Interfaces").Return(ifaces, tc.expectError).Once()
}

func (ts *TestSuite) TestGetNames() {
	for _, tc := range ts.testCases {
		tc := tc
		ts.T().Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Setup
			mock := NewWiFi(t)
			ts.setupMock(mock, tc)
			service := wifi.New(mock)

			// Execute
			names, err := service.GetNames()

			// Verify
			if tc.expectError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectError.Error())
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.names, names)
		})
	}
}

func (ts *TestSuite) TestGetAddresses() {
	for _, tc := range ts.testCases {
		tc := tc
		ts.T().Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Setup
			mock := NewWiFi(t)
			ts.setupMock(mock, tc)
			service := wifi.New(mock)

			// Execute
			addrs, err := service.GetAddresses()

			// Verify
			if tc.expectError != nil {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectError.Error())
				return
			}

			require.NoError(t, err)

			expectedAddrs := make([]net.HardwareAddr, 0, len(tc.macStrings))
			for _, s := range tc.macStrings {
				mac, pErr := net.ParseMAC(s)
				require.NoError(t, pErr, "invalid test data")
				expectedAddrs = append(expectedAddrs, mac)
			}

			require.Equal(t, expectedAddrs, addrs)
		})
	}
}
