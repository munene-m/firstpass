package main

import (
	"flag"
	"fmt"

	"github.com/munene-m/firstpass/internal/utils"
)

func main(){
	length := flag.Int("l", 15, "Length of the password")
	useUpper := flag.Bool("upper", false, "Use uppercase characters")
	useLower := flag.Bool("lower", false, "Include lowercase characters")
	useNumbers := flag.Bool("numbers", false, "Include numbers")
	useSpecial := flag.Bool("special", false, "Include special characters")
	count := flag.Int("c", 1, "number of passwords to generate")
	copy := flag.Bool("cp", false, "Copy to clipboard")

	flag.Parse()

	for i := 0; i < *count; i++ {
		password := utils.GeneratePassword(*length, *useUpper, *useLower, *useSpecial, *useNumbers, *copy)
		fmt.Printf("Generated password: %s\n", password)
	}
}
