package main

import "fmt"

func main() {
	entries := []int{10, 20, 30, 40, 50}
	for _, val := range entries {
		fmt.Println(val)
	}

	scores := map[string]int{
		"Filipe": 0,
		"Gilberto": -1,
		"Elisabete":  12,
	}

	for name, score := range scores {
		fmt.Println(name, ":", score)
	}
}
