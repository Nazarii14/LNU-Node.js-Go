package main

import "fmt"

func areAnagrams(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	var str1Map = make(map[rune]int)
	var str2Map = make(map[rune]int)

	for _, char := range str1 {
		str1Map[char]++
	}

	for _, char := range str2 {
		str2Map[char]++
	}

	for key, value := range str1Map {
		if str2Map[key] != value {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(areAnagrams("listen", "silent"))
	fmt.Println(areAnagrams("hello", "world"))
}
