package main

import (
	"./test"
	"fmt"
)

func main() {
	test.TT = "jjjjjj"
	test.Name_a()
	test.Name_b()
	fmt.Println(test.TT)
}
