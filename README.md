Para rodar os exercícios, vá até a pasta exercises/ e rode o comando "go run ."
Há um arquivo para cada lista, e um arquivo seletor que permite selecionar a lista e a questão que será avaliada.

package main

import (
	"fmt"
	"math"
)

var QuestsList03 = []func(){q3_01, q3_02, q3_03, q3_04, q3_05, q3_06, q3_07, q3_08, q3_09,
	q3_10, q2_11, q2_12, q2_13, q2_14, q2_15, q2_16, q2_17, q2_18,
	q2_19, q2_20, q2_21, q2_22, q2_23, q2_24, q2_25, q2_26, q2_27,
}

func init() {
}

func q3_01() {
	var n, m int
	fmt.Print("Digite o número de elementos do vetor: ")
	fmt.Scan(&n)
	vector := InputToList(n)
	fmt.Print("Quantos números quer verificar? ")
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scan(&x)
		for j := 0; j < len(vector); j++ {
			if vector[j] == float64(x) {
				fmt.Print("Encontrei o número.\n")
				break
			}
			if j == len(vector)-1 {
				fmt.Print("Não encontrei o número.\n")
			}
		}

	}
}
func q3_02() {
	var n, k, c int
	for n < 1 || n > 1000 {
		fmt.Print("Digite o valor de N: ")
		fmt.Scan(&n)
	}
	fmt.Print("Digite os valores da lista: ")
	v := InputToList(n)
	fmt.Print("Digite o valor de K: ")
	fmt.Scan(&k)
	for i := range v {
		if v[i] >= float64(k) {
			c += 1
		}
	}
	fmt.Printf("%v números do vetor são maiores do que K.", c)

}
func q3_03() {
	var n int
	fmt.Print("Digite o número de elementos da lista: ")
	fmt.Scan(&n)
	fmt.Print("Digite os elementos da lista: ")
	v := InputToList(n)
	fmt.Print(reverse(v))
}
func q3_04() {
	var n int
	fmt.Print("Digite o número de elementos da lista: ")
	fmt.Scan(&n)
	v := InputToList(n)
	for i := range v {
		fmt.Printf("%v ", v[i])
	}
}
func q3_05() {
	for {
		var n int
		fmt.Print("Digite o número de elementos da lista: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		v := InputToList(n)
		_, _, idx, m := findM(v)
		fmt.Printf("%v é o maior número, encontrado na posição %v.\n", m, idx)
	}
}
func q3_06() {
	var n int
	fmt.Print("Digite o tamanho da lista: ")
	fmt.Scan(&n)
	v := make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}
	fmt.Printf("O somatório da lista é de %v.\n", Sum(v))
}
func q3_07() {
	for {
		fmt.Print("Digite o número de elementos da lista: ")
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		v := InputToList(n)
		c := 0
		_, _, _, m := findM(v)
		for i := 0; i <= int(m); i++ {
			for j := range v {
				if int(v[j]) == i {
					c += 1
				}
			}
			fmt.Printf("(%v) %v\n", i, c)
		}
	}
}
func q3_08() {
	for {
		var n int
		fmt.Print("Digite o número em decimal a ser transformado para binário: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		bin := []float64{}
		for n > 1 {
			bin = append(bin, float64(n%2))
			n /= 2
		}
		bin = append(bin, float64(n))
		bin = reverse(bin)
		for i := range bin {
			fmt.Print(bin[i])
		}
		fmt.Println()
	}
}
func q3_09() {
	var n int
	for n < 2 || n > 1000 {
		fmt.Print("Digite o valor de N: ")
		fmt.Scan(&n)
	}
	v := [][]float64{}
	for i := 0; i < n; i++ {
		inp := InputToList(3)
		v = append(v, inp)
	}
	d := func(a []float64, b []float64) float64 {
		return math.Sqrt(math.Pow(a[0]-b[0], 2) + math.Pow(a[1]-b[1], 2) + math.Pow(a[2]-b[2], 2))
	}
	for i := 0; i < len(v)-1; i++ {
		fmt.Printf("A distância entre o ponto %v e o ponto %v é de %.2f.\n", i+1, i+2, d(v[i], v[i+1]))
	}
}
func q3_10() {

}

func findM(list []float64) (int, float64, int, float64) {
	var min, max float64 = list[0], list[0]
	var idx1, idx2 int = 0, 0
	for i := range list {
		if list[i] < min {
			min = list[i]
			idx1 = i
		}
		if list[i] > max {
			max = list[i]
			idx2 = i
		}
	}
	return idx1, min, idx2, max
}

func reverse(list []float64) []float64 {
	nv := make([]float64, len(list))
	for i := range list {
		nv[i] = float64(list[len(list)-i-1])
	}
	return nv
}
