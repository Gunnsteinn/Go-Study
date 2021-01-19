// |****************************************|
// |				EJERCICIO 2				|
// |****************************************|
//	_Crea un tipo struct persona
//	Adjunta el método hablar usando un receptor de tipo puntero
//		*persona
//	Crea un tipo interfaz humano
//		Para implícitamente implementar esa interfaz, un tipo humano debe tener el método hablar
//	Crea la función “diAlgo”
//		Haz que tome un humano como parámetro
//		Haz que llame al método hablar
//	Muestra lo siguiente en tu código
//		PUEDES pasar un valor de tipo *persona a diAlgo
//		NO puedes pasar un valor de tipo persona a diAlgo

package main

import "fmt"

const name = "Hi, I´m:"

type person struct {
	Name string
}

func (p *person) talk() string {
	return p.Name
}

type humano interface {
	talk() string
}

func saySomething(hu humano) {
	fmt.Println(name, hu.talk())
}

func main() {
	v := person{Name: "Pepe"}

	saySomething(&v)
	fmt.Println(name, v.talk())
}
