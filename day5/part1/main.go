package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"unicode"
)

func main() {
	t := time.Now()
	polymer := ""
	file, _ := os.Open("input.txt")
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		polymer = reader.Text()

	}
	newGuy := []rune{}
	for _, rletter := range polymer {
		if len(newGuy) == 0 {
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
	fmt.Println(len(string(newGuy)))
	endtime := time.Since(t)
	fmt.Println(endtime)

}
