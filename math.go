package main

import "fmt"

func main()  {
	fmt.Println(Soma(20,30))
	fmt.Println(Soma(30,10))
	fmt.Println(Soma(25,25))
}

func Soma(a int, b int) int {
	return a + b
}