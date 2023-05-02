package main

import (
	"fmt"

	"github.com/brayhank02/introduccion_go/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1200)
	fmt.Println(estado, texto)
}
