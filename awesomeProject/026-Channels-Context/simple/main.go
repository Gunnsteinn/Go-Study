// En servidores de Go, cada request entrante es manejado por su propia gorutina. Los request
// handlers usualmente inicializan gorutinas adicionales para acceder a backends tales como
// bases de datos o servicios RPC. El conjunto de gorutinas que trabajan en un request
// típicamente necesitan acceso a valores que son específicos del request, tales como identidad
// del usuario final, tokens de autorización y un tiempo límite del request. Cuando un request es
// cancelado o termina su tiempo, todas las gorutinas trabajando para ese request deberían
// finalizar rápidamente de manera que el sistema pueda reclamar cualquier recurso que estén
// usando. En Google, desarrollaron un paquete context que hace fácil pasar valores que
// pertenecen al scope del request, señales de cancelación y deadlines a través de APIs hacia las
// gorutinas que están involucradas en manejar un request. El paquete está públicamente
// disponible como context. Este artículo describe cómo usar el paquete, provee un ejemplo
// completo y funcional.
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("chqueo de error 1:", ctx.Err())
	fmt.Println("num gorutinas 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Trabajando", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("chequeo de error:", ctx.Err())
	fmt.Println("num gorutinas 2:", runtime.NumGoroutine())

	fmt.Println("A punto de cancelar context.")
	cancel()
	fmt.Println("context cancelado.")

	time.Sleep(time.Second * 2)
	fmt.Println("chequeo de error 3:", ctx.Err())
	fmt.Println("num gorutinas 3:", runtime.NumGoroutine())
}
