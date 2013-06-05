package main

import "fmt"

//INCOMPLETE
func invalidIndex(st string) int {
	var i int
	var brackets = map[string]string{
		"{": "}",
		"(": ")",
		"[": "]",
	}
	var buffer []string
	for _, s := range st {
		if b, ok := brackets[string(s)]; ok {
			buffer = append(buffer, string(s))
		} else {
			lastB := buffer[len(buffer)-1]
			buffer = buffer[0:(len(buffer) - 2)]
			if string(s) != brackets[lastB] {
				fmt.Println("Found a break")
				return i
			}
			fmt.Println(buffer[len(buffer)-1])
			fmt.Println(b)
			fmt.Println(string(s))
		}
		i++
	}
	return i
}

func main() {
	input := "{[[])]}"
	index := invalidIndex(input)
	fmt.Println(index)
}
