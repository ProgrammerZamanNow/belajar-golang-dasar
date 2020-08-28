package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Eko")
	// helper.sayGoodbye("Eko") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
