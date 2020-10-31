package stringBuilder

import "strings"

func buildString() {
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
}

func concatString() {
	var str string
	for i := 0; i < 1000; i++ {
		str += "a"
	}
}

func appendString() {
	var str string
	for i := 0; i < 1000; i++ {
		str = string(append([]byte(str), []byte("a")...))
	}
}
