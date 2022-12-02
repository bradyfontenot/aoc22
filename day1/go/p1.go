package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	current     int = 1
	elves       map[int]int
	maxCalories int
	maxElf      int
)

func main() {

	filePath := os.Args[1]
	elves = make(map[int]int)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		calories := sc.Text()

		if calories == "" {
			current += 1
			continue
		}

		c, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}

		elves[current] += c

		if maxCalories < elves[current] {
			maxCalories = elves[current]
			maxElf = current
		}
	}

	fmt.Printf("Elf: \t\t %d\n", maxElf)
	fmt.Printf("Max Calories: \t %d\n", maxCalories)

}
