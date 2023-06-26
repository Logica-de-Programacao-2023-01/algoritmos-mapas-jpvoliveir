package main

import "fmt"

func somarValoresMapa(m map[string]int) int {
	soma := 0
	for _, valor := range m {
		soma += valor
	}
	return soma
}

func main() {
	mapa := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	total := somarValoresMapa(mapa)
	fmt.Println(total)
}
