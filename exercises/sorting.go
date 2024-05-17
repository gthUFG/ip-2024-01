// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// // Insertion sorting

// func main() {
// 	fmt.Print("Digite a lista: ")
// 	l := inputToList()

// 	insertionSorting := Method{
// 		name: "Insertion Sorting",
// 		sort: insertion,
// 	}
// 	sortMethods := []Method{insertionSorting}

// 	fmt.Print("Digite o método a ser utilizado: ")
// 	var x int
// 	fmt.Scan(&x)
// 	time1 := time.Now()
// 	sortMethods[x].sort(l)
// 	time2 := time.Now()
// 	execTime := time2.Sub(time1)

// 	fmt.Println("Lista ordenada: ", l)
// 	fmt.Printf("O método %v foi utilizado, e foram gastos \033[1;96m%v\033[m para ordenar a lista.\n", sortMethods[x].name, execTime)
// }

// type Method struct {
// 	name string
// 	sort func(l []float64) []float64
// }

// func insertion(l []float64) []float64 {
// 	for i := 0; i < len(l)-1; i++ {
// 		for j := i + 1; j > 0; j-- {
// 			if l[j-1] < l[j] {
// 				break
// 			} else {
// 				l[j], l[j-1] = l[j-1], l[j]
// 			}
// 		}
// 	}
// 	return l
// }

// func inputToList(params ...int) []float64 {
// 	reader := bufio.NewReader(os.Stdin)
// 	inp, _ := reader.ReadString('\n')
// 	notes := toFloat64(strings.Split(strings.TrimSpace(inp), " "))
// 	if len(params) != 0 {
// 		if len(notes) != params[0] {
// 			fmt.Println("\n\033[1;31mO input está incorreto.\n\033[m")
// 			return []float64{}
// 		}
// 	}
// 	return notes
// }
// func toFloat64(l1 []string) []float64 {
// 	var floats = []float64{}
// 	for i := range l1 {
// 		l1[i] = strings.TrimSuffix(l1[i], "\n")
// 		float, _ := strconv.ParseFloat(l1[i], 32)
// 		floats = append(floats, float)
// 	}
// 	return floats
// }
