package main

import (
	"fmt"
)

func retainFirsthalf(man string) string {
	char := len(man)
	data := (char + 1)
	return man[:data]
}

func main() {
	fmt.Println(retainFirsthalf(learn2earn))
}
