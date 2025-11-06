package main

import (
	"fmt"
	"strings"
)

func firstword(d string) string {
	d = strings.TrimSpace(d)   // trim the start and end space
	parts := strings.Fields(d) //splitting by space
	if len(parts) == 0 {
		return "" //if there is no word it print empty string
	}
	return parts[0] // print the first part
}

func main() {
	fmt.Println(firstword("hello world"))      //our output "hello"
	fmt.Println(firstword("fatima is a girl")) //our output "fatima"
}
