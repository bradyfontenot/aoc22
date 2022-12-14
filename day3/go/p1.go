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
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		sack := sc.Text()

		bag := make(map[rune]rune)
		for i, item := range sack {
			if i <= (len(sack)/2 - 1) {

				bag[item] = item
			} else if _, ok := bag[item]; ok {
				sum += priority[item]
				break
			}

		}

	}

	fmt.Println(sum)
}
