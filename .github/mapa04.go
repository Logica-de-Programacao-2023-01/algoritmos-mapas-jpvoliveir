package main

import (
	"fmt"
	"sort"
	"strings"
)

func ordenarString(s string) string {
	// Converter a string em um slice de caracteres
	chars := strings.Split(s, "")

	// Ordenar os caracteres em ordem alfabética
	sort.Strings(chars)

	// Juntar os caracteres ordenados em uma única string
	return strings.Join(chars, "")
}

func agruparAnagramas(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		// Obter a versão ordenada da palavra como chave para o mapa
		chave := ordenarString(palavra)

		// Adicionar a palavra ao grupo de anagramas correspondente na chave do mapa
		anagramas[chave] = append(anagramas[chave], palavra)
	}

	return anagramas
}

func main() {
	palavras := []string{"amor", "roma", "carro", "arco", "casa", "saco"}

	resultado := agruparAnagramas(palavras)

	for _, grupo := range resultado {
		fmt.Println(grupo)
	}
}
