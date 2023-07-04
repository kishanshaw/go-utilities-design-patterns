package main

import (
	"fmt"
)

func main() {
	var mapUsers = make(map[string]int)
	mapUsers["Kishan"] = 28
	mapUsers["Jyoti"] = 21
	fmt.Println(mapUsers)
}
