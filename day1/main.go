package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func readInputToList() []int {
	var workingList []int
	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Print("Counldn't read file")
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		line := reader.Text()
		lineNumber, _ := strconv.Atoi(line)
		workingList = append(workingList, lineNumber)
	}
	return workingList
}
func checkForDups(newNum int, alreadyHave *[]int) bool {
	// fmt.Printf("Already have = %v\n", alreadyHave)
	for _, values := range *alreadyHave {
		if newNum == values {
			fmt.Printf("Found %v in alreadyhave\n", newNum)
			return true
		}
	}
	return false
}
func main() {
	workingList := readInputToList()
	counter := 0
	interations := 1
	var alreadyseen []int
	t := time.Now()

	for index := 0; index <= len(workingList); index++ {
		if index == len(workingList) {
			interations = interations + 1
			index = 0
		}
		newNum := workingList[index]
		counter = counter + newNum
		mycheck := checkForDups(counter, &alreadyseen)
		if !mycheck {
			alreadyseen = append(alreadyseen, counter)
			// interations = interations + 1
		} else {
			index = len(workingList) + 2
		}
	}
	fmt.Println(time.Since(t))
	fmt.Printf("%v\n", interations)
}
