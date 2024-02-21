package main

import (
	"fmt"

	"github.com/exemorelli/golang-proyect/variables"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1612)
	fmt.Println(estado, texto)
}
