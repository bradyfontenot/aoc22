package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum int
	var counter int
	var matches [][]string
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		pairs := strings.Split(line, ",")
		aS := strings.Split(pairs[0], "-")
		bS := strings.Split(pairs[1], "-")

		var a [2]int
		var b [2]int
		for i := 0; i < 2; i++ {
			a[i], _ = strconv.Atoi(aS[i])
			b[i], _ = strconv.Atoi(bS[i])
		}

		if a[0] < b[0] && a[1] < b[0] || b[0] < a[0] && b[1] < a[0] {
			sum += 1
			matches = append(matches, pairs)
		}
		counter += 1
	}

	fmt.Println(counter - sum)
}
