package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	var startPort int
	var endPort int
	var verboseMode string
	var verboseOn bool
	net

	fmt.Print("Let's scan some ports.\nEnter starting port number: ")
	fmt.Scan(&startPort)
	fmt.Print("Enter the ending port number:")
	fmt.Scan(&endPort)
	fmt.Print("Show errors? (y/) ")
	fmt.Scan(&verboseMode)
	if strings.ToLower(verboseMode) == "y" {
		verboseOn = true
	} else {
		fmt.Print("Not displaying errors.\n")
	}

	fmt.Print("Starting port scan...\n")
	for i := startPort; i <= endPort; i++ {
		port := strconv.FormatInt(int64(i), 10)
		conn, err := net.Dial("tcp", "10.0.0.1:"+port)
		if err == nil {
			fmt.Println("Port", i, "is open")
			conn.Close()
		} else if verboseOn == true {
			fmt.Print("Port Number: port", err, "\n")
		}
	}
	fmt.Print("=^_^= Request complete. You deserve coffee. |_|>")
}
