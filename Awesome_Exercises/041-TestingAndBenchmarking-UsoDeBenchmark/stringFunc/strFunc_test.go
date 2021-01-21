// Package stringfunc provides strings utils.
package stringfunc

import (
	"fmt"
	"testing"
)

type textExample struct {
	Phrases  string
	Expected int
}

type textExampleReplace struct {
	Phrases  string
	Expected string
}

type textArrayExample struct {
	ArrPhrases []string
	Expected   int
}

var textExamples = []textExample{
	{"There's nothing in computing that can't be broken by another level of indirection.", 13},
	{"Sockets are the X windows of IO interfaces.", 8},
	{"There's no such thing as a simple cache bug.", 9},
	{"Caches aren't architecture, they're just optimization.", 6},
	{"Languages that try to disallow idiocy become themselves idiotic.", 9},
	{"Such is modern computing: everything simple is made too complicated because it's easy to fiddle with; everything complicated stays complicated because it's hard to fix.", 25},
	{"When there is no type hierarchy you don't have to manage the type hierarchy.", 14},
	{"A smart terminal is not a smartass terminal, but rather a terminal you can educate.", 15},
	{"Not only is UNIX dead, it's starting to smell really bad.", 11},
	{"Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is", 39},
	{"Procedure names should reflect what they do; function names should reflect what they return", 14},
	{"If POSIX threads are a good thing, perhaps I don't want to know what they're better than.", 17},
	{"Narrowness of experience leads to narrowness of imagination", 8},
	{"Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming.", 28},
	{"Object-oriented design is the roman numerals of computing.", 8},
	{"Eventually, I decided that thinking was not getting me very far and it was time to try building.", 18},
	{"Using Unix is the computing equivalent of listening only to music by David Cassidy.", 14},
	{"Why would you have a language that is not theoretically exciting? Because it's very useful.", 15},
}

var textExamplesReplace = []textExampleReplace{
	{"There's nothing in computing that can't be broken by another level of indirection.", "There's nothing in computing th@t c@n't be broken by @nother level of indirection."},
	{"Sockets are the X windows of IO interfaces.", "Sockets @re the X windows of IO interf@ces."},
	{"There's no such thing as a simple cache bug.", "There's no such thing @s @ simple c@che bug."},
	{"Caches aren't architecture, they're just optimization.", "C@ches @ren't @rchitecture, they're just optimiz@tion."},
	{"Languages that try to disallow idiocy become themselves idiotic.", "L@ngu@ges th@t try to dis@llow idiocy become themselves idiotic."},
	{"Such is modern computing: everything simple is made too complicated because it's easy to fiddle with; everything complicated stays complicated because it's hard to fix.", "Such is modern computing: everything simple is m@de too complic@ted bec@use it's e@sy to fiddle with; everything complic@ted st@ys complic@ted bec@use it's h@rd to fix."},
	{"When there is no type hierarchy you don't have to manage the type hierarchy.", "When there is no type hier@rchy you don't h@ve to m@n@ge the type hier@rchy."},
	{"A smart terminal is not a smartass terminal, but rather a terminal you can educate.", "A sm@rt termin@l is not @ sm@rt@ss termin@l, but r@ther @ termin@l you c@n educ@te."},
	{"Not only is UNIX dead, it's starting to smell really bad.", "Not only is UNIX de@d, it's st@rting to smell re@lly b@d."},
	{"Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is.", "Rule 1. You c@n't tell where @ progr@m is going to spend its time. Bottlenecks occur in surprising pl@ces, so don't try to second guess @nd put in @ speed h@ck until you've proven th@t's where the bottleneck is."},
	{"Procedure names should reflect what they do; function names should reflect what they return.", "Procedure n@mes should reflect wh@t they do; function n@mes should reflect wh@t they return."},
	{"If POSIX threads are a good thing, perhaps I don't want to know what they're better than.", "If POSIX thre@ds @re @ good thing, perh@ps I don't w@nt to know wh@t they're better th@n."},
	{"Narrowness of experience leads to narrowness of imagination.", "N@rrowness of experience le@ds to n@rrowness of im@gin@tion."},
	{"Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming.", "D@t@ domin@tes. If you've chosen the right d@t@ structures @nd org@nized things well, the @lgorithms will @lmost @lw@ys be self-evident. D@t@ structures, not @lgorithms, @re centr@l to progr@mming."},
	{"Object-oriented design is the roman numerals of computing.", "Object-oriented design is the rom@n numer@ls of computing."},
	{"Eventually, I decided that thinking was not getting me very far and it was time to try building.", "Eventu@lly, I decided th@t thinking w@s not getting me very f@r @nd it w@s time to try building."},
	{"Using Unix is the computing equivalent of listening only to music by David Cassidy.", "Using Unix is the computing equiv@lent of listening only to music by D@vid C@ssidy."},
	{"Why would you have a language that is not theoretically exciting? Because it's very useful.", "Why would you h@ve @ l@ngu@ge th@t is not theoretic@lly exciting? Bec@use it's very useful."}}

var textArrayExamples = []textArrayExample{
	{[]string{"There's", "nothing", "in", "computing", "that", "can't", "be", "broken", "by", "another", "level", "of", "indirection."}, 82},
	{[]string{"Sockets", "are", "the", "X", "windows", "of", "IO", "interfaces."}, 43},
	{[]string{"There's", "no", "such", "thing", "as", "a", "simple", "cache", "bug."}, 44},
	{[]string{"Caches", "aren't", "architecture,", "they're", "just", "optimization."}, 54},
	{[]string{"Languages", "that", "try", "to", "disallow", "idiocy", "become", "themselves", "idiotic."}, 64},
	{[]string{"Such", "is", "modern", "computing:", "everything", "simple", "is", "made", "too", "complicated", "because", "it's", "easy", "to", "fiddle", "with;", "everything", "complicated", "stays", "complicated", "because", "it's", "hard", "to", "fix."}, 168},
	{[]string{"When", "there", "is", "no", "type", "hierarchy", "you", "don't", "have", "to", "manage", "the", "type", "hierarchy."}, 76},
	{[]string{"A", "smart", "terminal", "is", "not", "a", "smartass", "terminal,", "but", "rather", "a", "terminal", "you", "can", "educate."}, 83},
	{[]string{"Not", "only", "is", "UNIX", "dead,", "it's", "starting", "to", "smell", "really", "bad."}, 57},
	{[]string{"Rule", "1.", "You", "can't", "tell", "where", "a", "program", "is", "going", "to", "spend", "its", "time.", "Bottlenecks", "occur", "in", "surprising", "places,", "so", "don't", "try", "to", "second", "guess", "and", "put", "in", "a", "speed", "hack", "until", "you've", "proven", "that's", "where", "the", "bottleneck", "is."}, 211},
	{[]string{"Procedure", "names", "should", "reflect", "what", "they", "do;", "function", "names", "should", "reflect", "what", "they", "return."}, 92},
	{[]string{"If", "POSIX", "threads", "are", "a", "good", "thing,", "perhaps", "I", "don't", "want", "to", "know", "what", "they're", "better", "than."}, 89},
	{[]string{"Narrowness", "of", "experience", "leads", "to", "narrowness", "of", "imagination."}, 60},
	{[]string{"Data", "dominates.", "If", "you've", "chosen", "the", "right", "data", "structures", "and", "organized", "things", "well,", "the", "algorithms", "will", "almost", "always", "be", "self-evident.", "Data", "structures,", "not", "algorithms,", "are", "central", "to", "programming."}, 197},
	{[]string{"Object-oriented", "design", "is", "the", "roman", "numerals", "of", "computing."}, 58},
	{[]string{"Eventually,", "I", "decided", "that", "thinking", "was", "not", "getting", "me", "very", "far", "and", "it", "was", "time", "to", "try", "building."}, 96},
	{[]string{"Using", "Unix", "is", "the", "computing", "equivalent", "of", "listening", "only", "to", "music", "by", "David", "Cassidy."}, 83},
	{[]string{"Why", "would", "you", "have", "a", "language", "that", "is", "not", "theoretically", "exciting?", "Because", "it's", "very", "useful."}, 91},
}

//func TestCustomSpiltWordByWord(t *testing.T) {
//	for _, value := range textExamples {
//		got := CustomSpiltWordByWord(value.Phrases)
//		if len(got) != value.Expected {
//			t.Errorf("got %v expected %v", len(got), value.Expected)
//		}
//	}
//}
//
//func ExampleCustomSpiltWordByWord() {
//	// >>> {"There's nothing in computing that can't be broken by another level of indirection."}
//	// >>> {"Sockets are the X windows of IO interfaces."}
//	// >>> {"There's no such thing as a simple cache bug."}
//	// >>> {"Caches aren't architecture, they're just optimization."}
//	// >>> {"Languages that try to disallow idiocy become themselves idiotic."}
//	// >>> {"Such is modern computing: everything simple is made too complicated because it's easy to fiddle with; everything complicated stays complicated because it's hard to fix."}
//	// >>> {"When there is no type hierarchy you don't have to manage the type hierarchy."}
//	// >>> {"A smart terminal is not a smartass terminal, but rather a terminal you can educate."}
//	// >>> {"Not only is UNIX dead, it's starting to smell really bad."}
//	// >>> {"Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is"}
//	// >>> {"Procedure names should reflect what they do; function names should reflect what they return"}
//	// >>> {"If POSIX threads are a good thing, perhaps I don't want to know what they're better than."}
//	// >>> {"Narrowness of experience leads to narrowness of imagination"}
//	// >>> {"Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming."}
//	// >>> {"Object-oriented design is the roman numerals of computing."}
//	// >>> {"Eventually, I decided that thinking was not getting me very far and it was time to try building."}
//	// >>> {"Using Unix is the computing equivalent of listening only to music by David Cassidy."}
//	// >>> {"Why would you have a language that is not theoretically exciting? Because it's very useful."}
//
//	for _, value := range textExamples {
//		fmt.Println(len(CustomSpiltWordByWord(value.Phrases)))
//	}
//	//Output:
//	// 13
//	// 8
//	// 9
//	// 6
//	// 9
//	// 25
//	// 14
//	// 15
//	// 11
//	// 39
//	// 14
//	// 17
//	// 8
//	// 28
//	// 8
//	// 18
//	// 14
//	// 15
//}
//
//func BenchmarkCustomSpiltWordByWord(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		for _, value := range textExamples {
//			CustomSpiltWordByWord(value.Phrases)
//		}
//	}
//}

func TestNativeSpilt(t *testing.T) {
	for _, value := range textExamples {
		got := NativeSpilt(value.Phrases, " ")
		if len(got) != value.Expected {
			t.Errorf("got %v expected %v", len(got), value.Expected)
		}
	}
}

func ExampleNativeSpilt() {
	// >>> {"There's nothing in computing that can't be broken by another level of indirection."}
	// >>> {"Sockets are the X windows of IO interfaces."}
	// >>> {"There's no such thing as a simple cache bug."}
	// >>> {"Caches aren't architecture, they're just optimization."}
	// >>> {"Languages that try to disallow idiocy become themselves idiotic."}
	// >>> {"Such is modern computing: everything simple is made too complicated because it's easy to fiddle with; everything complicated stays complicated because it's hard to fix."}
	// >>> {"When there is no type hierarchy you don't have to manage the type hierarchy."}
	// >>> {"A smart terminal is not a smartass terminal, but rather a terminal you can educate."}
	// >>> {"Not only is UNIX dead, it's starting to smell really bad."}
	// >>> {"Rule 1. You can't tell where a program is going to spend its time. Bottlenecks occur in surprising places, so don't try to second guess and put in a speed hack until you've proven that's where the bottleneck is"}
	// >>> {"Procedure names should reflect what they do; function names should reflect what they return"}
	// >>> {"If POSIX threads are a good thing, perhaps I don't want to know what they're better than."}
	// >>> {"Narrowness of experience leads to narrowness of imagination"}
	// >>> {"Data dominates. If you've chosen the right data structures and organized things well, the algorithms will almost always be self-evident. Data structures, not algorithms, are central to programming."}
	// >>> {"Object-oriented design is the roman numerals of computing."}
	// >>> {"Eventually, I decided that thinking was not getting me very far and it was time to try building."}
	// >>> {"Using Unix is the computing equivalent of listening only to music by David Cassidy."}
	// >>> {"Why would you have a language that is not theoretically exciting? Because it's very useful."}

	for _, value := range textExamples {
		fmt.Println(len(NativeSpilt(value.Phrases, " ")))
	}
	//Output:
	// 13
	// 8
	// 9
	// 6
	// 9
	// 25
	// 14
	// 15
	// 11
	// 39
	// 14
	// 17
	// 8
	// 28
	// 8
	// 18
	// 14
	// 15
}

func BenchmarkNativeSpilt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range textExamples {
			NativeSpilt(value.Phrases, " ")
		}
	}
}

func TestCustomJoin(t *testing.T) {
	for _, value := range textArrayExamples {
		got := CustomJoin(value.ArrPhrases)
		if len(got) != value.Expected {
			t.Errorf("got %v expected %v", len(got), value.Expected)
		}

	}
}

func ExampleCustomJoin() {
	//{[]string{"There's", "nothing", "in", "computing", "that", "can't", "be", "broken", "by", "another", "level", "of", "indirection."}, 82},
	//{[]string{"Sockets", "are", "the", "X", "windows", "of", "IO", "interfaces."}, 43},
	//{[]string{"There's", "no", "such", "thing", "as", "a", "simple", "cache", "bug."}, 44},
	//{[]string{"Caches", "aren't", "architecture,", "they're", "just", "optimization."}, 54},
	//{[]string{"Languages", "that", "try", "to", "disallow", "idiocy", "become", "themselves", "idiotic."}, 64},
	//{[]string{"Such", "is", "modern", "computing:", "everything", "simple", "is", "made", "too", "complicated", "because", "it's", "easy", "to", "fiddle", "with;", "everything", "complicated", "stays", "complicated", "because", "it's", "hard", "to", "fix."}, 168},
	//{[]string{"When", "there", "is", "no", "type", "hierarchy", "you", "don't", "have", "to", "manage", "the", "type", "hierarchy."}, 76},
	//{[]string{"A", "smart", "terminal", "is", "not", "a", "smartass", "terminal,", "but", "rather", "a", "terminal", "you", "can", "educate."}, 83},
	//{[]string{"Not", "only", "is", "UNIX", "dead,", "it's", "starting", "to", "smell", "really", "bad."}, 57},
	//{[]string{"Rule", "1.", "You", "can't", "tell", "where", "a", "program", "is", "going", "to", "spend", "its", "time.", "Bottlenecks", "occur", "in", "surprising", "places,", "so", "don't", "try", "to", "second", "guess", "and", "put", "in", "a", "speed", "hack", "until", "you've", "proven", "that's", "where", "the", "bottleneck", "is."}, 211},
	//{[]string{"Procedure", "names", "should", "reflect", "what", "they", "do;", "function", "names", "should", "reflect", "what", "they", "return."}, 92},
	//{[]string{"If", "POSIX", "threads", "are", "a", "good", "thing,", "perhaps", "I", "don't", "want", "to", "know", "what", "they're", "better", "than."}, 89},
	//{[]string{"Narrowness", "of", "experience", "leads", "to", "narrowness", "of", "imagination."}, 60},
	//{[]string{"Data", "dominates.", "If", "you've", "chosen", "the", "right", "data", "structures", "and", "organized", "things", "well,", "the", "algorithms", "will", "almost", "always", "be", "self-evident.", "Data", "structures,", "not", "algorithms,", "are", "central", "to", "programming."}, 197},
	//{[]string{"Object-oriented", "design", "is", "the", "roman", "numerals", "of", "computing."}, 58},
	//{[]string{"Eventually,", "I", "decided", "that", "thinking", "was", "not", "getting", "me", "very", "far", "and", "it", "was", "time", "to", "try", "building."}, 96},
	//{[]string{"Using", "Unix", "is", "the", "computing", "equivalent", "of", "listening", "only", "to", "music", "by", "David", "Cassidy."}, 83},
	//{[]string{"Why", "would", "you", "have", "a", "language", "that", "is", "not", "theoretically", "exciting?", "Because", "it's", "very", "useful."}, 91}

	for _, value := range textArrayExamples {
		fmt.Println(len(CustomJoin(value.ArrPhrases)))
	}
	//Output:
	// 82
	// 43
	// 44
	// 54
	// 64
	// 168
	// 76
	// 83
	// 57
	// 211
	// 92
	// 89
	// 60
	// 197
	// 58
	// 96
	// 83
	// 91
}

func BenchmarkCustomJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range textArrayExamples {
			CustomJoin(value.ArrPhrases)
		}
	}
}

func TestNativeJoin(t *testing.T) {
	for _, value := range textArrayExamples {
		got := NativeJoin(value.ArrPhrases)
		if len(got) != value.Expected {
			t.Errorf("got %v expected %v", len(got), value.Expected)
		}

	}
}

func ExampleNativeJoin() {
	//{[]string{"There's", "nothing", "in", "computing", "that", "can't", "be", "broken", "by", "another", "level", "of", "indirection."}, 82},
	//{[]string{"Sockets", "are", "the", "X", "windows", "of", "IO", "interfaces."}, 43},
	//{[]string{"There's", "no", "such", "thing", "as", "a", "simple", "cache", "bug."}, 44},
	//{[]string{"Caches", "aren't", "architecture,", "they're", "just", "optimization."}, 54},
	//{[]string{"Languages", "that", "try", "to", "disallow", "idiocy", "become", "themselves", "idiotic."}, 64},
	//{[]string{"Such", "is", "modern", "computing:", "everything", "simple", "is", "made", "too", "complicated", "because", "it's", "easy", "to", "fiddle", "with;", "everything", "complicated", "stays", "complicated", "because", "it's", "hard", "to", "fix."}, 168},
	//{[]string{"When", "there", "is", "no", "type", "hierarchy", "you", "don't", "have", "to", "manage", "the", "type", "hierarchy."}, 76},
	//{[]string{"A", "smart", "terminal", "is", "not", "a", "smartass", "terminal,", "but", "rather", "a", "terminal", "you", "can", "educate."}, 83},
	//{[]string{"Not", "only", "is", "UNIX", "dead,", "it's", "starting", "to", "smell", "really", "bad."}, 57},
	//{[]string{"Rule", "1.", "You", "can't", "tell", "where", "a", "program", "is", "going", "to", "spend", "its", "time.", "Bottlenecks", "occur", "in", "surprising", "places,", "so", "don't", "try", "to", "second", "guess", "and", "put", "in", "a", "speed", "hack", "until", "you've", "proven", "that's", "where", "the", "bottleneck", "is."}, 211},
	//{[]string{"Procedure", "names", "should", "reflect", "what", "they", "do;", "function", "names", "should", "reflect", "what", "they", "return."}, 92},
	//{[]string{"If", "POSIX", "threads", "are", "a", "good", "thing,", "perhaps", "I", "don't", "want", "to", "know", "what", "they're", "better", "than."}, 89},
	//{[]string{"Narrowness", "of", "experience", "leads", "to", "narrowness", "of", "imagination."}, 60},
	//{[]string{"Data", "dominates.", "If", "you've", "chosen", "the", "right", "data", "structures", "and", "organized", "things", "well,", "the", "algorithms", "will", "almost", "always", "be", "self-evident.", "Data", "structures,", "not", "algorithms,", "are", "central", "to", "programming."}, 197},
	//{[]string{"Object-oriented", "design", "is", "the", "roman", "numerals", "of", "computing."}, 58},
	//{[]string{"Eventually,", "I", "decided", "that", "thinking", "was", "not", "getting", "me", "very", "far", "and", "it", "was", "time", "to", "try", "building."}, 96},
	//{[]string{"Using", "Unix", "is", "the", "computing", "equivalent", "of", "listening", "only", "to", "music", "by", "David", "Cassidy."}, 83},
	//{[]string{"Why", "would", "you", "have", "a", "language", "that", "is", "not", "theoretically", "exciting?", "Because", "it's", "very", "useful."}, 91}

	for _, value := range textArrayExamples {
		fmt.Println(len(NativeJoin(value.ArrPhrases)))
	}
	//Output:
	// 82
	// 43
	// 44
	// 54
	// 64
	// 168
	// 76
	// 83
	// 57
	// 211
	// 92
	// 89
	// 60
	// 197
	// 58
	// 96
	// 83
	// 91
}

func BenchmarkNativeJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range textArrayExamples {
			NativeJoin(value.ArrPhrases)
		}
	}
}

func TestCustomReplace(t *testing.T) {
	for _, value := range textExamplesReplace {
		got := CustomReplace(value.Phrases, "a", "@")
		if got != value.Expected {
			t.Errorf("got %v expected %v", got, value.Expected)
		}
	}
}

func ExampleCustomReplace() {
	for _, value := range textExamplesReplace {
		fmt.Println(CustomReplace(value.Phrases, "a", "@"))
	}
	//Output:
	// There's nothing in computing th@t c@n't be broken by @nother level of indirection.
	// Sockets @re the X windows of IO interf@ces.
	// There's no such thing @s @ simple c@che bug.
	// C@ches @ren't @rchitecture, they're just optimiz@tion.
	// L@ngu@ges th@t try to dis@llow idiocy become themselves idiotic.
	// Such is modern computing: everything simple is m@de too complic@ted bec@use it's e@sy to fiddle with; everything complic@ted st@ys complic@ted bec@use it's h@rd to fix.
	// When there is no type hier@rchy you don't h@ve to m@n@ge the type hier@rchy.
	// A sm@rt termin@l is not @ sm@rt@ss termin@l, but r@ther @ termin@l you c@n educ@te.
	// Not only is UNIX de@d, it's st@rting to smell re@lly b@d.
	// Rule 1. You c@n't tell where @ progr@m is going to spend its time. Bottlenecks occur in surprising pl@ces, so don't try to second guess @nd put in @ speed h@ck until you've proven th@t's where the bottleneck is.
	// Procedure n@mes should reflect wh@t they do; function n@mes should reflect wh@t they return.
	// If POSIX thre@ds @re @ good thing, perh@ps I don't w@nt to know wh@t they're better th@n.
	// N@rrowness of experience le@ds to n@rrowness of im@gin@tion.
	// D@t@ domin@tes. If you've chosen the right d@t@ structures @nd org@nized things well, the @lgorithms will @lmost @lw@ys be self-evident. D@t@ structures, not @lgorithms, @re centr@l to progr@mming.
	// Object-oriented design is the rom@n numer@ls of computing.
	// Eventu@lly, I decided th@t thinking w@s not getting me very f@r @nd it w@s time to try building.
	// Using Unix is the computing equiv@lent of listening only to music by D@vid C@ssidy.
	// Why would you h@ve @ l@ngu@ge th@t is not theoretic@lly exciting? Bec@use it's very useful.
}

func BenchmarkCustomReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range textExamplesReplace {
			CustomReplace(value.Phrases, "a", "@")
		}
	}
}

func TestNativeReplace(t *testing.T) {
	for _, value := range textExamplesReplace {
		got := NativeReplace(value.Phrases, "a", "@")
		if got != value.Expected {
			t.Errorf("got %v expected %v", got, value.Expected)
		}
	}
}

func ExampleNativeReplace() {
	for _, value := range textExamplesReplace {
		fmt.Println(NativeReplace(value.Phrases, "a", "@"))
	}
	//Output:
	// There's nothing in computing th@t c@n't be broken by @nother level of indirection.
	// Sockets @re the X windows of IO interf@ces.
	// There's no such thing @s @ simple c@che bug.
	// C@ches @ren't @rchitecture, they're just optimiz@tion.
	// L@ngu@ges th@t try to dis@llow idiocy become themselves idiotic.
	// Such is modern computing: everything simple is m@de too complic@ted bec@use it's e@sy to fiddle with; everything complic@ted st@ys complic@ted bec@use it's h@rd to fix.
	// When there is no type hier@rchy you don't h@ve to m@n@ge the type hier@rchy.
	// A sm@rt termin@l is not @ sm@rt@ss termin@l, but r@ther @ termin@l you c@n educ@te.
	// Not only is UNIX de@d, it's st@rting to smell re@lly b@d.
	// Rule 1. You c@n't tell where @ progr@m is going to spend its time. Bottlenecks occur in surprising pl@ces, so don't try to second guess @nd put in @ speed h@ck until you've proven th@t's where the bottleneck is.
	// Procedure n@mes should reflect wh@t they do; function n@mes should reflect wh@t they return.
	// If POSIX thre@ds @re @ good thing, perh@ps I don't w@nt to know wh@t they're better th@n.
	// N@rrowness of experience le@ds to n@rrowness of im@gin@tion.
	// D@t@ domin@tes. If you've chosen the right d@t@ structures @nd org@nized things well, the @lgorithms will @lmost @lw@ys be self-evident. D@t@ structures, not @lgorithms, @re centr@l to progr@mming.
	// Object-oriented design is the rom@n numer@ls of computing.
	// Eventu@lly, I decided th@t thinking w@s not getting me very f@r @nd it w@s time to try building.
	// Using Unix is the computing equiv@lent of listening only to music by D@vid C@ssidy.
	// Why would you h@ve @ l@ngu@ge th@t is not theoretic@lly exciting? Bec@use it's very useful.
}

func BenchmarkNativeReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, value := range textExamplesReplace {
			NativeReplace(value.Phrases, "a", "@")
		}
	}
}
