package domain

import (
	"fmt"
	"math/rand"
)

type User struct {
	id       int
	name     string
	basicPay int
	pension  int
}

type Contract struct {
	id       int
	basicPay int
}

type Employee struct {
	User
	active bool
}

func (p *User) CalculateSalary() int {
	return p.basicPay + p.pension
}

func (p *Contract) CalculateSalary() int {
	return p.basicPay
}

func NewEmployee(name string, basicPay int, pension int) *Employee {
	return &Employee{
		User: User{
			id:       rand.Int(),
			name:     name,
			basicPay: basicPay,
			pension:  pension,
		},
		active: true,
	}
}

func NewContract(basicPay int) *Contract {
	return &Contract{
		id:       rand.Int(),
		basicPay: basicPay,
	}
}

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
	CalculateSalary() int
}

func (e *Employee) deactivate() {
	e.active = false
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.active {
		fmt.Println("Cashier deactivated")
		return 0
	}
	var sum float32
	for _, v := range item {
		sum += v
	}
	return sum * 1.15
}
