package main

import "fmt"

func gerarSequenciaFibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	// Caso base: os primeiros dois números da sequência são 0 e 1
	sequencia[0] = 0
	sequencia[1] = 1

	// Gerar os números da sequência até o número n
	for i := 2; sequencia[i-1]+sequencia[i-2] <= n; i++ {
		sequencia[i] = sequencia[i-1] + sequencia[i-2]
	}

	return sequencia
}

func main() {
	n := 100

	resultado := gerarSequenciaFibonacci(n)

	for indice, numero := range resultado {
		fmt.Printf("Índice: %d, Número: %d\n", indice, numero)
	}
}
