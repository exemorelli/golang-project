package main

import (
	"fmt"
	//"github.com/exemorelli/golang-proyect/variables"
	"runtime"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	/*estado, texto := variables.ConviertoaTexto(1612)
	fmt.Println(estado, texto)*/

	if os := runtime.GOOS; os == "linux" || os == "OS X." {
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
	}

}
