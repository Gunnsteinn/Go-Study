/*Package mathme provides math solutions.

---------------------------------go doc----------------------------------

Go doc imprime la documentación para un paquete, constante, función, tipo, var o método.

Go doc acepta cero, uno, o dos argumentos.

 go doc <paquete>
 go doc <sim>[.<método>]
 go doc [<paquete>.]<sim>[.<método>]
 go doc [<paquete>.][<sim>.]<método>

----------------------------------godoc-----------------------------------

Corre como un servidor web y presenta la documentación como una página web
 godoc -http : 8081
--------------------------------------------------------------------------
*/
package mathme

// Sum add several numbers.
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Mul multiply two numbers.
func Mul(x, y float64) float64 {
	return x * y
}

// Div divide two numbers.
func Div(x, y float64) float64 {
	return x / y
}
