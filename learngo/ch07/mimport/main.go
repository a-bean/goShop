package main

import (
	"fmt"
	"learngo/ch07/nest"
)

func main() {
	p := nest.New()
	fmt.Println(p)
	fmt.Println(p.Name)
}
