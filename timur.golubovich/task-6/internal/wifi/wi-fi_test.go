package wifi_test

import (
	"errors"
	"net"
	"testing"

	"example_mock/internal/wifi"

	wifiMd "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
)

//go:generate mockery --name=WiFi --outpkg wifi_test --output . --testonly

// Test suite for WiFiService
type WiFiServiceTestSuite struct {
	mockWiFi *WiFi
	service  *wifi.Service
}

// Initializes the test suite
func setupWiFiServiceTest(t *testing.T) *WiFiServiceTestSuite {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)
	return &WiFiServiceTestSuite{mockWiFi: mockWiFi, service: &service}
}

// Test for successful retrieval of hardware addresses
func TestWiFiService_GetAddresses_Success(t *testing.T) {
	suite := setupWiFiServiceTest(t)

	mockInterfaces := []*wifiMd.Interface{
		{Name: "wlan0", HardwareAddr: net.HardwareAddr{0xFF, 0xEE, 0xDD, 0xCC, 0xBB, 0xAA}},
		{Name: "wlan1", HardwareAddr: net.HardwareAddr{0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF}},
	}

	suite.mockWiFi.On("Interfaces").Return(mockInterfaces, nil)

	addresses, err := suite.service.GetAddresses()

	assert.NoError(t, err, "Expected no error when retrieving addresses")
	assert.Len(t, addresses, len(mockInterfaces), "Expected the number of addresses to match mock interfaces")
	for i, iface := range mockInterfaces {
		assert.Equal(t, iface.HardwareAddr, addresses[i], "Expected address to match for interface %s", iface.Name)
	}
	suite.mockWiFi.AssertExpectations(t)
}

// Test for error during retrieval of hardware addresses
func TestWiFiService_GetAddresses_Error(t *testing.T) {
	suite := setupWiFiServiceTest(t)

	expectedErr := errors.New("failed to retrieve interfaces")
	suite.mockWiFi.On("Interfaces").Return([]*wifiMd.Interface{}, expectedErr)

	addresses, err := suite.service.GetAddresses()

	assert.Error(t, err, "Expected an error when retrieving addresses")
	assert.Nil(t, addresses, "Expected addresses to be nil on error")
	assert.Equal(t, expectedErr, errors.Unwrap(err), "Expected the unwrapped error to match")
	suite.mockWiFi.AssertExpectations(t)
}

// Test for successful retrieval of interface names
func TestWiFiService_GetNames_Success(t *testing.T) {
	suite := setupWiFiServiceTest(t)

	mockInterfaces := []*wifiMd.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
	}

	suite.mockWiFi.On("Interfaces").Return(mockInterfaces, nil)

	names, err := suite.service.GetNames()

	assert.NoError(t, err, "Expected no error when retrieving names")
	assert.Len(t, names, len(mockInterfaces), "Expected the number of names to match mock interfaces")
	for i, iface := range mockInterfaces {
		assert.Equal(t, iface.Name, names[i], "Expected name to match for interface %s", iface.Name)
	}
	suite.mockWiFi.AssertExpectations(t)
}

// Test for error during retrieval of interface names
func TestWiFiService_GetNames_Error(t *testing.T) {
	suite := setupWiFiServiceTest(t)

	expectedErr := errors.New("failed to retrieve interfaces")
	suite.mockWiFi.On("Interfaces").Return([]*wifiMd.Interface{}, expectedErr)

	names, err := suite.service.GetNames()

	assert.Error(t, err, "Expected an error when retrieving names")
	assert.Nil(t, names, "Expected names to be nil on error")
	assert.Equal(t, expectedErr, errors.Unwrap(err), "Expected the unwrapped error to match")
	suite.mockWiFi.AssertExpectations(t)
}
