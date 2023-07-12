package main

import (
	"fmt"
	"runtime"
)

func main() {

	/*estado, texto := variables.ConviertoaTexto(1677)
	fmt.Println(estado)
	fmt.Println(texto) */

	os := runtime.GOOS

	if os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows")
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Esto es Linux")

	case "Darwin":
		fmt.Println("Esto es darwin")

	default:
		fmt.Printf("%s \n", os)

	}
}
