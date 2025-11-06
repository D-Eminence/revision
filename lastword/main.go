package main

import (
	"fmt"
	"strings"
)

func lastword(r string) string { //our program function
	r = strings.TrimSpace(r)  //used to trim space at the start and end
	part := strings.Fields(r) //it split the string by space
	if len(part) == 0 {       //if there is no word it should print an empty string
		return ""
	}
	return part[len(part)-1] //it return the last word
}

func main() {
	fmt.Println(lastword("hello word")) //our output "word"
	fmt.Println(lastword("hello"))      //our output "hello"
	fmt.Println(lastword(""))           //our output "empty string"

}
