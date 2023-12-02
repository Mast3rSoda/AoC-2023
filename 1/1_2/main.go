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

	constNumberNames := map[int][]rune{1: []rune("one"), 2: []rune("two"), 3: []rune("three"), 4: []rune("four"), 5: []rune("five"), 6: []rune("six"), 7: []rune("seven"), 8: []rune("eight"), 9: []rune("nine")}
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
			exitFlag := false
			runes := []rune(line)

			//for each line we copy the key (number) with its length as the value
			numLenMap := make(map[int]int)
			for k := range constNumberNames {
				numLenMap[k] = 0
			}

			//find the first number
			for _, r := range runes {
				num := int(r - '0')
				//if we find a number we break out
				if 0 < num && 9 >= num {
					x = num * 10
					break
				}
				//else we search for the name representation of the number
				for k, v := range constNumberNames {
					if r == v[numLenMap[k]] {
						numLenMap[k]++
						if numLenMap[k] == len(v) {
							x = k * 10
							exitFlag = true
							break
						}
						//case for the same letter in a row
					} else if r != v[numLenMap[0]] {
						numLenMap[k] = 0
					}
				}
				if exitFlag {
					break
				}
			}

			//set the map value to length of the word
			for k := range numLenMap {
				numLenMap[k] = len(constNumberNames[k])
			}
			exitFlag = false

			//invert the iterator and find the last number
			for i := len(runes) - 1; i >= 0; i-- {

				num := int(runes[i] - '0')
				if 0 < num && 9 >= num {
					x += num
					break
				}

				//we go backwards
				for k, v := range constNumberNames {
					if runes[i] == v[numLenMap[k]-1] {
						numLenMap[k]--
						if numLenMap[k] == 0 {
							x += k
							exitFlag = true
							break
						}
						//same thing as in the previous loop, in case of repeating first letters
					} else if runes[i] != v[len(v)-1] {
						numLenMap[k] = len(v)
					}

				}
				if exitFlag {
					break
				}
			}
			fmt.Printf("found number %d in line %d\n", x, i+1)

			//atomic add since we use parallelism
			atomCounter.Add(int32(x))
			wg.Done()
		}(i, line)
	}

	wg.Wait()

	fmt.Printf("Total: %d", atomCounter.Load())

}
