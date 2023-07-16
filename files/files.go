package files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Auladeproyectos/godesde0/ejercicios"
)

var filename string = "./files/txt/tablas.txt"

func GrabarTabla() {

	var texto string = ejercicios.Tablademultiplicar()
	archivo, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.Tablademultiplicar()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar el archivo")
	}
}

func Append(filen string, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println("Error durante el Appened" + err.Error())
		return false
	}
	_, err = arch.WriteString(texto)

	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}

	arch.Close()
	return true

}

func Leoarchivo2() {
	archivo, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error durante la lectura del archivo" + err.Error())
		return
	}

	fmt.Println(string(archivo))
}

func Leoarchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error durante la lectura del archivo" + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}

	archivo.Close()
}
