package arreglos_slices

import "fmt"

var tablas []int = []int{2, 6, 8}
var arreglo [10]int = [10]int{8, 9, 3, 4, 6, 7, 8, 67}

func Muestroslices() {
	fmt.Println(tablas)

	porcion := arreglo[3:]   //Slices desde un vector desde la posicion 3 al final
	porcion2 := arreglo[:5]  //Slices desde la posicion 0 hasta la 5
	porcion3 := arreglo[6:8] //Slices desde la posicion 6 hasta la 8

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)

}

func Capacidad() {
	elemento := make([]int, 5, 20)

	fmt.Printf("largo %d, Capacidad %d", len(elemento), cap(elemento))

	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {

		nums = append(nums, i)

	}

	fmt.Printf("\nlargo %d, Capacidad %d", len(nums), cap(nums))
}
