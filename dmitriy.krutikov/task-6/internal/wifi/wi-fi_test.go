package wifi_test

import (
	"errors"
	"net"
	"testing"

	wifiExternal "github.com/mdlayher/wifi"
	"github.com/stretchr/testify/require"

	"example_mock/internal/wifi"
)

//go:generate mockery --all --testonly --quiet --outpkg wifi_test --output .

type testCase struct {
	addrs       []string
	names       []string
	errExpected error
}


var testTable = []testCase{ 
	{ 
		addrs: 		[]string{"00:11:22:33:44:55", "aa:bb:cc:dd:ee:ff"},
		names: 		[]string{"AA", "cd"}, 
		errExpected: nil,

	}, 
	{ 
		addrs: 		nil,
		names: 		nil, 
		errExpected: errors.New("error"),
	}, 
} 


func mockIfaces(addrs, names []string) []*wifiExternal.Interface { 
	var interfaces []*wifiExternal.Interface 
  
	for i, addrStr := range addrs { 
	   hwAddr, _ := net.ParseMAC(addrStr)   
	   iface := &wifiExternal.Interface{ 
		  Index:        i + 1, 
		  Name:         names[i], 
		  HardwareAddr: hwAddr, 
		  PHY:          1, 
		  Device:       1, 
	   } 
	   interfaces = append(interfaces, iface) 
	} 
  
	return interfaces 
 } 

func parseMACs(macStr []string) ([]net.HardwareAddr, error) { 
	var addrs []net.HardwareAddr 
  
	for _, str := range macStr { 
		addr, err := net.ParseMAC(str)
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, addr)
	} 
  
	return addrs, nil
}

func TestGetNames(t *testing.T) {
	mockWiFi := NewWiFi(t)
	service := wifi.New(mockWiFi)

	for i, tc := range testTable {
		mockWiFi.ExpectedCalls = nil
		mockWiFi.
			On("Interfaces").
			Return(mockIfaces(tc.addrs, tc.names), tc.errExpected).
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

	for i, tc := range testTable {
		mockWiFi.ExpectedCalls = nil
		mockWiFi.
			On("Interfaces").
			Return(mockIfaces(tc.addrs, tc.names), tc.errExpected).
			Once()

		resultAddrs, err := service.GetAddresses()
		if tc.errExpected != nil {
			require.Error(t, err, "test %d: expected error", i)
			continue
		}

		require.NoError(t, err)
		addrs, err := parseMACs(tc.addrs)
		if err != nil {
			t.Fatalf("test %d: parsing address error: %v", i, tc.addrs)
		}
		require.Equal(t, addrs, resultAddrs, "test %d: addresses not equal", i)
	}
}