package main

import (
	"fmt"
	"runtime"

	"github.com/julisman/api-gateway-go/server"
)

func main() {

	ConfigRuntime()
    server.StartServer()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}



