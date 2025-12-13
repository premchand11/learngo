package stringpackage

import (
	"fmt"
	"strings"
)

//basically go lang is made beaces of stig manipulation like the main focus is on strings

func Learnstrings() {
	data := "appleorangebanana"
	parts := strings.Split(data, "")
	fmt.Println(parts)

	str := "one two theree four five one one one one"
	count := strings.Count(str, "one")
	fmt.Println(count)

	gaps := "     hello     ,go    "
	trimmed := strings.TrimSpace(gaps)
	fmt.Println(trimmed)

	str1 := "prem"
	str2 := "chand"
	result := strings.Join([]string{str1, str2}, ",")
	fmt.Println(result)

}
