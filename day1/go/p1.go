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
	elf         map[int]int
	maxCalories int
	maxElf      int
)

func main() {

	filePath := os.Args[1]
	elf = make(map[int]int)

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

		elf[current] += c

		if maxCalories < elf[current] {
			maxCalories = elf[current]
			maxElf = current
		}
	}

	fmt.Printf("Elf: \t\t %d\n", maxElf)
	fmt.Printf("Max Calories: \t %d\n", maxCalories)

}
