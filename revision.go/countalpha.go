package main

import "fmt"

func CountAlpha(str string) int {
	counter := 0
	for _, n := range str {
		if n >= 'a' && n <= 'z' || n  >= 'A' && n <= 'Z' {
			counter++
		}
	}
	return counter
}


func main() {
	fmt.Println(CountAlpha("hello"))
	fmt.Println(CountAlpha("abdulsamad"))
}