package main

import (
	"fmt"
)

func main() {
	var l, n int
	fmt.Print("\n\033[0;36mQual lista? \033[m")
	fmt.Scanln(&l)
	var quests = [][]func(){QuestsList01, QuestsList02}
	if l <= len(quests) && l >= 1 {
		fmt.Print("\n\033[0;36mQual exercício? \033[m")
		fmt.Scanln(&n)
		if n <= len(quests[l-1]) && n >= 1 {
			fmt.Printf("\n\033[1;36m------------- Questão %v -------------\033[1;37m\n\n", n)
			quests[l-1][n-1]()
			fmt.Printf("\n\n\033[1;36m-------------------------------------\033[1;37m\n\n")
		} else {
			fmt.Printf("Questão inválida. Há %v questões.\n", len(quests[l-1]))
		}
	} else {
		fmt.Printf("Lista inválida. Há somente %v listas.\n", len(quests))
	}
}
