package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var QuestsList01 = []func(){q1_01, q1_02, q1_03, q1_04, q1_05, q1_06, q1_07, q1_08, q1_09, q1_10,
	q1_11, q1_12, q1_13, q1_14, q1_15, q1_16, q1_17, q1_18, q1_19, q1_20}

func init() {
}

func q1_01() {
	fmt.Print("Notas: ")
	// Jeito massa fácil gostoso (só aguenta três notas)
	// var n1, n2, n3 float64
	// fmt.Scanln(&n1, &n2, &n3)
	// notes := []float64{n1,n2,n3}

	// Jeito masoquista louco pirado (aguenta infinitas)
	notes := InputToList()

	var s float64
	for i := range notes {
		s += notes[i]
	}
	average := s / float64(len(notes))
	if average >= 6 {
		fmt.Printf("Aprovado, sua nota foi de %.2f.", average)
	} else {
		fmt.Printf("Reprovado, sua nota foi de %.2f.", average)
	}
}
func q1_02() {

	fmt.Println("Número de jogos: ")
	var n int
	fmt.Scan(&n)
	for i := 1; i < (n + 1); i++ {
		fmt.Printf("Jogo %v: ", i)
		data := InputToList()
		fmt.Printf("A renda do jogo foi de R$%.2f\n", (data[0] / 100 * (data[1]*1 + data[2]*5 + data[3]*10 + data[4]*20)))
	}
}
func q1_03() {

	fmt.Println("Digite três algarismos: ")
	var n1, n2, n3 string
	fmt.Scanln(&n1, &n2, &n3)
	if len(n1) != 1 || len(n2) != 1 || len(n3) != 1 {
		fmt.Println("Dígito inválido.")
		return
	}
	var n, _ = strconv.ParseFloat(n1+n2+n3, 32)
	fmt.Printf("O número é %v, seu quadrado é %v", n, math.Pow(n, 2))
}
func q1_04() {
	var s, q float64
	fmt.Println("Salário Mínimo: ")
	fmt.Scan(&s)
	fmt.Println("Consumo em kW: ")
	fmt.Scan(&q)
	fmt.Printf("Custo por kW: R$ %.2f\n", s*0.007)
	fmt.Printf("Custo do consumo: R$ %.2f\n", s*0.007*q)
	fmt.Printf("Custo com desconto: R$ %.2f", s*0.007*q*0.9)
}
func q1_05() {
	var acc, waste float64
	var kind string
	var value float64
	fmt.Scanln(&acc, &waste, &kind)
	if kind == "R" {
		value = 5 + 0.05*waste
	} else if kind == "C" {
		value = 500 + 0.25*(waste-80)
	} else if kind == "I" {
		value = 800 + 0.04*(waste-100)
	} else {
		fmt.Println("Tipo de consumidor inválido.")
		return
	}
	fmt.Printf("Conta: %v\n", acc)
	fmt.Printf("Valor da conta: R$%.2f", value)
}
func q1_06() {
	var n int
	fmt.Print("Número de temperaturas em ºF: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Temperatura %v: ", i+1)
		var t float64
		fmt.Scan(&t)
		fmt.Printf("%.2fºF equivalem a %.2fºC.\n", t, 5*(t-32)/9)
	}
}
func q1_07() {
	var t, r float64
	fmt.Print("Temperatura em ºF: ")
	fmt.Scan(&t)
	fmt.Print("Quantidade de chuva em polegadas: ")
	fmt.Scan(&r)
	fmt.Printf("O valor em ºC: %.2f\nA quantidade de chuva em mm: %.2f", 5*(t-32)/9, r*25.4)
}
func q1_08() {
	var r, h float64
	fmt.Print("Raio da lata: ")
	fmt.Scan(&r)
	fmt.Print("Altura da lata: ")
	fmt.Scan(&h)
	fmt.Printf("O valor do custo de fabricação é: R$%.2f.", (2*math.Pi*r*h+2*math.Pi*math.Pow(r, 2))*100)
}
func q1_09() {
	fmt.Print("Valores de A, B e C: ")
	coefs := InputToList(3)
	fmt.Printf("∆ = %v.", math.Pow(coefs[1], 2)-4*coefs[0]*coefs[2])
}
func q1_10() {
	fmt.Print("Escreva os quatro elementos da matriz (M11, M12, M21, M22): ")
	els := InputToList(4)
	fmt.Printf("O valor do determinante é de %v.", els[0]*els[3]-els[1]*els[2])
}
func q1_11() {
	fmt.Print("Digite um número: ")
	var n int
	fmt.Scan(&n)
	if n%3 == 0 && n%5 == 0 {
		fmt.Print("O número é divisível por 3 e por 5.")
	} else {
		fmt.Print("O nḿero não é divisível.")
	}
}
func q1_12() {
	var t float64
	fmt.Print("Por quanto tempo será alugado? ")
	fmt.Scan(&t)
	var price float64
	if t < 3 {
		price = t * 5
	} else if t >= 3 {
		price = t / 3 * 10
	} else {
		fmt.Print("Número de horas inválido.")
		return
	}
	fmt.Printf("O preço será de R$%.2f.", price)
}
func q1_13() {
	var x float32
	fmt.Print("Digite a nota do aluno: ")
	fmt.Scan(&x)
	var c string
	if x > 10 || x < 0 {
		fmt.Print("Nota inválida.")
		return
	}
	if x >= 9 {
		c = "A"
	} else if x >= 7.5 {
		c = "B"
	} else if x >= 6 {
		c = "C"
	} else {
		c = "D"
	}
	fmt.Printf("Nota: %v. Conceito: %v.", x, c)
}
func q1_14() {
	var h, a float64
	fmt.Print("Digite a altura da pirâmide em m³: ")
	fmt.Scan(&h)
	fmt.Print("Digite a aresta da base da pirâmide em m³: ")
	fmt.Scan(&a)
	fmt.Printf("O volume da pirâmide é de %.2f metros cúbicos.", (6*math.Pow(a, 2)*math.Sqrt(3)/4)*h/3)
}
func q1_15() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	for i := 2; i <= n; i += 2 {
		fmt.Printf("%v^2 = %v\n", i, i*i)
	}
}
func q1_16() {
	var s, f float64
	fmt.Print("Qual o salário a ser reajustado? ")
	fmt.Scan(&s)
	if s <= 300 {
		f = s * 1.5
	} else if s > 300 {
		f = s * 1.3
	} else {
		fmt.Print("Salário inválido.")
		return
	}
	fmt.Printf("Salário com reajuste: R$%.2f.", f)
}
func q1_17() {
	fmt.Print("Digite os valores de x e de y: ")
	var x, y int
	fmt.Scan(&x, &y)
	if x%2 == 0 {
		for i := 0; i < y; i++ {
			fmt.Print(x+i*2, " ")
		}
	} else {
		fmt.Print("O primeiro número não é par.")
	}
}
func q1_18() {
	var a1, r, n, s int
	fmt.Print("Digite os valores de a1, r e n: ")
	fmt.Scan(&a1, &r, &n)
	s = a1
	for i := 1; i < n; i++ {
		s += a1 + i*r
	}
	fmt.Printf("A soma dessa P.A. é de %v.", s)
}
func q1_19() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Print("Número inválido.")
		return
	}
	var s float64
	for i := 1; i <= n; i++ {
		s += 1 / float64(i)
	}
	fmt.Printf("%.6f", s)
}
func q1_20() {
	fmt.Print("Digite o tempo em horas, minutos e segundos: ")
	ts := InputToList(3)
	fmt.Printf("O tempo é de %v segundos.", ts[0]*3600+ts[1]*60+ts[2])
}

func InputToList(params ...int) []float64 {
	reader := bufio.NewReader(os.Stdin)
	inp, _ := reader.ReadString('\n')
	notes := toFloat64Elements(strings.Split(strings.TrimSpace(inp), " "))
	if len(params) != 0 {
		if len(notes) != params[0] {
			fmt.Println("\n\033[1;31mO input está incorreto.\n\033[m")
			return []float64{}
		}
	}

	return notes
}
func toFloat64Elements(l1 []string) []float64 {
	var floats = []float64{}
	for i := range l1 {
		l1[i] = strings.TrimSuffix(l1[i], "\n")
		float, _ := strconv.ParseFloat(l1[i], 32)
		floats = append(floats, float)
	}
	return floats
}

// func toFixed(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }
