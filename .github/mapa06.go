package main

import "fmt"

func somarContagens(listasDeContagens []map[string]int) map[string]int {
	resultado := make(map[string]int)

	for _, contagem := range listasDeContagens {
		for palavra, quantidade := range contagem {
			resultado[palavra] += quantidade
		}
	}

	return resultado
}

func main() {
	listaContagens := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 3, "world": 2, "golang": 1},
		{"hello": 1, "golang": 2},
	}

	resultado := somarContagens(listaContagens)

	fmt.Println(resultado)
}
