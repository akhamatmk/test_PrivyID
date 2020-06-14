package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "{a"

	// make regex just allow input [ ] { } and minimum 1 input
	re := regexp.MustCompile(`[\[\]\{\}]+${1}`)
	//get value match string 1
	debug := re.MatchString(str1)

	if !debug {
		// if other char in str1 will be print nil
		fmt.Println(nil)
	} else {
		// match string and just allow 2 char
		match, _ := regexp.MatchString(`\{}|\[]{2,2}`, str1)

		// print true if input [] or {}
		fmt.Println(match)
	}

}
