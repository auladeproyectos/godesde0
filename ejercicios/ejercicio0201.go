package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Tablademultiplicar() string {

	var numero int
	var err error
	var texto string

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

		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, numero*i)

	}

	return texto

}
