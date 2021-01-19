/*Package humanagetodogage provides convert solutions.

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
package humanagetodogage

// Convert is a func that convert human years in canine years.
func Convert(age int) int {
	return age * 7
}
