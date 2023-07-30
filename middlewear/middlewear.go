package middlewear

import "fmt"

func sumar(a, b int) int {
	return a + b

}

func restar(a, b int) int {
	return a - b

}

func multiplicar(a, b int) int {
	return a * b

}

func Mimiddleware() {
	fmt.Println("Inicio")

	result := OperacionesMidd(sumar)(2, 3)
	fmt.Println(result)

	result = OperacionesMidd(restar)(10, 3)
	fmt.Println(result)

	result = OperacionesMidd(multiplicar)(2, 4)
	fmt.Println(result)
}

func OperacionesMidd(f func(int, int) int) func(int, int) int {

	return func(a, b int) int {
		fmt.Println("Inicio de operaciones")
		return f(a, b)

	}
}
