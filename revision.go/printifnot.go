package main

import "fmt"

func PrintIfNot(str string) string {

	for len(str) <= 3 || str == " " {
		return "G\n"
	}
	return "InvalidInput\n"
}



func main() {
	fmt.Println(PrintIfNot("sfghdf"))
	fmt.Println(PrintIfNot("ghj"))
	fmt.Println(PrintIfNot(""))
}
