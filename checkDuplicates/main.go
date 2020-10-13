package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	//implement here
	uniqueMap := make(map[string]bool)
	var uniqueNames []string

	for _, dev := range developers {
		if _, exists := uniqueMap[dev.Name]; !exists {
			uniqueMap[dev.Name] = true
			uniqueNames = append(uniqueNames, dev.Name)
		}
	}
	return uniqueNames
}

func main() {
	fmt.Println("filter unique challenge")
}
