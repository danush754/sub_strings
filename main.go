package main

import (
	"fmt"
	"strings"
)

func main() {

	var dictionary = []string{"below", "down", "go", "going", "letgo", "horn", "how", "howdy", "it", "i", "low", "own", "part", "partner", "sit"}

	var subStrings = []string{"go", "low", "art"}

	substrings(dictionary, subStrings)
}

func substrings(dictionary []string, substrings []string) {

	var subStringCount = make(map[string]int)

	for _, dictWords := range dictionary {
		for _, substring := range substrings {

			if strings.Contains(dictWords, substring) {
				subStringCount[substring] += 1
			} else {
				continue
			}
		}
	}

	fmt.Println(subStringCount)

}
