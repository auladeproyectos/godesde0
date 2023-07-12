package main

import (
	"fmt"

	"github.com/Auladeproyectos/godesde0/variables"
)

func main() {

	estado, texto := variables.ConviertoaTexto(1677)
	fmt.Println(estado)
	fmt.Println(texto)

}
