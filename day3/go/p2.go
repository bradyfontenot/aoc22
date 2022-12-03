package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var priority = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var sum int
	var count int
	var group []map[rune]rune

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		count += 1

		bag := make(map[rune]rune)
		group = append(group, bag)
		for _, item := range sc.Text() {

			bag[item] = item
			if count%3 == 0 {
				_, ok1 := group[0][item]
				_, ok2 := group[1][item]
				if ok1 && ok2 {
					sum += priority[item]
					group = nil
					break
				}
			}
		}
	}

	fmt.Println(sum)
}

// create separate map of first 2 sacks
// check if item in 3rd sack in other 2
// sum
// reset
