// |****************************************|
// |				EJERCICIO 4				|
// |****************************************|
// Se pretende hacer un ejercicio donde se maneje el error a traves de strcut donde se pasa mas info que la del error
// en sí. Se construye un metodo que implementa el Error.
package main

import (
	"fmt"
	"log"
	"strconv"
)

type coordinateError struct {
	lat  string
	long string
	err  error
}

func (re coordinateError) Error() string {
	return fmt.Sprintf("[ERROR]: %v %v %v", re.lat, re.long, re.err)
}

func main() {
	_, err := geographicCoordinateValidation(31.428767,-64.1751695)
	if err != nil {
		log.Println(err)
	}
}

func geographicCoordinateValidation(lat, lng float64) (bool, error) {
	if lat > 0 || lng > 0 {
		return false, coordinateError{lat: FloatToString(lat), long: FloatToString(lng), err: fmt.Errorf("This is not SudAmerica. ")}
	}
	return true, nil
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}