Para rodar os exercícios, vá até a pasta exercises/ e rode o comando "go run ."
Há um arquivo para cada lista, e um arquivo seletor que permite selecionar a lista e a questão que será avaliada.

package main

func main() {
	// fmt.Print("\n\033[0;36mQual exercício? \033[m")
	// var n int
	// fmt.Scanln(&n)
	// if n <= len(QuestsList01) && n >= 1 {
	// 	fmt.Printf("\n\033[1;36m------------- Questão %v -------------\033[1;37m\n\n", n)
	// 	QuestsList01[n-1]()
	// 	fmt.Printf("\n\n\033[1;36m-------------------------------------\033[1;37m\n\n", n)
	// } else {
	// 	fmt.Printf("Questão inválida. Há %v questões.\n", len(QuestsList01))
	// }
	q2_04()
}

package main

import (
	"fmt"
	// "strings"
	"math"
)

var QuestsList01 = []func(){q01, q02, q03, q04, q05, q06, q07, q08, q09, q10, q11, q12, q13, q14, q15, q16, q17, q18, q19, q20}

func init() {
	// q2_01()

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
	var ops = [][]string{[]string{"soma", "+"}, []string{"subtração", "-"}, []string{"multiplicação", "x"}, []string{"divisão", "/"}}
	fmt.Print("Dê os valores de n, i, K e s: ")
	infs := InputToList(4)
	for i := 0; i < 4; i++ {
		fmt.Printf("Tabuada de %v:\n", ops[i][0])
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
