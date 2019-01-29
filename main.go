package main

import (
	"fmt"
	"strconv"
)

func theQuestion(i int) int {
	g := 10 - i
	r, err := strconv.Atoi(fmt.Sprintf("%d%d", i, g))

	if err != nil {
		panic(err)
	}

	return r
}

func main() {
	fmt.Println("vim-go")
}
