package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

type autoGenerated struct {
	Name    string `json:"Name"`
	Code    string `json:"Code"`
	Age     int    `json:"Age"`
	Country string `json:"Country"`
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var nameArr = []string{"Rodri", "Enzo", "Fran", "Harmy", "Gun"}

const size = 70000

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr1 []autoGenerated

	var wg sync.WaitGroup
	wg.Add(size)
	var mu sync.Mutex

	for i := 0; i < size; i++ {
		go func() {
			mu.Lock()
			x := autoGenerated{
				Name:    nameArr[rand.Intn(5-0)+0],
				Code:    randStringRunes(10),
				Age:     rand.Intn(100-1) + 1,
				Country: randomdata.Country(randomdata.FullCountry),
			}
			arr1 = append(arr1, x)
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()

	rankingsJSON, _ := json.Marshal(arr1)
	err := ioutil.WriteFile("output.json", rankingsJSON, 0644)
	if err != nil {
		fmt.Println("opening config file", err.Error())
	}

}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
