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
func secondsWwaiting(s string) int {
	i := []rune(s)
	return int(i[0]) - 4
}

type workers struct {
	WorkingON string
	worker    int
	TTC       int
	idle      bool
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
	ourWorker := map[int]*workers{}

	//For example
	for x := 1; x <= 5; x++ {
		ourWorker[x] = &workers{idle: true}
	}
	counter := 0
	for len(sets) > 0 {
		// for z := 0; z < 20; z++ {
		for x := 1; x <= 5; x++ {
			// fmt.Println(x)
			// fmt.Println(checkLater)
			sort.Strings(checkLater)
			if ourWorker[x].idle == true && (len(checkLater) > 0) {
				ourWorker[x].WorkingON = checkLater[0]
				ourWorker[x].TTC = secondsWwaiting(ourWorker[x].WorkingON)
				ourWorker[x].worker = x
				ourWorker[x].idle = false
				fmt.Println("USEING worker: ", x)
				fmt.Println(checkLater)
				checkLater = checkLater[1:]
				// finalString = append(finalString, ourWorker[x].WorkingON)
			} else if len(sets) == 1 && ourWorker[x].idle == true {
				for k := range sets {
					ourWorker[x].WorkingON = k
					ourWorker[x].TTC = secondsWwaiting(ourWorker[x].WorkingON)
					ourWorker[x].worker = x
					ourWorker[x].idle = false
					sets = nil

				}
			} else if ourWorker[x].TTC == 0 && ourWorker[x].idle == false {
				// if len(sets) == 1 {
				// for last := range sets {
				// 	finalLetter = last
				// 	fmt.Println(counter)
				// 	counter = counter + secondsWwaiting(last)
				// 	break
				// }
				for char, every := range sets {
					for in, str := range every {
						if str == ourWorker[x].WorkingON {
							sets[char] = append(sets[char][:in], sets[char][in+1:]...)

							// }
							// }
						}
					}
				}
				for char, every := range sets {
					if len(every) == 0 {
						checkLater = append(checkLater, char)
						delete(sets, char)

					}
					// }
				}
				finalString = append(finalString, ourWorker[x].WorkingON)
				ourWorker[x].WorkingON = ""
				ourWorker[x].idle = true
				counter++
				break

			}
			fmt.Println(ourWorker[x])
			ourWorker[x].TTC--
		}
		counter++

	}

	fmt.Println(counter + 1)
	fmt.Println(strings.Join(finalString, "") + finalLetter)
	fmt.Println(time.Since(t))
}
func main() {
	Start()
}
