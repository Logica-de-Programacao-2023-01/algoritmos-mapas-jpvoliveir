package main

import "fmt"

func contarFrequenciaCaracteres(s string) map[rune]int {
	frequencia := make(map[rune]int)

	for _, char := range s {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	texto := "Hello, world!"

	resultado := contarFrequenciaCaracteres(texto)

	for char, freq := range resultado {
		fmt.Printf("Caractere: %c, Frequência: %d\n", char, freq)
	}
}
