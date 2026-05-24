package main

import (
	"fmt"
	"thriftBinaryProtocol-service/service"
)

func main() {
	fmt.Println("Application is runnning")
	service.Start()
}
