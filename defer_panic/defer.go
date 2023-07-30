package defer_panic

import (
	"fmt"
	"log"
)

func VerDefer() {
	fmt.Println("Este es el primer mensaje")
	defer fmt.Println("Este es el ultimo mensaje")
	fmt.Println("Este es el segundo mensaje")
}

func VerPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó un panic \n %v", reco)
		}

	}()

	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	}

}
