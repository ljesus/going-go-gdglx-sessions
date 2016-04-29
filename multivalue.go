package main

import (
	"errors"
	"fmt"
)

func nexusPerGoogler(googlers, nexus int) (float32, error) {
	if googlers == 0 {
		return 0, errors.New("No Googlers! Let's not divide by zero")
	}

	return float32(nexus) / float32(googlers), nil
}

func main() {
	var googlers int = 50
	nexus := 40

	npg, _ := nexusPerGoogler(googlers, nexus)
	fmt.Println("Nexus per Googler:", npg)
}
