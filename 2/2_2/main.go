package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)

type game struct {
	maxRed   int
	maxGreen int
	maxBlue  int
}

func main() {

	//open and read file by line
	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	// var wg sync.WaitGroup
	var atomCounter atomic.Int32

	//for each line in file
	for i, line := range lines {
		// wg.Add(1)

		func(i int, line string) {

			splits := strings.Split(line, ": ")
			gg := game{0, 0, 0}

			gameSplits := strings.Split(splits[1], "; ")
			for _, g := range gameSplits {
				cubes := strings.Split(g, ", ")
				for _, cs := range cubes {
					sp := strings.Split(cs, " ")
					switch sp[1][0] {
					case 'r':
						if a, _ := strconv.Atoi(sp[0]); a > gg.maxRed {
							gg.maxRed = a
						}
					case 'g':
						if a, _ := strconv.Atoi(sp[0]); a > gg.maxGreen {
							gg.maxGreen = a
						}
					case 'b':
						if a, _ := strconv.Atoi(sp[0]); a > gg.maxBlue {
							gg.maxBlue = a
						}
					}
				}
			}
			atomCounter.Add(int32(gg.maxRed * gg.maxBlue * gg.maxGreen))
		}(i, line)
	}

	// wg.Wait()

	fmt.Printf("Total: %d", atomCounter.Load())

}
