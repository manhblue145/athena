package main

import (
	"fmt"

	"github.com/manhblue145/athena/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
