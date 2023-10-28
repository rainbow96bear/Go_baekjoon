package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	time int
	value int
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input := strings.Split(Scan(scanner), " ")
	N, _ := strconv.Atoi(input[0])
	K, _ := strconv.Atoi(input[1])

	answer := Process(N, K)

	fmt.Fprintln(writer, answer)
}

func Process(N, K int) int {
	if N == K {
		return 0
	}
    list := []Node{{0, N}}
	isVisit := make(map[int]bool)

	for len(list) > 0{
		now := list[0]
        list = list[1:]
		t := now.time
		v := now.value
		if v == K {
			return t
		}
		if isVisit[v*2] == false && v*2 <= 100000{
			list = append(list, Node{t, v*2})
			isVisit[v*2] = true
		}
		if v >= 0 && isVisit[v-1] == false {
			list = append(list, Node{t+1, v-1})
			isVisit[v-1] = true
		}
		if isVisit[v+1] == false && v+1 <= 100000{
			list = append(list, Node{t+1, v+1})
			isVisit[v+1] = true
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Scan(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := scanner.Text()
	return input
}