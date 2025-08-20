package main

import (
	"fmt"
)

func main() {
	//var name string = "Fasia"
	var age int = 20

	email := "mianaiad_f@silpakorn.edu"
	gpa := 3.5

	firstnam, lastname := "Fasia", "Mianaiad"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f", firstnam, lastname, age, email, gpa)
}
