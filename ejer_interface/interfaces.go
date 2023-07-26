package ejer_interface

import (
	"fmt"

	"github.com/Auladeproyectos/godesde0/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {

	hu.Respirar()

	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())

}
