package main

import (
	"fmt"
	"github.com/goburrow/modbus"
	"github.com/tbrandon/mbserver"
	"time"
)

var serverPort = "0.0.0.0:6161"

func main() {

	serv := mbserver.NewServer()
	err := serv.ListenTCP(serverPort)
	if err != nil {
		fmt.Println("%v\n", err)
	}
	defer serv.Close()

	setRegisters()

	// keep server up
	for {
		time.Sleep(2 * time.Second)
	}
}

func setRegisters() {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler(serverPort)
	err := handler.Connect()
	if err != nil {
		fmt.Println("%v\n", err)
		return
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(0, 3, []byte{0, 3, 0, 4, 0, 5})
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	client.WriteSingleRegister(0, 34)
	client.WriteSingleRegister(1, 53)
	client.WriteSingleRegister(2, 55)
	client.WriteSingleRegister(61, 61)

	results, err := client.ReadHoldingRegisters(0, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results %v\n", results)

	results1, err := client.ReadHoldingRegisters(0, 1)
	fmt.Println(results1)

	results2, _ := client.ReadHoldingRegisters(1, 1)
	fmt.Println(results2)

	results3, _ := client.ReadHoldingRegisters(2, 1)
	fmt.Println(results3)

	results4, _ := client.ReadHoldingRegisters(61, 1)
	fmt.Println(results4)

}
