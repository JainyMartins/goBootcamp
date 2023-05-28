package exercicio4

import "fmt"

func Exercicio4() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}

	fmt.Println(employees["Benjamin"])
	for key, element := range employees {
		if element > 21 {
			fmt.Println("Nome do funcionário:", key, "=>", "Idade:", element)
		}
		continue
	}
	employees["Frederico"] = 25
	fmt.Println(employees)

	delete(employees, "Pedro")
	fmt.Println(employees)
}
