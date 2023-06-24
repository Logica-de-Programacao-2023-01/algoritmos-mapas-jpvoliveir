package main

import (
	"fmt"
	"strings"
)

func contarOcorrenciasPalavras(texto string) map[string]int {
	// Remover pontuações e converter para minúsculas
	texto = strings.ToLower(texto)
	texto = strings.ReplaceAll(texto, ",", "")
	texto = strings.ReplaceAll(texto, ".", "")
	texto = strings.ReplaceAll(texto, "!", "")
	texto = strings.ReplaceAll(texto, "?", "")

	// Dividir o texto em palavras
	palavras := strings.Fields(texto)

	// Contagem das ocorrências de cada palavra
	ocorrencias := make(map[string]int)
	for _, palavra := range palavras {
		ocorrencias[palavra]++
	}

	return ocorrencias
}

func main() {
	texto := "Olá, mundo! Olá, programação. Olá, Go!"
	resultado := contarOcorrenciasPalavras(texto)
	fmt.Println(resultado)
}
