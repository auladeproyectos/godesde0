package mapas

import "fmt"

func MostrarMapas() {
	/*
		paises := make(map[string]string)

		paises["Mexico"] = "D.F"
		paises["Argentina"] = "Buenos aires"

		fmt.Println(paises)
		fmt.Println(paises["Mexico"]) */

	campeonato := map[string]int{
		"Barcelona":   7,
		"Real Madrid": 9,
		"Chivas":      6,
		"Nacional":    8,
	}
	fmt.Println(campeonato)
	/*
		for equipo, puntaje := range campeonato {
			fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)

		} */

	delete(campeonato, "Chivas")

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Real Madrid"]

	fmt.Printf("El puntaje captura es %d y el equipo existe %t \n", puntaje, existe)

}
