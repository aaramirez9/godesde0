package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	//fmt.Println(paises)
	paises["Ecuador"] = "Quito"
	paises["Arrgentina"] = "Buenos Aires"
	//fmt.Println(paises)
	//fmt.Println(paises["Argentina"])

	campeonato := map[string]int{
		"Barcelona": 20,
		"Quito":     50,
		"Nacional":  34,
		"Macará":    84,
	}
	/*fmt.Println(campeonato)
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}*/

	delete(campeonato, "Nacional")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Quito"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe =%t \n", puntaje, existe)
}
