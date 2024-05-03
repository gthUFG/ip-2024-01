package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var QuestsList03 = []func(){
	q3_01, q3_02, q3_03, q3_04, q3_05, q3_06, q3_07, q3_08, q3_09, q3_10, q3_11, q3_12, q3_13, q3_14, q3_15, q3_16, q3_17, q3_18, q3_19, q3_20, q3_21, q3_22, q3_23, q3_24, q3_25, q3_26, q3_27, q3_28,
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
			c += count(v, float64(i))
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
	var n int
	fmt.Print("Digite o número de notas: ")
	fmt.Scan(&n)
	v := []float64{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		v = append(v, float64(x))
	}
	fmt.Printf("A nota %v aparece %v vezes.\n", v[len(v)-1], count(v, v[len(v)-1]))
	_, _, idx, max := findM(v)
	fmt.Printf("A nota %v se encontra na %vº posição.\n", max, idx)
}
func q3_11() {
	var n int
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&n)
	v := InputToList(n)
	_, min, _, max := findM(v)
	fmt.Printf("%v\n%v\n%v\n%v\n", v, reverse(v), max, min)
}
func q3_12() {
	var n int
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&n)
	v := []float64{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		v = append(v, float64(x))
	}
	v = sort(v)
	for i := range v {
		fmt.Printf("%v\n", v[i])
	}
}
func q3_13() {
	v := IndInputs()
	number, m := v[0], 0
	for i := range v {
		if count(v, v[i]) > m || (count(v, v[i]) == m && v[i] < number) {
			m = count(v, v[i])
			number = v[i]
		}
	}
	fmt.Printf("\n%v\n%v\n", number, m)
}
func q3_14() {
	v := sort(IndInputs())
	fmt.Printf("\nMediana: %.2f\n", []float64{v[len(v)/2], (v[len(v)/2] + v[len(v)/2-1]) / 2}[BToi(len(v)%2 == 0)])
}
func q3_15() {
	var n int
	fmt.Print("Digite o número de testes: ")
	fmt.Scan(&n)
	for k := 0; k < n; k++ {
		var n2 int
		fmt.Print("Digite o número de elementos: ")
		fmt.Scan(&n2)
		v := InputToList(n2)
		c, minDiff := 0, float64(1000)
		for i := range v {
			for j := i + 1; j < len(v); j++ {
				c += 1
				if minDiff > module(v[i]-v[j]) {
					minDiff = module(v[i] - v[j])
				}
			}
		}
		fmt.Printf("A menor diferença é de %v. Foram realizados %v testes.\n\n", minDiff, c)
	}
}
func q3_16() {
	v1 := InputToList(2)
	v2 := InputToList(int(v1[0]))
	c, v3 := 0, []int{}
	for i := range v2 {
		if v2[i] <= 0 {
			c += 1
			v3 = append(v3, i+1)
		}
	}
	if c >= int(v1[1]) {
		fmt.Println("Não será cancelada.")
		for i := range v3 {
			fmt.Println(v3[i])
		}
	} else {
		fmt.Println("Será cancelada.")
	}
}
func q3_17() {
	var n int
	fmt.Print("Digite o número de elementos: ")
	fmt.Scan(&n)
	v, c := InputToList(n), 0
	for i := range v {
		if count(v, v[i]) == 1 {
			c += 1
		}
	}
	fmt.Printf("Há %v números únicos.", c)
}
func q3_18() {
	var n int
	fmt.Print("Digite o número de CPFs a serem analisados: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		v := InputToList(11)
		sum1, sum2 := 0, 0
		for i := range v[:9] {
			sum1 += int(v[i]) * (i + 1)
			sum2 += int(v[i]) * (9 - i)
		}
		if sum1%11 == 10 {
			sum1 = 11
		}
		if sum2%11 == 10 {
			sum2 = 11
		}
		if sum1%11 == int(v[9]) && sum2%11 == int(v[10]) {
			fmt.Println("CPF válido.")
		} else {
			fmt.Println("CPF inválido.")
		}
	}
}
func q3_19() {
	v := IndInputs()
	v2 := []float64{}
	for i := range v {
		alreadyAdded := false
		for j := range v2 {
			if v2[j] == v[i] {
				alreadyAdded = true
			}
		}
		if !alreadyAdded {
			v2 = append(v2, v[i])
		}
	}
	for i := range v2 {
		fmt.Println(v2[i])
	}
}
func q3_20() {
	var n int
	fmt.Print("Digite o número de pontos: ")
	fmt.Scan(&n)
	v := [][]float64{}
	for i := 0; i < n; i++ {
		p := InputToList(3)
		v = append(v, p)
	}
	for i := 0; i < len(v)-1; i++ {
		var m float64
		for j := range v[i] {
			if v[i][j]-v[i+1][j] > m {
				m = v[i][j] - v[i+1][j]
			}
		}
		fmt.Printf("Entre os pontos %v e %v: %.2f.\n", i+1, i+2, m)
	}
}
func q3_21() {
	v := sort(IndInputs())
	v1, v2 := []float64{}, []float64{}
	for i := range v {
		if int(v[i])%2 == 0 {
			v1 = append(v1, v[i])
		} else {
			v2 = append(v2, v[i])
		}
	}
	v2 = reverse(v2)
	fmt.Println()
	for i := range v1 {
		fmt.Println(v1[i])
	}
	for i := range v2 {
		fmt.Println(v2[i])
	}
}
func q3_22() {
	for {
		fmt.Print("Digite a quantidade de caracteres e o limite de tamanho: ")
		info := InputToList(2)
		if int(info[0]) <= 0 {
			break
		}
		fmt.Print("Digite o número: ")
		var n string
		fmt.Scan(&n)
		try := ""
		for i := int(info[0] - 1); i >= 0; i-- {
			if len(try) < int(info[1]) {
				try = string(n[i]) + try
			} else {
				if try[0] < n[i] {
					try = string(n[i]) + try[1:]
				}
			}
		}
		fmt.Println(try)
	}
}
func q3_23() {
	fmt.Print("Digite as frases, separadas por \";\": ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	vows := make([]int, 11)
	separated := false
	for i := 0; i < len(str); i++ {
		vowels := [][]string{{"a", "A"}, {"e", "E"}, {"i", "I"}, {"o", "O"}, {"u", "U"}, {";", ";"}}
		for j := 0; j < len(vowels); j++ {
			if string(str[i]) == vowels[j][0] || string(str[i]) == vowels[j][1] {
				if j == 5 {
					separated = true
					vows[10] += 1
				} else {
					vows[j+[]int{0, 5}[BToi(separated)]] += 1
				}
			}
		}
	}
	if vows[10] != 1 {
		fmt.Print(vows[10], "Formato inválido.\n")
		return
	}
	var distance float64
	for i := 0; i < 5; i++ {
		distance += math.Pow(float64(vows[i]-vows[i+5]), 2)
	}
	distance = math.Sqrt(distance)
	fmt.Printf("%v\n%v\nA distância entre as strings é de %.2f.\n", vows[:5], vows[5:10], distance)
}
func q3_24() {
	for {
		var n int
		fmt.Print("Digite o número de elementos: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		v := InputToList(n)
		fmt.Print(sort(v), '\n')
	}
}
func q3_25() {
	fmt.Print("Quais foram os números sorteados? ")
	ns, n := sort(InputToList(6)), 0
	fmt.Print("Houveram quantos apostadores? ")
	fmt.Scan(&n)
	counter := []int{0, 0, 0}
	for i := 0; i < n; i++ {
		c := 0
		v := sort(InputToList(6))
		for i := range v {
			for j := range ns {
				if v[i] == ns[j] {
					c += 1
				}
			}
		}
		if c == 6 {
			counter[0] += 1
		} else if c == 5 {
			counter[1] += 1
		} else if c == 4 {
			counter[2] += 1
		}
	}
	for i := range counter {
		fmt.Printf("Houve %v acertador(es) da %v.\n", counter[i], []string{"sena", "quina", "quadra"}[i])
	}
}
func q3_26() {
	var n int
	fmt.Print("Digite o número de casos de teste: ")
	fmt.Scan(&n)
	for k := 0; k < n; k++ {
		fmt.Print("Digite o número de cada anão: \n")
		v := sort(IndInputs(9))
	out:
		for i := range v {
			for j := range v {
				if i != j && Sum(v)-v[i]-v[j] == 100 {
					fmt.Print("Os anões são os de número:\n")
					for anao := range v {
						if anao != i && anao != j {
							fmt.Println(v[anao])
						}
					}
					fmt.Println("\n")
					break out
				}
			}
		}
	}
}
func q3_27() {
	var n, n2 int
	fmt.Print("Digite quantos elementos possui o primeiro vetor: ")
	fmt.Scan(&n)
	fmt.Print("Digite quantos elementos possui o segundo vetor: ")
	fmt.Scan(&n2)
	v1 := IndInputs(n)
	v2 := IndInputs(n2)
	for i := range v2 {
		v1 = append(v1, v2[i])
	}
	v1 = sort(v1)
	for i := range v1 {
		fmt.Print("\n", v1[i])
	}
}
func q3_28() {
	var n1, n2 int
	fmt.Print("Digite quantos elementos possui o primeiro vetor: ")
	fmt.Scan(&n1)
	fmt.Print("Digite quantos elementos possui o segundo vetor: ")
	fmt.Scan(&n2)
	v1 := InputToList(n1)
	v2 := InputToList(n2)
	v := append(v1[:], v2[:]...)
	var u, n []float64
	isInside := func(arr []float64, elem float64) bool {
		for i := range arr {
			if arr[i] == elem {
				return true
			}
		}
		return false
	}
	for i := range v {
		if !isInside(u, v[i]) {
			if isInside(v1, v[i]) || isInside(v2, v[i]) {
				u = append(u, v[i])
			}
		}
		if !isInside(n, v[i]) {
			if isInside(v1, v[i]) && isInside(v2, v[i]) {
				n = append(n, v[i])
			}
		}
	}
	fmt.Println(sort(u))
	fmt.Print(sort(n))
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

func count(list []float64, e float64) int {
	c := 0
	for i := range list {
		if list[i] == e {
			c++
		}
	}
	return c
}

func copy(list []float64) []float64 {
	nl := []float64{}
	for i := range list {
		nl = append(nl, list[i])
	}
	return nl
}

func sort(list []float64) []float64 {
	var sorted, cp = make([]float64, 0, len(list)), list[:]
	for i := 0; i < len(list); i++ {
		_, min, _, _ := findM(cp)
		sorted = append(sorted, min)
		for j := range cp {
			if cp[j] == min && len(cp) > 1 {
				cp[j] = cp[len(cp)-1]
				cp = cp[:len(cp)-1]
				break
			}
		}
	}
	return sorted
}

func IndInputs(params ...int) []float64 {
	var n int
	if len(params) == 0 {
		fmt.Printf("Digite o número de %v: ", params[0])
		fmt.Scan(&n)
	} else {
		n = params[0]
	}

	v := []float64{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		v = append(v, float64(x))
	}
	return v
}

func BToi(stat bool) int {
	if stat {
		return 1
	} else {
		return 0
	}
}

func NumMap(arr []float64, f func(x float64) float64) []float64 {
	nl := []float64{}
	for i := range arr {
		nl = append(nl, f(arr[i]))
	}
	return nl
}
