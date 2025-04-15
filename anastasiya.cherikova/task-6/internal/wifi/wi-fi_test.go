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

// testCase defines input data and expected results for WiFi interface tests
type testCase struct {
	expectedMACs  []string // List of expected MAC addresses
	expectedNames []string // List of expected interface names
	expectedError error    // Expected error (if any)
}

// testScenarios contains all test cases for WiFi interface operations
var testScenarios = []testCase{
	{
		expectedMACs:  []string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		expectedNames: []string{"wlan0", "wlan1"},
	},
	{
		expectedError: errors.New("interface error"),
	},
}

// TestService_GetMACAddresses verifies correct extraction of MAC addresses from WiFi interfaces
func TestService_GetMACAddresses(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := myWifi.New(mockWiFi)

	for scenarioID, tc := range testScenarios {
		t.Run(fmt.Sprintf("Scenario%d", scenarioID), func(t *testing.T) {
			// Reset mock expectations between test cases
			mockWiFi.ExpectedCalls = nil

			// Configure mock behavior
			mockWiFi.On("Interfaces").
				Return(createMockInterfaces(tc.expectedMACs, tc.expectedNames), tc.expectedError).
				Once()

			// Execute method under test
			addrs, err := service.GetAddresses()

			// Validate results
			if tc.expectedError != nil {
				require.ErrorContains(t, err, tc.expectedError.Error(),
					"Scenario %d: Expected error missing", scenarioID)
				return
			}

			require.NoError(t, err, "Scenario %d: Unexpected error", scenarioID)
			require.ElementsMatch(t, parseMACStrings(tc.expectedMACs), addrs,
				"Scenario %d: MAC address mismatch", scenarioID)
		})
	}
}

// TestService_GetInterfaceNames validates proper retrieval of WiFi interface names
func TestService_GetInterfaceNames(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := myWifi.New(mockWiFi)

	for scenarioID, tc := range testScenarios {
		t.Run(fmt.Sprintf("Scenario%d", scenarioID), func(t *testing.T) {
			// Reset mock expectations between test cases
			mockWiFi.ExpectedCalls = nil

			// Configure mock behavior
			mockWiFi.On("Interfaces").
				Return(createMockInterfaces(tc.expectedMACs, tc.expectedNames), tc.expectedError).
				Once()

			// Execute method under test
			names, err := service.GetNames()

			// Validate results
			if tc.expectedError != nil {
				require.ErrorContains(t, err, tc.expectedError.Error(),
					"Scenario %d: Expected error missing", scenarioID)
				return
			}

			require.NoError(t, err, "Scenario %d: Unexpected error", scenarioID)
			require.Equal(t, tc.expectedNames, names,
				"Scenario %d: Interface name mismatch", scenarioID)
		})
	}
}

// createMockInterfaces generates mock WiFi interface objects from MAC strings and names
func createMockInterfaces(macStrings []string, names []string) []*wifiLib.Interface {
	var interfaces []*wifiLib.Interface

	for idx, macStr := range macStrings {
		// Create base interface with default values
		iface := &wifiLib.Interface{
			Index:        idx + 1,
			Name:         fmt.Sprintf("eth%d", idx+1), // Default name
			HardwareAddr: parseMAC(macStr),
		}

		// Override name if provided
		if idx < len(names) && names[idx] != "" {
			iface.Name = names[idx]
		}

		interfaces = append(interfaces, iface)
	}

	return interfaces
}

// parseMAC converts MAC address string to net.HardwareAddr
func parseMAC(macStr string) net.HardwareAddr {
	hwAddr, err := net.ParseMAC(macStr)
	if err != nil {
		panic(fmt.Sprintf("invalid MAC address in test data: %s", macStr))
	}
	return hwAddr
}

// parseMACStrings converts slice of MAC strings to HardwareAddr objects
func parseMACStrings(macStrings []string) []net.HardwareAddr {
	var addrs []net.HardwareAddr
	for _, s := range macStrings {
		addrs = append(addrs, parseMAC(s))
	}
	return addrs
}
