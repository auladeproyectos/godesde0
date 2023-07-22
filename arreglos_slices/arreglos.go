package arreglos_slices

import "fmt"

var tabla [10]int
var matriz [20][30]int

func MuestroArreglos() {
	tabla[2] = 3
	tabla[6] = 34
	tabla[8] = 345

	tabla2 := [10]string{"Juan", "Pablo", "Carlos"}

	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	matriz[8][15] = 56

	fmt.Println(matriz)
}
