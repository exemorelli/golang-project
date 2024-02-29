package main

import (
	"fmt"
	//"github.com/exemorelli/golang-project/variables"
	"github.com/exemorelli/golang-project/ejercicios"
	// "runtime"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1612)
	fmt.Println(estado, texto)*/

	/*if os := runtime.GOOS; os == "linux" || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

	palabra, mensaje := ejercicios.StringToInt("300")
	fmt.Println(palabra)
	fmt.Println(mensaje)
}
