package main

import (
	"fmt"
	"os"
)

func main() {

	var hostname string
	var err error
	hostname, err = os.Hostname()

	fmt.Printf("Operating System Hostname %s", hostname)

	if err != nil {
		fmt.Printf("Error trying to obtain system hostname %s", err)
	}
}
