package main

import (
	"fmt"
)


func main(){
	// Exercício 1
	// ex01()
	// ex02()
	
	ex03()
}

func ex01(){
	var str string
	fmt.Print("Digite o número de palavras da frase: ")
	var n int
	fmt.Scan(&n)
	fmt.Print("Digite a frase que será invertida: ")
	for i:=0; i<n; i++{
		var char string
		fmt.Scan(&char)
		str += char + " " 
	}
	var final string
	for i := range(str){
		if i>0{
			final+=string(str[len(str)-i-1])
		}
	}
	fmt.Print(final)
}

func ex02(){
	troca := func(vector []int, change bool, x int, y int) (int, int){
		var indX, indY int
		for i := range(vector){
			if vector[i] == x {
				indX = i
			}
			if vector[i] == y {
				indY = i
			}
		}
		if change {
			vector[indX] = y
			vector[indY] = x
		}
		return indX, indY
	}
	trocaOpostosSeMenor := func(vector []int, x int, y int){
		indX, indY := troca(vector, false, x, y)
		if x < y && indX < indY && indX == len(vector)-indY-1 {
			troca(vector, true, x, y)
		}
	}
	var x, y int
	fmt.Print("Digite o valor de x: ")
	fmt.Scan(&x)
	fmt.Print("Digite o valor de y: ")
	fmt.Scan(&y)
	var n int
	fmt.Print("Digite o número de elementos do vetor: ")
	fmt.Scan(&n)
	var vector []int
	fmt.Print("Digite, cada elemento de uma vez, o vetor que será passado. ")
	for i:=0; i<n; i++{
		var n2 int
		fmt.Scan(&n2)
		vector = append(vector, n2)
	}
	trocaOpostosSeMenor(vector, x, y)
	fmt.Print(vector)
}

func ex03() {
	sort := func(vector []int) []int{
		for i:=0; i<len(vector); i++{
			for j:=0; j<len(vector)-i-1; j++{
				if vector[j] < vector[j+1] {
					vector[j], vector[j+1] = vector[j+1], vector[j]
				}
			}
		}
		return vector
	}

	var n int
	fmt.Print("Digite o número de elementos do vetor: ")
	fmt.Scan(&n)
	var vector []int
	fmt.Print("Digite, cada elemento de uma vez, o vetor que será passado. ")
	for i:=0; i<n; i++{
		var n2 int
		fmt.Scan(&n2)
		vector = append(vector, n2)
	}
	fmt.Print(sort(vector))
}
