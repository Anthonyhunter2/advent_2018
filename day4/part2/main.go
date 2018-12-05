package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type guard struct {
	numOfSleeps  int
	minutesSlept []int
}

func main() {
	t := time.Now()
	events := make(map[int][]string)
	// building our map that holds our timestamp as the key
	// and either awake/asleep/gaurd #
	file, _ := os.Open("./input.txt")
	//	file, _ := os.Open("./sample.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	sortedKeys := []int{}
	for data.Scan() {
		line := data.Text()
		splitLine := strings.Split(line, "] ")
		timeLayout := strings.Trim(splitLine[0], "[")
		layout := "2006-01-02 15:04"
		t, err := time.Parse(layout, timeLayout)
		if err != nil {
			fmt.Println(err)
		}
		ts := int(t.Unix())
		text := strings.Split(splitLine[1:][0], " ")
		events[ts] = text
		sortedKeys = append(sortedKeys, ts)
	}
	sort.Ints(sortedKeys)
	var start time.Time
	var gid string
	guardstats := make(map[string]*guard)
	for _, sts := range sortedKeys {
		switch events[sts][0] {
		case "Guard":
			gid = events[sts][1]
			_, ok := guardstats[gid]
			if !ok {
				guardstats[gid] = &guard{
					numOfSleeps:  0,
					minutesSlept: make([]int, 60),
				}
			}
		case "falls":
			start = time.Unix(int64(sts), 0)
		case "wakes":
			end := time.Unix(int64(sts), 0)
			endtime := end.Sub(start)
			// track the start minute so we know where to iterate out list from
			smin := start.Minute()
			for i := smin; i < int(endtime.Minutes())+smin; i++ {
				guardstats[gid].numOfSleeps++
				guardstats[gid].minutesSlept[i]++
			}

		}
	}
	minAsleep := 0
	inOfMin := 0
	pradictableSleeper := ""
	for dude, lists := range guardstats {
		for inMax, max := range lists.minutesSlept {
			if max > minAsleep {
				minAsleep = max
				pradictableSleeper = dude
				inOfMin = inMax
			}
		}
	}
	intpradictableSleeper, _ := strconv.Atoi(strings.TrimLeft(pradictableSleeper, "#"))
	answer2 := intpradictableSleeper * inOfMin
	fmt.Println("Guy: ", pradictableSleeper, "Minute: ", inOfMin, "Answer: ", answer2)
	fmt.Println("Total Time: ", time.Since(t))
	os.Exit(0)

	// Firgure out who slept the most
	// sleptmost := 0
	// lazyDude := ""
	// for guy, stats := range guardstats {
	// 	if stats.numOfSleeps > sleptmost {
	// 		sleptmost = stats.numOfSleeps
	// 		lazyDude = guy
	// 	}
	// }
	// // Here we will find out which minute dude slept the most
	// indexOfSleepiestMinute := 0
	// total := 0
	// for insleep, minSlept := range guardstats[lazyDude].minutesSlept {
	// 	if minSlept > total {
	// 		total = minSlept
	// 		indexOfSleepiestMinute = insleep
	// 	}
	// }
	// // since we never removed the # from the index well, do that now so we can turn it to an int
	// // so we can do math on it
	// intLazyDude, _ := strconv.Atoi(strings.TrimLeft(lazyDude, "#"))
	// answer := intLazyDude * indexOfSleepiestMinute
	// fmt.Println("Guy :", lazyDude, "Day :", indexOfSleepiestMinute, "Answer = ", answer)
	// fmt.Println("Total Time: ", time.Since(t))
}
