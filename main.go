package main

import (
	"fmt"

	"github.com/aaramirez9/godesde0/goroutines"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	//estado, texto := variables.ConviertoaTexto(3)
	//fmt.Println(estado)
	//fmt.Println(texto)

	//	if os := runtime.GOOS; os == "Linux" || os == "OS x" {
	//		fmt.Println("Esto no es windows ", os)
	//	} else {
	//		fmt.Println("Esto es windows ", os)
	//	}

	//	switch os := runtime.GOOS; os {
	//	case "linux":
	//		fmt.Println("Esto es linux ")
	//	case "darwin":
	//		fmt.Println("Esto es darwin ")
	//	default:
	//		fmt.Printf("%s \n", os)
	//	}

	//numero, entero := ejercicios.ConviertoaEntero("alexa")
	//fmt.Println(numero)
	//fmt.Println(entero)

	// Entero, mensaje := ejercicios.ConviertoaEntero("299")

	// fmt.Printf("hola %s \n", mensaje)
	// fmt.Printf("hola %d", Entero)
	//teclado.IngresoNumeros()

	//iteraciones.Iterar()
	//fmt.Println(ejercicios.TablaMultiplica())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.LlamarClosure()
	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglo()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUsuario()
	//Pedro := new(modelos.Hombre)
	//e.HumanosRespirando(Pedro)

	//Maria := new(modelos.Mujer)
	//e.HumanosRespirando(Maria)
	//defer_panic.VemosDefer()
	//defer_panic.EjemploPanic()
	canal1 := make(chan bool)
	go goroutines.MiNombreLentooo("Alexandra A", canal1)

	fmt.Println("Estoy aquí")
	<-canal1

}
