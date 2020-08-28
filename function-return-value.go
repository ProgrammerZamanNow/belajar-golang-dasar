package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
