package main

import (
	"fmt"
)

func calcularDivisaoConta(gastos map[string]float64) map[string]float64 {
	totalGastos := 0.0

	// Calcular o total de gastos
	for _, valor := range gastos {
		totalGastos += valor
	}

	// Calcular o valor que cada pessoa deve receber ou pagar
	valorDivisao := totalGastos / float64(len(gastos))
	resultado := make(map[string]float64)

	for nome, valor := range gastos {
		resultado[nome] = valor - valorDivisao
	}

	return resultado
}

func main() {
	gastos := map[string]float64{
		"João":  100.0,
		"Maria": 50.0,
		"Pedro": 75.0,
		"Carla": 60.0,
		"André": 90.0,
	}

	resultado := calcularDivisaoConta(gastos)

	for nome, valor := range resultado {
		fmt.Printf("%s deve pagar/receber: R$ %.2f\n", nome, valor)
	}
}
