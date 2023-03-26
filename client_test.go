package modbus

import (
	"fmt"
	"testing"
	"time"
)

type MyLog struct {
}

func (l *MyLog) Write(dir int, url string, station string, data []byte) {
	fmt.Printf("%d %s %s [% x]\n", dir, url, station, data)
}

func TestClient(t *testing.T) {
	var client *ModbusClient
	var err error

	// for an RTU (serial) device/bus
	client, err = NewClient(&ClientConfiguration{
		URL:      "tcp://127.0.0.1:502",
		Speed:    9600,        // default
		DataBits: 8,           // default, optional
		Parity:   PARITY_NONE, // default, optional
		StopBits: 1,           // default if no parity, optional
		Timeout:  300 * time.Millisecond,
		LSaver:   &MyLog{},
	})

	if err != nil {
		// error out if client creation failed
	}

	// now that the client is created and configured, attempt to connect
	err = client.Open()
	if err != nil {
		// error out if we failed to connect/open the device
		// note: multiple Open() attempts can be made on the same client until
		// the connection succeeds (i.e. err == nil), calling the constructor again
		// is unnecessary.
		// likewise, a client can be opened and closed as many times as needed.
	}

	// read a single 16-bit holding register at address 100
	var reg16 uint16
	reg16, err = client.ReadRegister(1, HOLDING_REGISTER)
	if err != nil {
		// error out
	} else {
		// use value
		fmt.Printf("value: %v\n", reg16)        // as unsigned integer
		fmt.Printf("value: %v\n", int16(reg16)) // as signed integer
	}

	// close the TCP connection/serial port
	client.Close()
}
