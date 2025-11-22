package main

import "fmt"

func RetainFirstHalf(str string) string {
	can := len(str)
	data := (can-1)/2
	return str[:data]
}


func main() {
	fmt.Println(RetainFirstHalf("hello world"))
}