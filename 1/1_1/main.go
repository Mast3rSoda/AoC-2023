package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

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

	var wg sync.WaitGroup
	var atomCounter atomic.Int32

	//for each line in file
	for i, line := range lines {
		wg.Add(1)

		//declare a goroutine, cause we can do this in parallel
		go func(i int, line string) {
			x := 0

			//find the first number
			for _, rune := range line {
				num := int(rune - '0')
				if 0 > num || 9 < num {
					continue
				}
				x = num * 10
				break
			}

			//invert the iterator and find the last number
			runes := []rune(line)
			for i := len(runes) - 1; i >= 0; i-- {

				num := int(runes[i] - '0')
				if 0 > num || 9 < num {
					continue
				}
				x += num
				break
			}
			fmt.Printf("found number %d in line %d\n", x, i+1)

			//a
			atomCounter.Add(int32(x))
			wg.Done()
		}(i, line)
	}

	wg.Wait()

	fmt.Printf("Total: %d", atomCounter.Load())

}
