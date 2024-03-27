package core

import (
	"fmt"
	"log"
)

func ConfirmAction(action string) bool {
	var input string

	fmt.Printf("%s? (y): ", action)

	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}

	switch input {
	case "y", "yes", "Y", "YES":
		return true
	default:
		return false
	}
}
