package main

import "fmt"

func main() {
	defer gti()
	fmt.Println("Salam")
}

func gti() {
	fmt.Println("Function here")
}
