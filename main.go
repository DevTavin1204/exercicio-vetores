package main

import "fmt"

func main() {
    // Slice
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros) 
	numeros = append(numeros, 6, 7, 8)
	fmt.Println(numeros, len(numeros), cap(numeros))
}