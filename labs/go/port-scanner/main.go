package main

import (
	"fmt"
	"portscanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	open := port.ScanPort("tcp", "localhost", 9090)
	fmt.Printf("Port Open: %t\n", open)
}
