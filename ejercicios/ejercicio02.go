package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func NumeroTeclado() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese el numero:")

		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())

			if err != nil {
				continue
			} else {
				break
			}
		}

	}
	fmt.Println("Se generara la tabla de multiplicar del numero ", numero)

	for i := 1; i < 11; i++ {

		fmt.Println(numero, " x ", i, "=", numero*i)

	}

}
