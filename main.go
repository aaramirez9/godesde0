package main

import (
	"fmt"
	"runtime"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	//estado, texto := variables.ConviertoaTexto(3)
	//fmt.Println(estado)
	//fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux" || os == "OS x" {
		fmt.Println("Esto no es windows ", os)
	} else {
		fmt.Println("Esto es windows ", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux ")
	case "darwin":
		fmt.Println("Esto es darwin ")
	default:
		fmt.Printf("%s \n", os)
	}
}
