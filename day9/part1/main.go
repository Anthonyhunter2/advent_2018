package main

import (
	"container/ring"
	"fmt"
	"time"
)

var currentPlayers int

func whosPlaying(turn int, players int) int {
	turn = turn + 1
	if turn > players {
		turn = 1
	}
	return turn
}

// Start just helps us wrap our benchmarking
func Start() {
	t := time.Now()
	lastmar := 72058 * 100

	//creates our players
	numOfPlayers := 426
	players := map[int]int{}
	for x := 1; x <= numOfPlayers; x++ {
		players[x] = 0
	}
	// create our circle
	board := &ring.Ring{Value: 0}

	//player 1 will start
	playerPlaying := 1
	for mar := 1; mar <= lastmar; mar++ {
		if mar%23 == 0 {
			board = board.Move(-8)
			r := board.Unlink(1)
			players[playerPlaying] += mar + r.Value.(int)
			board = board.Next()

		} else {
			board = board.Next()

			//The way to add something to the current ring is to create a new ring with a length of 1
			// give it the value of the marble and add it to our current ring
			r := ring.New(1)
			r.Value = mar
			// fmt.Println(r.Value)
			board.Link(r)
			board = board.Next()

		}
		playerPlaying = whosPlaying(playerPlaying, numOfPlayers)
	}
	// board.Do(func(p interface{}) {
	// 	fmt.Printf("%v,", p.(int))
	// })
	//fmt.Println(players)
	highestNum := []int{0, 0}
	for k, v := range players {
		if v > highestNum[1] {
			highestNum[1] = v
			highestNum[0] = k
		}
	}
	fmt.Println(highestNum)
	fmt.Println(time.Since(t))
}

// func steve() {
// 	lastmar := 25

// 	//creates our players
// 	numOfPlayers := 9
// 	players := map[int]int{}
// 	for x := 1; x <= numOfPlayers; x++ {
// 		players[x] = 0
// 	}
// 	circle := &ring.Ring{Value: 0}
// 	//player 1 will start
// 	playerPlaying := 1
// 	for m := 1; m <= lastmar; m++ {
// 		if m%23 == 0 {
// 			circle = circle.Move(-8)
// 			r := circle.Unlink(1)
// 			players[playerPlaying] += m + r.Value.(int)
// 			circle = circle.Next()
// 		} else {
// 			circle = circle.Next()

// 			r := ring.New(1)
// 			r.Value = m

// 			circle.Link(r)
// 			circle = circle.Next()
// 		}
// 		playerPlaying = whosPlaying(playerPlaying, numOfPlayers)

// 	}
// 	circle.Do(func(p interface{}) {
// 		fmt.Printf("%v,", p.(int))
// 	})
// 	fmt.Println(players)
// }
func main() {
	Start()
}
