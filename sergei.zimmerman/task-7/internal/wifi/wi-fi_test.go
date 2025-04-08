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

func TestWiFiService_GetAddresses_Success(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	mockInterfaces := []*wifiMd.Interface{
		{Name: "wlan0", HardwareAddr: net.HardwareAddr{0x00, 0x1A, 0x2B, 0x3C, 0x4D, 0x5E}},
		{Name: "wlan1", HardwareAddr: net.HardwareAddr{0x00, 0x1B, 0x2C, 0x3D, 0x4E, 0x5F}},
	}

	mockWiFi.On("Interfaces").Return(mockInterfaces, nil)

	addrs, err := service.GetAddresses()

	assert.NoError(t, err)
	assert.Len(t, addrs, 2)
	assert.Equal(t, mockInterfaces[0].HardwareAddr, addrs[0])
	assert.Equal(t, mockInterfaces[1].HardwareAddr, addrs[1])
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetAddresses_Error(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	expectedErr := errors.New("failed to retrieve interfaces")
	mockWiFi.On("Interfaces").Return([]*wifiMd.Interface{}, expectedErr)

	addrs, err := service.GetAddresses()

	assert.Error(t, err)
	assert.Nil(t, addrs)
	assert.Equal(t, expectedErr, errors.Unwrap(err))
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Success(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	mockInterfaces := []*wifiMd.Interface{
		{Name: "wlan0"},
		{Name: "wlan1"},
	}

	mockWiFi.On("Interfaces").Return(mockInterfaces, nil)

	names, err := service.GetNames()

	assert.NoError(t, err)
	assert.Len(t, names, 2)
	assert.Equal(t, "wlan0", names[0])
	assert.Equal(t, "wlan1", names[1])
	mockWiFi.AssertExpectations(t)
}

func TestWiFiService_GetNames_Error(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	expectedErr := errors.New("failed to retrieve interfaces")
	mockWiFi.On("Interfaces").Return([]*wifiMd.Interface{}, expectedErr)

	names, err := service.GetNames()

	assert.Error(t, err)
	assert.Nil(t, names)
	assert.Equal(t, expectedErr, errors.Unwrap(err))
	mockWiFi.AssertExpectations(t)
}
