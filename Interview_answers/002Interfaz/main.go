package main

import (
	"002Interfaz/domain"
	"fmt"
)

var cashier domain.Cashier
var admin domain.Admin

func main() {
	fmt.Println("Hellow World.")
	cashier = domain.NewEmployee("Guille", 200, 10)
	contract := domain.NewContract(300)
	admin = domain.NewEmployee("Vero", 2000, 100)

	total := cashier.CalcTotal(90, 65, 93.6)
	fmt.Println(total)

	admin.DeactivateEmployee(cashier)
	total1 := cashier.CalcTotal(90, 65, 93.6)
	fmt.Println(total1)

	fmt.Println(cashier.CalculateSalary())
	fmt.Println(contract.CalculateSalary())

}
