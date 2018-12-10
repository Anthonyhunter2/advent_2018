package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type node struct {
	ind           int
	NumchildNodes int
	MetaNum       int
	childNodes    []node
	metalist      []int
}

// Start is just here so we can wrap tests if we want to
func Start() {
	t := time.Now()
	file, _ := os.Open("input.txt")
	defer file.Close()
	fullListAsInts := []int{}
	// tree := []node{}
	data := bufio.NewScanner(file)
	for data.Scan() {
		each := strings.Split(data.Text(), " ")
		for _, nums := range each {
			addme, _ := strconv.Atoi(string(nums))
			fullListAsInts = append(fullListAsInts, addme)
		}

	}

	rootnode := grabAllChildren(0, fullListAsInts)
	fmt.Println(sumMeta(rootnode))

	fmt.Println(time.Since(t))
}
func grabAllChildren(find int, flist []int) node {
	rootnode := node{ind: find, NumchildNodes: flist[find], MetaNum: flist[find+1]}
	offset := rootnode.ind + 2
	for c := 0; c < rootnode.NumchildNodes; c++ {
		childnode := grabAllChildren(offset, flist)
		rootnode.childNodes = append(rootnode.childNodes, childnode)
		offset = offset + getLength(childnode)

	}
	for mn := 0; mn < rootnode.MetaNum; mn++ {
		rootnode.metalist = append(rootnode.metalist, flist[offset+mn])
	}
	return rootnode
}

func sumMeta(Node node) int {
	sum := 0
	for _, v := range Node.childNodes {
		sum = sum + sumMeta(v)
	}
	for _, v := range Node.metalist {
		sum = sum + v
	}
	return sum
}

func getLength(Node node) int {
	length := 2
	for i := 0; i < Node.NumchildNodes; i++ {
		length = length + getLength(Node.childNodes[i])
	}
	length = length + Node.MetaNum
	return length
}
func main() {
	Start()
}
