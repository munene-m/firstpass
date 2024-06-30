package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/munene-m/firstpass/internal/utils"
)

func main(){
	length := flag.Int("l", 15, "Length of the password")
	useUpper := flag.Bool("U", false, "Use uppercase characters")
	useLower := flag.Bool("L", false, "Include lowercase characters")
	useNumbers := flag.Bool("n", false, "Include numbers")
	useSpecial := flag.Bool("s", false, "Include special characters")
	count := flag.Int("c", 1, "number of passwords to generate")
	copy := flag.Bool("cp", false, "Copy to clipboard")

	flag.Parse()

	var passwords []string

	for i := 0; i < *count; i++ {
		password := utils.GeneratePassword(*length, *useUpper, *useLower, *useSpecial, *useNumbers)
		passwords = append(passwords, password)
	}

	if *copy {
		clipboardContent := strings.Join(passwords, "\n")
		err := utils.CopyToClipboard(clipboardContent)
		if err != nil {
			log.Fatalf("Failed to copy to clipboard: %v", err)
		}
		fmt.Println("Password(s) copied to clipboard.")
	} else {
		for _, password := range passwords {
			fmt.Printf("Generated password: %s\n ", password)
		}
	}
	
}

