package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blacklistAdmin(name string) bool {
//	return name == "admin"
//}
//
//func blacklistRoot(name string) bool {
//	return name == "root"
//}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "root"
	})
}
