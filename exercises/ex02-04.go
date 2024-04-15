package main

import (
	"fmt"
)

func init() {

}

//! Exercício 1

func Ex01() {
	fmt.Println("Digite três números: ")
	var n1, n2, n3 float32
	fmt.Scanln(&n1, &n2, &n3)
	numbers := []float32{n1, n2, n3}

	fmt.Printf("Em ordem crescente: %v", sort(numbers))
}

func min(list []float32) float32 {
	m := list[0]
	for i := range list {
		if list[i] < m {
			m = list[i]
		}
	}
	return m
}
func sort(list []float32) []float32 {
	var sorted, cp = make([]float32, 0, len(list)), list[:]
	for i := 0; i < len(list); i++ {
		sorted = append(sorted, min(cp))
		for j := range cp {
			if cp[j] == min(cp) && len(cp) > 1 {
				cp[j] = cp[len(cp)-1]
				cp = cp[:len(cp)-1]
				break
			}
		}
	}
	return sorted
}

//! Exercício 2

func ex02() {
	fmt.Println("Digite três números: ")
	var n1, n2, n3 float32
	fmt.Scanln(&n1, &n2, &n3)
	fmt.Printf("O maior número é %v\n", max([]float32{n1, n2, n3}))
}

func max(list []float32) float32 {
	m := list[0]
	for i := range list {
		if list[i] > m {
			m = list[i]
		}
	}
	return m
}

//! Exercício 3

func ex03() {
	var p, desc float32
	fmt.Println("O preço do produto é: ")
	fmt.Scanln(&p)
	fmt.Println("O desconto é, em %: ")
	fmt.Scanln(&desc)
	fmt.Printf("O valor final é de %v.", p*(100-desc)/100)
}
