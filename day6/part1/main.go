package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type coords struct {
	xv   float64
	yv   float64
	seen int
}

func amIHere(bad []int, check int) bool {
	for _, i := range bad {
		if i == check {
			return true
		}
	}
	return false
}

// Start is just here so we can wrap tests if we want to
func Start() {
	t := time.Now()
	minx := 500
	maxx := 0
	miny := 500
	maxy := 0
	mycoords := []coords{}
	file, _ := os.Open("input2.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := strings.TrimRight(strings.TrimLeft(data.Text(), "[ "), "] ")
		splitLine := strings.Split(line, ", ")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		fx := float64(x)
		fy := float64(y)

		addMe := coords{xv: fx, yv: fy, seen: 0}
		mycoords = append(mycoords, addMe)
		if minx > x {
			minx = x
		}
		if miny > y {
			miny = y
		}
		if maxx < x {
			maxx = x
		}
		if maxy < y {
			maxy = y
		}
	}
	infinate := []int{}

	for xrange := minx; xrange <= maxx; xrange++ {
		for yrange := miny; yrange <= maxy; yrange++ {
			closest := 5000
			shared := false
			index := 0
			for where, each := range mycoords {
				myx := each.xv
				myy := each.yv
				howFar := int((math.Abs(myx-float64(xrange)) + math.Abs(myy-float64(yrange))))
				if howFar < closest {
					closest = howFar
					index = where
					shared = false
					if (yrange == maxy) || (xrange == maxx) || (xrange == minx) || (yrange == miny) {
						if !amIHere(infinate, where) {
							infinate = append(infinate, where)
						}
					}

				} else if howFar == closest {
					shared = true
				}
			}
			if !shared {
				mycoords[index].seen++
			}
		}

	}
	for _, cheater := range infinate {
		mycoords[cheater].seen = 0
	}

	highestNum := 0
	for _, highest := range mycoords {
		if highest.seen > highestNum {
			highestNum = highest.seen
		}
	}
	fmt.Println(highestNum)
	fmt.Printf("MinX: %v, MaxX: %v, MinY: %v, MaxY %v\n", minx, maxx, miny, maxy)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
