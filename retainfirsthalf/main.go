package main

import (
	"fmt"
)

func retainFirsthalf(man string) string { // our program function
	char := len(man)       //used for the lenght of out character
	data := (char + 1) / 2 //used to split our character
	return man[:data]      //retain our fisrt word and leave the rest
}

func main() {
	fmt.Println(retainFirsthalf("learn2earn")) // our output "learn"
}
