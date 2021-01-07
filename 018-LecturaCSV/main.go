// |****************************************|
// |		   EJERCICIO LECTURA CSV		|
// |****************************************|
//	El fichero cotizacion.csv contiene las cotizaciones de las empresas del IBEX35 con las siguientes columnas:
//	Nombre (nombre de la empresa), Final (precio de la acción al cierre de bolsa), Máximo (precio máximo de la acción
//	durante la jornada), Mínimo (precio mínimo de la acción durante la jornada), Volumen (Volumen al cierre de bolsa),
//	Efectivo (capitalización al cierre en miles de euros).
//
//	   - Construir una función reciba el fichero de cotizaciones y devuelva un struct con los datos del fichero.
//
//	   - Construir una función que reciba el struct devuelto por la función anterior y cree un fichero en formato
//	   csv con el mínimo, el máximo y la media de dada columna.
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	rows := openCsv()

	var arrFinal, arrMXimo, arrMNimo, arrVolumen, arrEfectivo []string
	for _, arr := range rows[1:] {
		arrFinal = append(arrFinal, arr[1])
		arrMXimo = append(arrMXimo, arr[2])
		arrMNimo = append(arrMNimo, arr[3])
		arrVolumen = append(arrVolumen, arr[4])
		arrEfectivo = append(arrEfectivo, arr[5])
	}

	matrix := [][]string{
		{"Columna", "Máximo", "Mínimo", "Media"},
		{"Final", max(strToFloat(arrFinal)), min(strToFloat(arrFinal)), media(strToFloat(arrFinal))},
		{"Máximo", max(strToFloat(arrMXimo)), min(strToFloat(arrMXimo)), media(strToFloat(arrMXimo))},
		{"Mínimo", max(strToFloat(arrMNimo)), min(strToFloat(arrMNimo)), media(strToFloat(arrMNimo))},
		{"Volumen", max(strToFloat(arrVolumen)), min(strToFloat(arrVolumen)), media(strToFloat(arrVolumen))},
		{"Efectivo", max(strToFloat(arrEfectivo)), min(strToFloat(arrEfectivo)), media(strToFloat(arrEfectivo))},
	}
	writeCsv(matrix)

	fmt.Println("Finish")
}

func openCsv() [][]string {
	f, err := os.Open("C:/Users/u631170/go/src/awesomeProject/018-LecturaCSV/cotizacion.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func min(list []float64) string {

	minimus := list[0]

	for _, value := range list {
		if value < minimus {
			minimus = value
		}
	}

	return floatTostr(minimus)
}

func max(list []float64) string {

	minimus := list[0]

	for _, value := range list {
		if value > minimus {
			minimus = value
		}
	}

	return floatTostr(minimus)
}

func media(list []float64) string {

	var val float64
	for _, f := range list {
		val = val + f
	}
	return floatTostr(val / float64(len(list)))
}

func floatTostr(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 1, 64)
}

func strToFloat(list []string) []float64 {
	var t2 []float64

	for _, i := range list {
		j, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}

	return t2
}

func writeCsv(matrix [][]string) {

	f, err := os.Create("018-LecturaCSV/resultado.csv")
	defer f.Close()

	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range matrix {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
