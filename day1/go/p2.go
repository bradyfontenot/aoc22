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
	current     int = 1
	elf         map[int]int
	maxCalories int
	topThree    [3]elfie
)

type elfie struct {
	index    int
	calories int
}

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

		if elf[current] > topThree[2].calories {
			topThree[2].index = current
			topThree[2].calories = elf[current]
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
