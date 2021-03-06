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
	//building our xy coordinate grid
	grid := make(map[string][]int)

	file, _ := os.Open("./input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := removeWhitespace(data.Text())
		reqID := strings.Split(line, "@")[0]
		xOffSet, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[0], ",")[0])
		yOffSet, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[0], ",")[1])
		xRun, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[1], "x")[0])
		yRun, _ := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "@")[1], ":")[1], "x")[1])
		grid[reqID] = []int{xOffSet, yOffSet, xRun, yRun}
	}
	// 	for addToX := 0; addToX < xRun; addToX++ {
	// 		for addToY := 0; addToY < yRun; addToY++ {
	// 			grid[xOffSet+addToX][yOffSet+addToY]++
	// 		}
	// 	}
	// }

	fabric := make([][]uint8, 1000)
	for i := range fabric {
		fabric[i] = make([]uint8, 1000)
	}

	for _, v := range grid {
		et := v[0] + v[2]
		bl := v[1] + v[3]
		for i := v[0]; i < et; i++ {
			for j := v[1]; j < bl; j++ {
				fabric[i][j]++
			}
		}
	}
	multpileCuts := 0
	for xaxis := range grid {
		//	for yaxis, values := range grid[xaxis] {
		for _, values := range grid[xaxis] {
			if values > 1 {
				multpileCuts++
				//fmt.Printf("x: %v y: %v has more than one elf cuts\n", xaxis, yaxis)
			}
		}
	}
	fmt.Println("Squares with multiple cuts: ", multpileCuts)
}
