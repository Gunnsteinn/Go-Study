# Golang:

1. Enumerar todos los tipos de datos en golang

- Integers
  - Signed
    - int
    - int8
    - int16
    - int32
    - int64
  - Unsigned
    - uint
    - uint8
    - uint16
    - uint32
    - uint64
    - uintptr
  - Floats
    - float32
    - float64
  - Complex Numbers
    - complex64
    - complex128
  - Byte
  - Rune
  - String
  - Boolean

---

2. ¿Qué es la interfaz? ¿por que lo usas?

```
En Golang, una interfaz es un conjunto de firmas de metodos. Si un tipo tiene una definicion para esos metodos, se dice que implementa la interfaz. A diferencia de otros lenguajes, la asociacion de un tipo con una interfaz es implicita.

La principal ventaja del uso de interfaces en un lenguaje estatico como Golang es la reutilizacion de codigo.

Un tipo puede implementar mas de una interfaz.
Ejemplo: 002Interfaz
```

---

3. ¿Qué son la concurrencia y el paralelismo? y cual es la diferencia entre ambos?

```
La concurrencia es cuando dos o más tareas pueden empezar, ejecutarse y completarse en períodos de tiempos superpuestos. Esto no quiere decir que ambos procesos corran a la misma vez. Mientras que en el paralelismo las tareas se ejecutarán literalmente a la misma vez, esto solo es posible cuando tenemos más de un procesador de otra manera es imposible realizar tareas en paralelo y posiblemente hablemos de tareas ejecutándose concurrentemente.

```

---

4. difference between goroutines and threads?

```
Sample text here...
```

---

5. what are channels for?

```
Sample text here...
```

---

6. can you do something in goroutines without channels?

```
Sample text here...
```

7. difference between arrays and slices?

```
Sample text here...
```

8. what is a closure? define it.

```
Sample text here...
```

9. what is the size of int(not int8, int16)? where is the size of int defined?

```
Sample text here...
```

10. What is runtime? runtime package?

```
Sample text here...
```

11. how can you get how many cores your computer has?

```
Sample text here...
```

12. How would you tell a goroutine to use less less core than what you have?

```
Sample text here...
```

13. how would you determine the type of a variable? which package to use for it?

```
Sample text here...
```

14. What all types a map can store?

```
Sample text here...
```

15. ¿Qué es Runa?

Las runes o runas son sinónimo de un tipo int32. Es el tipo de variable por defecto cuando defines un carácter, utilizamos comillas sencillas para declararlo. Si no especificas byte u otro tipo de dato, go dará por sentado que se trata de una rune.

```
package main

import "fmt"

func main() {
	var runa rune = 'A'
	fmt.Printf("type:%T, value:%v\n", runa, runa)
}
```

> output: type:int32, value:65

16. what are services in golang?

```
Sample text here...
```

17. Define a singleton?

```
Sample text here...
```

18. Abstraction? Interfaces? Implementation?

```
Sample text here...
```

19. When do you use a structure in golang?

```
Sample text here...
```

20. How do define struct type functions? and how can you access those functions? Can you define such functions to other datatypes as well?

```
Sample text here...
```

21. Why are there no classes in Go?

```
Sample text here...
```

22. Compile time and runtime?

```
Sample text here...
```

23. Do you need to convert the type of a variable of interface{} type passed in a function as an argument?

```
Sample text here...
```

24. Why are goroutines light-weight? Explain reason.

```
Sample text here...
```

25. Do threads share memory? how?

```
Sample text here...
```

26. if capacity is not defined in slice, what would the capacity be?

```
Sample text here...
```
