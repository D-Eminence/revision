package main

import "fmt"

func CountChar(str string, c rune) int {
    count := 0
		for _, r := range str {
		if r == c {
			count++
		}
	}
	return count
}


func main() {
	fmt.Println(CountChar("my world is lovely", 'o'))
}