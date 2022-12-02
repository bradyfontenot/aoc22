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
	current  int = 1
	elves    map[int]int
	topThree [3]elf
)

type elf struct {
	index    int
	calories int
}

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

		if elves[current] > topThree[2].calories {
			topThree[2].index = current
			topThree[2].calories = elves[current]
		}
		sort.Slice(topThree[:], func(i int, j int) bool { return topThree[i].calories > topThree[j].calories })
	}

	var sum int
	for i := range topThree {
		fmt.Printf("Elf: %d \t\t Calories: %d\n", topThree[i].index, topThree[i].calories)
		sum += topThree[i].calories
	}
	fmt.Println("********************************************")
	fmt.Printf("Total: \t\t\t Calores: %d\n", sum)
}
