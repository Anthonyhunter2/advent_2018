package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// stole this from SOF: https://stackoverflow.com/questions/32081808/strip-all-whitespace-from-a-string
func removeWhitespace(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}
func main() {
	// building a full list of IDs so i can remove the id's as they conflict
	var IDS []string
	//building our xy coordinate grid
	grid := make(map[int]map[int]string)
	for x := 0; x < 1000; x++ {
		grid[x] = map[int]string{}
		for y := 0; y < 1000; y++ {
			grid[x][y] = ""
		}
	}
	file, _ := os.Open("./input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := removeWhitespace(data.Text())
		reqID := strings.Split(line, "@")[0]
		IDS = append(IDS, reqID)
		xOffSet, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[0], ",")[0])
		yOffSet, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[0], ",")[1])
		xRun, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[1], "x")[0])
		yRun, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[1], "x")[1])
		for addToX := 0; addToX < xRun; addToX++ {
			for addToY := 0; addToY < yRun; addToY++ {
				if len(grid[xOffSet+addToX][yOffSet+addToY]) == 0 {
					grid[xOffSet+addToX][yOffSet+addToY] = reqID
				} else {
					// Store from SOF: https://stackoverflow.com/questions/34070369/removing-a-string-from-a-slice-in-go
					// a = a[:i+copy(a[i:], a[i+1:])]
					for index := range IDS {
						if IDS[index] == grid[xOffSet+addToX][yOffSet+addToY] {
							IDS = IDS[:index+copy(IDS[index:], IDS[index+1:])]
							break
						}
					}
					for index := range IDS {
						if IDS[index] == reqID {
							IDS = IDS[:index+copy(IDS[index:], IDS[index+1:])]
							break
						}
					}

				}
			}
		}
	}
	fmt.Println("The one good elf was request:", IDS[0])

}
