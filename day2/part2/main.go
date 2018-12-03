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


func checkForCloseMatch(word string, mylist []string) {
	match := len(word)
	wordList := strings.Split(word, "")
	for _, words := range mylist {
		checkagainst := strings.Split(words, "")
		closeMatch := 0

		for index := range checkagainst {
			if wordList[index] == checkagainst[index] {
				closeMatch++
			}
		}
		if closeMatch == match-1 {
			fmt.Printf("%s was only 1 character off %s\n", word, words)
			for indexs := range wordList {
				if wordList[indexs] != checkagainst[indexs] {
					newword := wordList[:indexs+copy(wordList[indexs:], wordList[indexs+1:])]
					fmt.Println("Resulting in: ", strings.Join(newword, ""))
				}
			}

			os.Exit(0)
		}

	}
}
func main() {
	var itemList []string
	file, _ := os.Open("./input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := data.Text()
		itemList = append(itemList, line)
	}
	for _, word := range itemList {
		checkForCloseMatch(word, itemList)
	}

}
