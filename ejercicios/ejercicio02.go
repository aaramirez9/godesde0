package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int

var err error

func TablaMultiplica() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ingrese numero: ")

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		for i := 1; i <= 10; i++ {
			if err != nil {
				fmt.Println("Error")
				TablaMultiplica()
			}
			fmt.Printf("%d x %d = %d \n", numero1, i, numero1*i)
		}

	}

}
