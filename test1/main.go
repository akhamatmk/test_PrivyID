package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	st1, st2 := getStringLongsolution(19374048, "aiueobcd")
	fmt.Println(st1)
	fmt.Println(st2)

	st3, st4 := getStringFastSolution(19374048, "aiueobcd")
	fmt.Println(st3)
	fmt.Println(st4)
}

//
func getStringLongsolution(v1 int, v2 string) (string, string) {
	// convert conversi int to string
	s := strconv.Itoa(v1)

	// make string to slice so can be loop for search value
	slice1 := strings.Split(s, "")
	slice2 := strings.Split(v2, "")

	// initiation index to search data
	index := -1

	// looping slice params 1
	for i, v := range slice1 {

		// if data search equal next break loop
		if v == "7" {
			index = i
			break
		}
	}

	// check if index still get value -1 data seacrh in loop not found
	if index == -1 {
		return "", ""
	}

	// get len data slice 1 for initate max lopping
	maxSlice1 := len(slice1)

	// iniate string 1
	finalString1 := ""

	// looping for get string 1
	for i := index; i < maxSlice1; i++ {
		finalString1 += slice1[i]
	}

	// get len data slice 2 for initate max lopping
	maxSlice2 := len(slice2)

	// initiate string 2
	finalString2 := ""

	// looping for get string 1
	for i := index; i < maxSlice2; i++ {
		finalString2 += slice2[i]
	}

	// return result string 1 and string 2
	return finalString1, finalString2
}

func getStringFastSolution(v1 int, v2 string) (string, string) {
	// convert conversi int to string
	s := strconv.Itoa(v1)
	// search index contain search
	i := strings.Index(s, "7")

	// if index lower than 0 return empty string
	if i < 0 {
		return "", ""
	}

	// return result string 1 and string 2
	return s[i:], v2[i:]
}
