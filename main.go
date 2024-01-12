package main

import (
	"fmt"

	"github.com/aaramirez9/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(3)
	fmt.Println(estado)
	fmt.Println(texto)
}
