package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkList(mystring string) (twos int, threes int) {
	mymap := make(map[string]int)
	mylist := strings.Split(mystring, "")
	for _, character := range mylist {
		numberOfOccurances := strings.Count(mystring, character)
		if numberOfOccurances == 2 {
			mymap["Twos"] = 1
		} else if numberOfOccurances == 3 {
			mymap["Threes"] = 1
		}
	}
	return mymap["Twos"], mymap["Threes"]
}
func main() {
	totalTwos := 0
	totalThrees := 0
	file, _ := os.Open("./input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := data.Text()
		twos, threes := checkList(line)
		totalTwos = totalTwos + twos
		totalThrees = totalThrees + threes
	}
	fmt.Printf("Twos: %v, Threes: %v\n", totalTwos, totalThrees)
	fmt.Println(totalThrees * totalTwos)

}
