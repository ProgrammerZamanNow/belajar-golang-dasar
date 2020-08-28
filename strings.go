package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Contains("Eko Kurniawan", "Budi"))

	fmt.Println(strings.Split("Eko Kurniawna Khannedy", " "))

	fmt.Println(strings.ToLower("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Khannedy"))
	fmt.Println(strings.ToTitle("eko kurniawan khannedy"))

	fmt.Println(strings.Trim("      Eko Kurniawan     ", " "))
	fmt.Println(strings.ReplaceAll("Eko Joko Eko", "Eko", "Budi"))
}
