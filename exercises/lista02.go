package main

import (
	"fmt"
	"math"

	// "strings"
	"strconv"
)

var QuestsList02 = []func(){q2_01, q2_02, q2_03, q2_04, q2_05, q2_06, q2_07, q2_08, q2_09,
	q2_10, q2_11, q2_12, q2_13, q2_14, q2_15, q2_16, q2_17, q2_18,
	q2_19, q2_20, q2_21, q2_22, q2_23, q2_24, q2_25, q2_26, q2_27}

func init() {

}

func q2_01() {
	sitList := []string{"Aprovado", "Reprovado por nota", "Reprovado por frequência", "Reprovado por nota e frequência"}
	for {
		data := InputToList(16)
		if data[0] == float64(-1) {
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
			results := []float64{infs[0] + float64(j)*infs[3] + infs[1], infs[0] - float64(j)*infs[3] + infs[1],
				infs[0] * (float64(j)*infs[3] + infs[1]), infs[0] / (float64(j)*infs[3] + infs[1])}
			fmt.Printf("%.2f %v %.2f = %.2f\n", infs[0], ops[i][1], float64(j)*infs[3]+infs[1], results[i])
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
	var n, c int
	fmt.Scan(&n)
	// var finals = [][]int{}
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			if i != j {
				c += 1
				fmt.Printf("Final %v: Time%v X Time%v\n", c, i, j)
			}
			// var alreadyCounted = false
			// if i != j {
			// 	for k := range finals {
			// 		if (finals[k][0] == i && finals[k][1] == j) || (finals[k][0] == j && finals[k][1] == i) {
			// 			alreadyCounted = true
			// 			break
			// 		}
			// 	}
			// } else {
			// 	alreadyCounted = true
			// }
			// if !alreadyCounted {
			// 	finals = append(finals, []int{i, j})
			// 	fmt.Printf("Final %v: Time%v X Time%v\n", len(finals), i, j)
			// }
		}
	}
}
func q2_09() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	isPrime := 1
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			isPrime = 0
		}
	}
	fmt.Printf("%v %vé primo.", n, []string{"não ", ""}[isPrime])
}
func q2_10() {
	fmt.Print("Digite as matrículas, o valor da hora e o tempo dos trabalhadores: ")
	for {
		info := InputToList(3)
		if info[0] == 0 {
			break
		}
		defer fmt.Printf("Matrícula: %v, Salário: %.2f\n", int(info[0]), info[1]*info[2])
	}
}
func q2_11() {
	fmt.Print("Digite o número cuja raiz será calculada: ")
	var rk float64 = 1
	var n, e float64
	fmt.Scan(&n)
	fmt.Print("Digite o erro máximo a ser tolerado: ")
	fmt.Scan(&e)
	diff := n - math.Pow(rk, 2)
	for diff > e {
		rk = (rk + n/rk) / 2
		diff = module(n - math.Pow(rk, 2))
		defer fmt.Printf("r: %.9f, erro: %.9f\n", rk, diff)
	}
}
func q2_12() {
	var n, n0, nf float64
	max := []float64{0, 0, 0}
	fmt.Print("Digite o valor normal do ingresso: ")
	fmt.Scan(&n)
	fmt.Print("Digite o valor mínimo do ingresso: ")
	fmt.Scan(&n0)
	fmt.Print("Digite o valor máximo do ingresso: ")
	fmt.Scan(&nf)
	if n0 >= nf {
		fmt.Print("Intervalo inválido.")
		return
	}
	for i := n0; i <= nf; i++ {
		var q, l float64
		if i < n {
			q = 120 + (n-i)/0.5*25
		} else {
			q = 120 - (i-n)/0.5*30
		}
		l = q*i - (200 + 0.05*q)
		fmt.Printf("V: %.2f, N: %v, L: %.2f\n", i, q, l)
		if l > max[1] {
			max[0] = i
			max[1] = l
			max[2] = q
		}
	}
	for i := range max {
		fmt.Printf("%v: %v\n", []string{"Melhor valor final", "Lucro", "Número de ingressos"}[i], max[i])
	}
}
func q2_13() {
	var n int
	fmt.Print("Digite o número de grãos de milho em cada casa: ")
	fmt.Scan(&n)
	fmt.Printf("Pode-se colocar %v grãos de milho.", 32*3*n-n)
}
func q2_14() {
	fmt.Print("Digite as dimensões m * n da matriz: ")
	ds := InputToList(2)
	for i := 1; i <= int(ds[0]); i++ {
		for j := 1; j <= int(ds[1]); j++ {
			if i > j {
				if j > 1 {
					fmt.Print(" - ")
				}
				fmt.Printf("(%v,%v)", i, j)
			}
		}
		fmt.Println()
	}
}
func q2_15() {
	reverse := func(n int) int {
		cp := strconv.Itoa(n)
		var final string
		for i := range cp {
			final += string(cp[len(cp)-(i+1)])
		}
		f2, _ := strconv.Atoi(final)
		return f2
	}
	var n int
	fmt.Print("Digite o número de testes: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var a, b, m int
		fmt.Printf("Teste %v: ", i+1)
		fmt.Scanln(&a, &b)
		if (a < 100 || a > 999) || (b < 100 || b > 999) {
			fmt.Print("Digitos inválidos. Devem conter três dígitos.")
			return
		}
		if reverse(a) > reverse(b) {
			m = reverse(a)
		} else {
			m = reverse(b)
		}
		fmt.Printf("O maior é %v.\n", m)
	}
}
func q2_16() {
	var n int
	fmt.Print("Digite a hipotenusa limite: ")
	fmt.Scan(&n)
	for k := 1; k <= n; k++ {
		for i := 1; i < k; i++ {
			for j := k - 1; j > i; j-- {
				if i*i+j*j == k*k {
					fmt.Printf("Hipotenusa = %v, catetos %v e %v.\n", k, i, j)
				}
			}
		}
	}
}
func q2_17() {
	fmt.Println("Coloque os dados de cada produto: ")
	data := []float64{0, 0, 0, 0, 0, 0, 0}
	for {
		prodInfo := InputToList()
		if len(prodInfo) != 4 {
			break
		}
		code := prodInfo[0]
		cPrice := prodInfo[1]
		vPrice := prodInfo[2]
		q := prodInfo[3]
		profit := vPrice / cPrice
		if profit > 1.2 {
			data[2] += 1
		} else if profit >= 1.1 && profit <= 1.2 {
			data[1] += 1
		} else {
			data[0] += 1
		}
		if profit > data[4] {
			data[3] = code
		}
		if q > data[5] {
			data[4] = code
		}
		data[5] += cPrice * q
		data[6] += vPrice * q
	}
	displayPhrases := []string{
		"Quantidade de mercadoias que geraram lucro menor que 10%%: %v",
		"Quantidade de mercadorias que geraram lucro maior ou igual a 10%%, e menor ou igual a 20%%: %v",
		"Quantidade de mercadorias que geraram lucro maior do que 20%%: %v",
		"Código da mercadoria que gerou maior lucro: %v",
		"Código da mercadoria mais vendida: %v",
		"Valor total de compras: %.2f, valor total de vendas: %.2f, percentual de lucro total: %.2f%%.",
	}
	for i := range displayPhrases {
		if i < len(displayPhrases)-1 {
			fmt.Printf(displayPhrases[i]+"\n", int(data[i]))
		} else {
			fmt.Printf(displayPhrases[i]+"\n\n", data[5], data[6], 100*(data[6]/data[5]-1))
		}
	}
}
func q2_18() {
	fmt.Print("Digite três números para calcular seu MMC: ")
	ns := InputToList(3)
	mmc := 1
	i := 2
	for sum(ns) > 3 {
		divided := false
		cp := []float64{ns[0], ns[1], ns[2]}
		for j := range ns {
			if int(ns[j])%i == 0 {
				ns[j] /= float64(i)
				divided = true
			}
		}
		if divided {
			mmc *= i
			fmt.Printf("%v :%v\n", cp, i)
		} else {
			i++
		}
	}
	fmt.Printf("MMC: %v.", mmc)
}
func q2_19() {
	fmt.Print("Número de cubos analisados: ")
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print(i, "*", i, "*", i, " = ")
		for c := i*i - (i - 1); c <= i*i+(i-1); c += 2 {
			if c > i*i-(i-1) {
				fmt.Print("+")
			}
			fmt.Print(c)
		}
		fmt.Println()
	}
}
func q2_20() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	fmt.Print(n, ": ")
	d := divs(n)
	for i := range d {
		if i != 0 {
			fmt.Print(" + ")
		}
		fmt.Print(d[i])
	}

	fmt.Print(" = ", sum(d), ". ")
	if int(sum(d)) == n {
		fmt.Print("O número é perfeito.")
	} else {
		fmt.Print("O número não é perfeito.")
	}
}
func q2_21() {
	for {
		var n int
		fmt.Print("Número de elementos: ")
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		fmt.Println("Digite os elementos: ")
		l := InputToList(n)
		for i := range l {
			if i > 0 && l[i-1] > l[i] {
				fmt.Print("Desordenada.\n")
				break
			}
			if i == len(l)-1 {
				fmt.Print("Ordenada.\n")
			}
		}
	}
}
func q2_22() {
	fmt.Print("Digite um número: ")
	var n float64
	fmt.Scan(&n)
	l := []float64{n, 1}
	c := 0
	for l[0] != float64(int(l[0])) {
		l[0] *= 10
		l[1] *= 10
		c += 1
	}
	i := 2
	for l[0]/float64(i) != 1 {
		if int(l[0])%i == 0 && int(l[1])%i == 0 {
			l[0] /= float64(i)
			l[1] /= float64(i)
		} else {
			i++
		}
	}
	fmt.Printf("%v / %v", l[0], l[1])
}
func q2_23() {
	var n int
	fmt.Print("Quantos pares de números amigos quer encontrar? ")
	fmt.Scan(&n)
	if n < 1 || n > 9 {
		fmt.Print("Número inválido ou muito alto.")
		return
	}
	counter, last := 0, 0
	i := 1
	for counter < n {
		for {
			i++
			j := int(sum(divs(i)))
			if i != last && j != last && j != i {
				if int(sum(divs(j))) == i {
					counter += 1
					last = i
					fmt.Printf("(%v, %v)\n", i, j)
					break
				}
			}
		}
	}
}
func q2_24() {
	var n int
	var s, x float64
	fmt.Print("Digite o número cujo cosseno será calculado: ")
	fmt.Scan(&x)
	fmt.Print("Digite o número de termos da série: ")
	fmt.Scan(&n)
	xDeg := 180 * float64(x) / math.Pi
	for i := 0; i <= n; i++ {
		s += math.Pow(-1, float64(i)) * math.Pow(float64(x), 2*float64(i)) / float64(fact(2*i))
	}
	fmt.Printf("\nO cosseno de %v radianos (≅%.2fº) é de %.6f.", x, xDeg, s)
}
func q2_25() {
	var n int
	var s, x float64
	fmt.Print("Digite o número cujo cosseno será calculado: ")
	fmt.Scan(&x)
	fmt.Print("Digite o número de termos da série: ")
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		s += math.Pow(x, float64(i)) / float64(fact(i))
	}
	fmt.Printf("\nO valor de e^%v é de %.6f.", x, s)
}
func q2_26() {
	var n int
	var s, x float64
	fmt.Print("Digite o número cujo seno será calculado: ")
	fmt.Scan(&x)
	fmt.Print("Digite o número de termos da série: ")
	fmt.Scan(&n)
	xDeg := 180 * float64(x) / math.Pi
	for i := 0; i <= n; i++ {
		s += math.Pow(-1, float64(i)) * math.Pow(float64(x), 2*float64(i)+1) / float64(fact(2*i+1))
	}
	fmt.Printf("\nO seno de %v radianos (≅%.2fº) é de %.6f.", x, xDeg, s)
}
func q2_27() {
	var n int
	fmt.Print("Digite o número a ser decomposto: ")
	fmt.Scan(&n)
	if n <= 1 {
		fmt.Printf("Fatoração impossível para %v.", n)
		return
	}
	fmt.Printf("%v = ", n)
	decs := []int{2}
	for n != 1 {
		for i := decs[len(decs)-1]; i <= n; i++ {
			if n%i == 0 {
				decs = append(decs, i)
				n /= i
				fmt.Printf(" %v", decs[len(decs)-1])
				if n != 1 {
					fmt.Print(" x")
				}
				break
			}
		}
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
	if n <= 1 {
		return 1
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

func divs(n int) []float64 {
	f := []float64{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			f = append(f, float64(i))
		}
	}
	return f
}
