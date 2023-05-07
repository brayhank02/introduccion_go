package main

import (
	"fmt"
	"runtime"

	"github.com/brayhank02/introduccion_go/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1200)
	fmt.Println(estado, texto)

	os := runtime.GOOS

	if os == "windows" {
		fmt.Println("Operating System:", os)
	} else if os == "mac" {
		fmt.Println("Operating System:", os)
	} else {
		fmt.Println("Operating System:", os)
	}
}
