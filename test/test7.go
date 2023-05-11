package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	oldStr := []string{"1", "2", "3", "2", "2", "1"}

	mapString := make(map[string]bool)
	for _, num := range oldStr {
		mapString[num] = true
	}

	newStr := maps.Keys[map[string]bool](mapString)

	slices.Sort[]()

	fmt.Printf("%#v", newStr)
}
