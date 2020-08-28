package main

import "fmt"

func getFullName() (string, string, string) {
	return "Eko", "Kurniawan", "Khannedy"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}
