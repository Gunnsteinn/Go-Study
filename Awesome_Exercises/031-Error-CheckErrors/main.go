// |****************************************|
// |				Errors                  |
// |****************************************|
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("031-Error-CheckErrors/file.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	bs, errf := ioutil.ReadAll(f)
	if errf != nil {
		fmt.Println(errf)
		return
	}

	fmt.Println(string(bs))
}
