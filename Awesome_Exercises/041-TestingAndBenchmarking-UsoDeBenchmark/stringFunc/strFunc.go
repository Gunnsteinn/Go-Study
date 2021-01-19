/*Package stringfunc provides strings utils.

----------------------------------godoc-----------------------------------

Godoc extracts and generates documentation for Go programs.

It runs as a web server and presents the documentation as a web page.
 godoc -http : 8081

----------------------------------Gofmt-----------------------------------

	Gofmt reformats Go source code.

----------------------------------Govet-----------------------------------

	Govet is concerned with correctness.

----------------------------------Golint-----------------------------------

	Golint is concerned with coding style.
	Golint prints out style mistakes.
	Golint is in use at Google, and it seeks to match the accepted style of the open source Go project.
*/
package stringfunc

import (
	"strings"
)

// CustomSpiltWordByWord splits word by word in the text string.
func CustomSpiltWordByWord(s string) []string {
	var newArrString []string
	auxStr := ""
	for i, r := range s {
		aux := string(r)
		if aux == " " {
			newArrString = append(newArrString, auxStr)
			auxStr = ""
		} else if i == len(s)-1 {
			auxStr += string(r)
			newArrString = append(newArrString, auxStr)
		} else {
			auxStr += string(r)
		}
	}
	return newArrString
}

// NativeSpilt split slices s into all substrings separated by aux and returns a slice of the substrings between those separators.
func NativeSpilt(s, aux string) []string {
	return strings.Split(s, aux)
}

// CustomJoin concatenates the words of an array of characters.
func CustomJoin(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

// NativeJoin concatenates the elements of its first argument to create a single string.
// The separator string sep is placed between elements in the resulting string.
func NativeJoin(xs []string) string {
	return strings.Join(xs, " ")
}

// CustomReplace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
func CustomReplace(s, old, new string) string {
	var auxStr []string
	for _, r := range s {
		aux := string(r)
		if aux == old {
			aux = new
		}
		auxStr = append(auxStr, aux)
	}
	var outStr string
	for _, s := range auxStr {
		outStr += s
	}
	return outStr
}

// NativeReplace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1
// replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.
func NativeReplace(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}
