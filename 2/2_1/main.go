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

	bag := game{12, 13, 14}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	// var wg sync.WaitGroup
	var atomCounter atomic.Int32

	//for each line in file
	for i, line := range lines {
		// wg.Add(1)

		//declare a goroutine, cause we can do this in parallel
		func(i int, line string) {

			splits := strings.Split(line, ": ")
			gameNum, _ := strconv.Atoi(strings.Split(splits[0], " ")[1])
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
			if gg.maxRed <= bag.maxRed && gg.maxBlue <= bag.maxBlue && gg.maxGreen <= bag.maxGreen {
				atomCounter.Add(int32(gameNum))
			}
		}(i, line)
	}

	// wg.Wait()

	fmt.Printf("Total: %d", atomCounter.Load())

}
