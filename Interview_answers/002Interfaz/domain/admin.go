package domain

type Admin interface {
	DeactivateEmployee(c Cashier)
}

func (e *Employee) DeactivateEmployee(c Cashier) {
	c.deactivate()
}
