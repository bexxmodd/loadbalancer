package main

import (
	"fmt"

	"server"
)

func main() {
	fmt.Println("Hello Load Balancer")
	server.Serve()
}
