package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func existsInList(s string, l []string) bool {
	for _, check := range l {
		if s == check {
			return true
		}
	}
	return false
}

// Start is a main wrapper so we can do benchmarks
func Start() {
	t := time.Now()
	sets := make(map[string][]string)
	begins := []string{}
	finishs := []string{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := strings.Split(data.Text(), " ")
		begin := line[7]
		finish := line[1]
		begins = append(begins, begin)
		finishs = append(finishs, finish)
		sets[begin] = append(sets[begin], finish)
	}
	checkLater := []string{}
	for _, each := range finishs {
		if !existsInList(each, begins) {
			if !existsInList(each, checkLater) {
				checkLater = append(checkLater, each)
			}
		}
	}
	finalString := []string{}
	finalLetter := ""
	for len(sets) > 0 {
		sort.Strings(checkLater)
		removeChar := checkLater[0]
		fmt.Println(checkLater)
		checkLater = checkLater[1:]
		finalString = append(finalString, removeChar)
		if len(sets) == 1 {
			for last := range sets {
				finalLetter = last
			}
		}
		for char, every := range sets {
			badin := strings.Index(strings.Join(every, ""), removeChar)
			if badin != -1 {
				// a = append(a[:i], a[i+1:]...)
				// // or
				// a = a[:i+copy(a[i:], a[i+1:])]
				every = every[:badin+copy(every[badin:], every[badin+1:])]
				sets[char] = every

			}
			if len(every) == 0 {
				checkLater = append(checkLater, char)
				delete(sets, char)
			}
		}
	}
	fmt.Println(strings.Join(finalString, "") + finalLetter)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
