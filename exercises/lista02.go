package main

import (
	"fmt"
	"math"
)

var QuestsList02 = []func(){q2_01, q2_02, q2_03, q2_04, q2_05, q2_06, q2_07, q2_08, q2_11}

func init() {

}

func q2_01() {
	keep := true
	sitList := []string{"Aprovado", "Reprovado por nota", "Reprovado por frequência", "Reprovado por nota e frequência"}
	for keep {
		data := InputToList(16)
		if data[0] == float64(-1) {
			keep = false
			break
		}
		var situation int
		mat := int(data[0])
		grade := 0.7*(sum(data[1:9])/8) + 0.15*(sum(data[9:14])/5) + 0.15*(data[14])
		if grade < 6 {
			situation += 1
		}
		if data[15]/128 < float64(3)/4 {
			situation += 2
		}
		defer fmt.Printf("Matrícula: %v, Nota final: %.2f, Situação Final: %v.\n", mat, grade, sitList[situation])
	}
}
func q2_02() {
	var a, b float64
	fmt.Print("População no país A: ")
	fmt.Scan(&a)
	fmt.Print("População no país B: ")
	fmt.Scan(&b)
	var n int
	for a*math.Pow(1.03, float64(n)) < b*math.Pow(1.015, float64(n)) {
		n += 1
	}
	fmt.Printf("São necessário %v anos.", n)
}
func q2_03() {
	var n int
	fmt.Scan(&n)
	fmt.Printf("%v! = %v", n, fact(n))
}
func q2_04() {
	var ops = [][]string{{"soma", "+"}, {"subtração", "-"}, {"multiplicação", "x"}, {"divisão", "/"}}
	fmt.Print("Dê os valores de n, i, K e s: ")
	infs := InputToList(4)
	for i := 0; i < 4; i++ {
		fmt.Printf("\nTabuada de %v:\n", ops[i][0])
		for j := 0; j < int(infs[2]); j++ {
			var result float64
			if i == 0 {
				result = infs[0] + float64(j)*infs[3] + infs[1]
			} else if i == 1 {
				result = infs[0] - float64(j)*infs[3] + infs[1]
			} else if i == 2 {
				result = infs[0] * (float64(j)*infs[3] + infs[1])
			} else {
				result = infs[0] / (float64(j)*infs[3] + infs[1])
			}
			fmt.Printf("%.2f %v %.2f = %.2f\n", infs[0], ops[i][1], float64(j)*infs[3]+infs[1], result)
		}
	}
}
func q2_05() {
	var n int
	fmt.Print("Digite o tamanho da sequência: ")
	fmt.Scan(&n)
	fmt.Print("Digite a sequência: ")
	seq := InputToList(n)
	m := 0
	for i := range seq {
		c := 0
		for j := i; j < len(seq)-1; j++ {
			if seq[j] < seq[j+1] {
				c += 1
				if c > m {
					m = c
				}
			} else {
				break
			}
		}
	}
	fmt.Printf("\nO comprimento do segmento crescente máximo é de %v.", m)
}
func q2_06() {
	var n int
	fmt.Print("Digite o tamanho da sequência: ")
	fmt.Scan(&n)
	fmt.Print("Digite a sequência: ")
	seq := InputToList(n)
	m := 0
	for i := range seq {
		c := 1
		for j := i; j < len(seq)-1; j++ {
			if seq[j] == seq[j+1] {
				c += 1
				if c > m {
					m = c
				}
			} else {
				break
			}
		}
	}
	fmt.Printf("\nA maior subsequência de elementos iguais possui %v elementos.", m)
}
func q2_07() {
	fmt.Print("Digite a sequência com final 0: ")
	seq := InputToList()
	var sp, np, si, ni float64
	for i := range seq {
		if seq[i] == 0 {
			seq = seq[0:i]
		}
	}
	for i := range seq {
		if int(seq[i])%2 == 0 {
			sp += seq[i]
			np += 1
		} else {
			si += seq[i]
			ni += 1
		}
	}
	fmt.Printf("Média par: %.2f.\nMédia impar: %.2f.", sp/np, si/ni)
}
func q2_08() {
	fmt.Print("Digite o número de times: ")
	var n int
	fmt.Scan(&n)
	var finals = [][]int{}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var alreadyCounted = false
			if i != j {
				for k := range finals {
					if (finals[k][0] == i && finals[k][1] == j) || (finals[k][0] == j && finals[k][1] == i) {
						alreadyCounted = true
						break
					}
				}
			} else {
				alreadyCounted = true
			}
			if !alreadyCounted {
				finals = append(finals, []int{i, j})
				fmt.Printf("Final %v: Time%v X Time%v\n", len(finals), i, j)
			}
		}
	}

}
func q2_11() {
	fmt.Print("Digite o número cuja raiz será calculada: ")
	var rk float64 = 1
	var n, e float64
	fmt.Scan(&n)
	fmt.Print("Digite o número cuja raiz será calculada: ")
	fmt.Scan(&e)
	diff := n - math.Pow(rk, 2)
	for diff > e {
		rk = (rk + n/rk) / 2
		diff = module(n - math.Pow(rk, 2))
		defer fmt.Printf("r: %.9f, erro: %.9f\n", rk, diff)
	}
}

func module(n float64) float64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func fact(n int) int {
	if n == 1 {
		return n
	} else {
		return fact(n-1) * n
	}
}

func sum(list []float64) float64 {
	var s float64
	for i := range list {
		s += list[i]
	}
	return s
}
