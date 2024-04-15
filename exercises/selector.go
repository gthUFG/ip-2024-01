package main

import (
	"fmt"
)

func main() {
	fmt.Print("\n\033[0;36mQual exercício? \033[m")
	var n int
	fmt.Scanln(&n)
	var quests = []func(){Q01, Q02, Q03, Q04, Q05, Q06, Q07, Q08, Q09, Q10, Q11, Q12, Q13, Q14, Q15, Q16, Q17, Q18, Q19, Q20}
	if n <= len(quests) && n >= 1 {
		fmt.Printf("\n\033[1;36m------------- Questão %v -------------\033[1;37m\n\n", n)
		quests[n-1]()
		fmt.Printf("\n\n\033[1;36m-------------------------------------\033[1;37m\n\n", n)
	} else {
		fmt.Printf("Questão inválida. Há %v questões.\n", len(quests))
	}
}
