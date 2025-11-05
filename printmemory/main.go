package main

import (
	"fmt"
	"strings"
)

//main func contains an array of 10 byte
func PrintMemory(a [10]byte) {
// str store the ascii values
	str := ""
// it loops through the array
	for i, nbr := range a {
// convert the array values into a 2 didgit hexadecimal values
		fmt.Printf("%.2x",nbr)
// the condition is just to help us print 4 chars on each row and spaces btw the chars 
// just like dis
//  12 ab 3c 4f
// 5a 00 ff 9e
		if ((i+1)%4 == 0 && i != 0) || i == len(a)-1 {
			fmt.Println()
		} else {
			fmt.Print("-")
		}
// check if the values are set of printable character then store it under the str 
// and if the char are not printable it should represent it with 
		if nbr >= 33 && nbr <= 126{
			str += string(rune(nbr))
		} else {
			str += "."
		}
	}
	fmt.Println(str + strings.Repeat(".", 10-len(a)))
}
func main(){
	data := [10]byte{72, 101, 108, 108, 111, 40,33,60,70,120}
	PrintMemory(data)
}