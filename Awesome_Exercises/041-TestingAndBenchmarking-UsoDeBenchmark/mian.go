package main

import (
	"fmt"

	"github.com/Gunnsteinn/Go-Study/Awesome_Exercises/041-TestingAndBenchmarking-UsoDeBenchmark/stringFunc"
)

const s = `Rule 1: You can't tell where a program is going to spend its time. 
Bottlenecks occur in surprising places, so don't try to second guess and put
in a speed hack until you've proven that's where the bottleneck is.`

func main() {

	fmt.Printf("\n>>>Split & Concatenate text with NATIVE functions<<< \nRob Pike > %s\n", stringfunc.NativeJoin(stringfunc.NativeSpilt(s, " ")))
	fmt.Printf("\n>>>Split & Concatenate text with CUSTOM functions<<< \nRob Pike > %s\n\n", stringfunc.CustomJoin(stringfunc.CustomSpiltWordByWord(s)))

	fmt.Printf("\n>>>Replace text with NATIVE functions<<< \nRob Pike > %s\n", stringfunc.NativeReplace(s, "a", "@"))
	fmt.Printf("\n>>>Replace text with CUSTOM functions<<< \nRob Pike > %s\n\n", stringfunc.CustomReplace(s, "a", "@"))

}
