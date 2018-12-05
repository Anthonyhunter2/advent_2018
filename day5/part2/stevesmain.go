package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
	"unicode"
)

var wg sync.WaitGroup
var mx sync.Mutex
var final []int

func work(polymer string, letter rune) {
	newGuy := []rune{}

	for _, rletter := range polymer {
		if unicode.ToLower(rletter) == letter {
			continue
		} else if len(newGuy) == 0 {
			newGuy = append(newGuy, rletter)
		} else {
			compareLetter := newGuy[len(newGuy)-1]
			if unicode.ToLower(rletter) == unicode.ToLower(compareLetter) {
				if rletter == unicode.ToLower(rletter) {
					if compareLetter == unicode.ToUpper(compareLetter) {
						newGuy = newGuy[:len(newGuy)-1]
					} else {
						newGuy = append(newGuy, rletter)
					}

				} else if rletter == unicode.ToUpper(rletter) {
					if compareLetter == unicode.ToLower(compareLetter) {
						newGuy = newGuy[:len(newGuy)-1]
					} else {
						newGuy = append(newGuy, rletter)
					}
				}
			} else {
				newGuy = append(newGuy, rletter)
			}
		}
	}
	mx.Lock()
	final = append(final, len(newGuy))
	mx.Unlock()
	wg.Done()
}

func Start() {
	t := time.Now()
	polymer := ""
	file, _ := os.Open("input2.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		polymer = reader.Text()
	}
	const useCases = "abcdefghijklmnopqrstuvwxyz"

	for _, letter := range useCases {
		wg.Add(1)
		go func(s string, b rune) {
			work(s, b)
		}(polymer, letter)
	}
	wg.Wait()
	min := 10000
	for _, jawn := range final {
		if min > jawn {
			min = jawn
		}
	}
	fmt.Println(min)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
