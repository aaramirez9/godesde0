package ejercicios

import (
	"strconv"
)

func ConviertoaEntero(cadena string) (int, string) {
	var mensaje string
	Entero, _ := strconv.Atoi(cadena)
	if Entero > 100 {
		mensaje = "Es Mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}
	return Entero, mensaje

}
