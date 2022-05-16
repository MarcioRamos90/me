package main

import (
	"fmt"
	"portscanner/port"
)

func main() {
	fmt.Println("Port Scanner in Go")

	open := port.InitialScan("localhost")
	fmt.Println(open)
}
