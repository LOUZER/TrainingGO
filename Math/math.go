package main 

import "fmt"

func main() {
	resultado := soma(1, 1)
	fmt.Printf("%v", resultado)
}

func soma (a int, b int ) int{
	return a + b
}