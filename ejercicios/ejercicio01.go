package ejercicios

import (
	"strconv"
)

func StringToInt(word string) (int, string) {
	numero, err := strconv.Atoi(word)
	if err != nil {
		return 0, "Error al convertir a entero."
	}
	var texto string
	if numero > 100 {
		texto = "Es mayor a 100"
	} else {
		texto = "Es menor a 100"
	}
	return numero, texto
}
