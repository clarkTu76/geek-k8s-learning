package main

import "fmt"

func main() {

	s := []string{"i", "am", "stupid", "and", "week"}
	fmt.Print(s)
	s1 := []string{"i", "am", "smart", "and", "strong"}

	for index := range s {
		s[index] = s1[index]
	}
	fmt.Print(s)

}
