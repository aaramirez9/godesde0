package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplica() string {
	scanner := bufio.NewScanner(os.Stdin)

	var numero1 int

	var err error
	var texto string
	for {
		fmt.Println("Ingrese numero: ")
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}

	}

	for i := 1; i <= 10; i++ {
		texto += fmt.Sprintf("%d x %d = %d \n", numero1, i, numero1*i)
	}
	return texto
}
