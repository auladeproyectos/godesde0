package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia

	}
}

func LlamarClosure() {
	tablade1 := 2
	Mitabla := tabla(tablade1)

	for i := 1; i <= 10; i++ {
		fmt.Println(Mitabla())
	}
}
