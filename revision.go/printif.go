package main

import "fmt"

func PrintIf(str string) string {

	if len(str) >= 3 || str == " " {
		return "G\n"
	}
	return "InvalidInput\n"
}


func main() {
	fmt.Println(PrintIf("hello"))
	fmt.Println(PrintIf("Goi"))
	fmt.Println(PrintIf(""))
}