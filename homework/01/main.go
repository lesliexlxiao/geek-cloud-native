package main

import "fmt"

func main() {
	originalStrings := []string{"I","am","stupid","and","weak"}

	targetStrings := make([]string, 0)

	for _, value := range originalStrings {
		if value == "stupid" {
			value = "smart"
		} else if value == "weak" {
			value = "strong"
		}

		targetStrings = append(targetStrings, value)
	}

	fmt.Println(targetStrings)
}
