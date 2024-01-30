package main

import (
	"fmt"
	"time"
)

func main() {

	var year int
	year = time.Now().Year()

	anoDeNascimento := 1990

	age := year - anoDeNascimento

	message := "Hello, Go! I'm Leanderson."

	fmt.Println("Idade:", age)
	fmt.Println(message)
}
