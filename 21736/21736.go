package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)

	campusMap := make([][]string, N)
	for i:=0 ; i<N ; i++ {
		var input string
		fmt.Fscan(reader, &input)
		campusMap[i] = strings.Split(input, "")
	}

	x, y := FindStartingPoint(campusMap)
	answer := FindP(campusMap, x, y, N, M)
	if answer == 0 {
		fmt.Fprintln(writer, "TT")
		return
	}
	fmt.Fprintln(writer, answer)
}

func FindStartingPoint(campusMap [][]string) (x, y int) {
	for i:=0 ; i<len(campusMap) ; i++ {
		for j:=0 ; j<len(campusMap[i]) ; j++ {
			if campusMap[i][j] == "I" {
				return i, j
			}
		}
	}
	return 0, 0
}

func FindP(campusMap [][]string, x, y, maxX, maxY int) int {
	cnt := 0
	if campusMap[x][y] == "P" {
		cnt++
	}else if campusMap[x][y] == "X"{
		return 0
	}
	campusMap[x][y] = "X"

	if 0 <= x+1 && x+1 < maxX {
		cnt += FindP(campusMap, x+1, y, maxX, maxY)
	}
	if 0 <= x-1 && x-1 < maxX {
		cnt += FindP(campusMap, x-1, y, maxX, maxY)
	}
	if 0 <= y+1 && y+1 < maxY {
		cnt += FindP(campusMap, x, y+1, maxX, maxY)
	}
	if 0 <= y-1 && y-1 < maxY {
		cnt += FindP(campusMap, x, y-1, maxX, maxY)
	}

	return cnt
}