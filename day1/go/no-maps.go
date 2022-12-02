package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var (
	elves     []int
	topThree  [3]int
	totesCals int
)

type elf struct {
	index    int
	calories int
}

func main() {

	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		calories := sc.Text()

		if calories == "" {
			elves = append(elves, totesCals)
			totesCals = 0
			continue
		}

		c, err := strconv.Atoi(calories)
		if err != nil {
			log.Fatal(err)
		}

		totesCals += c

		if totesCals > topThree[0] {
			topThree[0] = totesCals
		}
		sort.Ints(topThree[:])
	}

	var sum int
	for i := range topThree {
		sum += topThree[i]
	}

	fmt.Printf("Total Calories: %d\n", sum)
}
