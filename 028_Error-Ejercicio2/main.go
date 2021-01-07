// |****************************************|
// |				EJERCICIO 2				|
// |****************************************|
// Crea un mensaje de error personalizado usando fmt.Errorf y errors.New. En este ejercicio solo
// se pretende entender el uso de la customizacion de errores.

package main

import (
	"encoding/json"
	"errors"
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
	bj, errj := aJSON(aux)
	if errj != nil {
		log.Fatalln(errj)
	}
	fmt.Println(string(bj))

	bs, errs := aSTRUCT(bj)
	if errs != nil {
		log.Fatalln(errj)
	}
	fmt.Println(bs.Phrases)
}

func aJSON(aux person) ([]byte, error) {
	bs, err := json.Marshal(aux)
	if err != nil {
		return []byte{}, fmt.Errorf("Error:  %v", err)
	}
	return bs, nil
}

func aSTRUCT(bj []byte) (person, error) {
	var toStruct person
	err := json.Unmarshal(bj, &toStruct)
	if err != nil {
		return person{}, errors.New(fmt.Sprintf("Error: %v.", err))
	}
	return toStruct, nil
}
