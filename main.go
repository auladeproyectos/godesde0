package main

import (
	e "github.com/Auladeproyectos/godesde0/ejer_interface"
	 "github.com/Auladeproyectos/godesde0/modelos"
)

func main() {

	/*estado, texto := variables.ConviertoaTexto(1677)
	fmt.Println(estado)
	fmt.Println(texto)

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

	numero, texto := ejercicios.Convertirstring("99")

	fmt.Println(numero)
	fmt.Println(texto)

	teclado.Ingresonumeros()

	iteraciones.Iterar()

	ejercicios.NumeroTeclado()

	fmt.Println(ejercicios.Tablademultiplicar())*/

	//files.GrabarTabla()
	//files.SumaTabla()

	//files.Leoarchivo()

	//funciones.Calculos()
	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slices.MuestroArreglos()

	//arreglos_slices.Muestroslices()
	//arreglos_slices.Capacidad()

	//mapas.MostrarMapas()

	// users.Altausuario()

	Pedro:= new(modelos.Hombre)

	e.HumanosRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanosRespirando(Maria)


}
