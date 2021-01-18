// |****************************************|
// |				EJERCICIO 1				|
// |****************************************|
// Simple ejercicio donde se explica log.Fatalln().
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Name     string
	LastName string
	Phrases  []string
}

func main() {
	aux := person{
		Name:     "Willy",
		LastName: "Colo",
		Phrases:  []string{"Hello", "World"},
	}

	// Struct to json
	bs, err := json.Marshal(aux)
	if err != nil {
		// Call os.Exit(). kill all goroutines and defer functions.
		log.Fatalln(err)
	}

	// pure format, slice of byte.
	fmt.Println(bs)
	// cast byte to string
	fmt.Println(string(bs))
}
