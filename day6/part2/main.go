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

// Start is just here so we can wrap tests if we want to
func Start() {
	t := time.Now()
	var minx = 500
	var maxx int
	var miny = 500
	var maxy int
	mycoords := []coords{}
	file, _ := os.Open("input.txt")
	defer file.Close()
	data := bufio.NewScanner(file)
	for data.Scan() {
		line := strings.TrimRight(strings.TrimLeft(data.Text(), "[ "), "] ")
		splitLine := strings.Split(line, ", ")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])

		addMe := coords{xv: float64(x), yv: float64(y), seen: 0}
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
	largetShared := 0
	for xrange := minx; xrange < maxx; xrange++ {
		for yrange := miny; yrange < maxy; yrange++ {
			sharedCounter := 0
			for _, each := range mycoords {
				myx := each.xv
				myy := each.yv
				howFar := int((math.Abs(myx-float64(xrange)) + math.Abs(myy-float64(yrange))))
				sharedCounter = sharedCounter + howFar
			}
			if sharedCounter < 10000 {
				largetShared++
			}
		}
	}
	fmt.Println(largetShared)
	fmt.Println(time.Since(t))
}

func main() {
	Start()
}
