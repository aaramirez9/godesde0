package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	seuencia := 0
	return func() int {
		seuencia++
		return numero * seuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	MiTabla := tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(MiTabla())
	}
}
