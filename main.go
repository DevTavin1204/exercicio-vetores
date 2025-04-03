package main

import "fmt"

func main() {
    // Slice
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros) 
	numeros = append(numeros, 6, 7, 8)
	fmt.Println(numeros, len(numeros), cap(numeros))
	// Slice
	nomes := []string{"otavio", "denise", "fabricio", "brenda", "lucas"}
	fmt.Println(nomes)
	rangeOne := nomes[0:2]
	fmt.Println(rangeOne)
	rangeTwo := nomes[3:5]
	fmt.Println(rangeTwo)
	rangeThree := nomes[2]
	fmt.Println(rangeThree)
}