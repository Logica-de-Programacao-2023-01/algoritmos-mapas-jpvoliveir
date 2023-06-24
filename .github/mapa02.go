package main

import "fmt"

func mergeMaps(map1, map2 map[string]int) map[string]int {
	result := make(map[string]int)

	// Copiar elementos do primeiro mapa para o resultado
	for key, value := range map1 {
		result[key] = value
	}

	// Adicionar ou atualizar elementos do segundo mapa no resultado
	for key, value := range map2 {
		result[key] = value
	}

	return result
}

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	map2 := map[string]int{
		"c": 4,
		"d": 5,
		"e": 6,
	}

	mergedMap := mergeMaps(map1, map2)
	fmt.Println(mergedMap)
}
