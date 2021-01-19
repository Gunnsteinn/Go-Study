/*----------------------------------godoc-----------------------------------

Corre como un servidor web y presenta la documentación como una página web
 godoc -http : 8081
--------------------------------------------------------------------------
 Package main solve some math operations. */
package main

import (
	"fmt"
	"github.com/Gunnsteinn/Go-Study/awesomeProject/036-Documentation-Ejercicio1/humanAgeToDogAge"
)

var humanAge = 10

type canine struct {
	Name string
	Age  int
}

func main() {
	c1 := canine{Name: "Renos", Age: humanagetodogage.Convert(humanAge)}

	fmt.Println(c1)
	fmt.Println(humanAge, " Human years are", c1.Age, "canine years.")
}
