package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	to int
	length int
}
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	tempInput := []int{}

	tempInput = Input(scanner)

	N, M, X := tempInput[0], tempInput[1], tempInput[2]
	RouteList := make(map[int][]Route)
	memberList := []int{}
	check := make(map[int]bool)
	for i:=0 ; i<M ; i++ {
		tempRoute := Input(scanner)
		from := tempRoute[0]
		to := tempRoute[1]
		length := tempRoute[2]

		RouteList[from] = append(RouteList[from], Route{to, length})
		if check[from] == false {
			memberList = append(memberList, from)
			check[from] = true
		}
	}
	maxLength := 0
	for i:=0 ; i<N ; i++ {
		from := memberList[i]
		to := X
		if from != to {	
			goLength :=FindRouteLength(from, to, RouteList, memberList)
			backLength := FindRouteLength(to, from, RouteList, memberList)
			total := goLength + backLength
			if maxLength < total {
				maxLength = total
			}
		}
	}
	fmt.Fprintln(writer, maxLength)
}

func FindRouteLength(from, to int, RouteList map[int][]Route, memberList []int) int {
	list := []int{from}
	dp := make(map[int]int)
	for _, v := range memberList {
		dp[v] = 1<<30
	}
	dp[from] = 0
	for len(list) > 0 {
		now := list[0]
		list = list[1:]
		for _, next := range RouteList[now] {
			if dp[next.to] > dp[now] + next.length {
				dp[next.to] = dp[now] + next.length
				list = append(list, next.to)
			}
		}
	}
	return dp[to]
}

func Input(scanner *bufio.Scanner) []int {
	scanner.Scan()

	list := strings.Split(scanner.Text(), " ")

	result := []int{}

	for _, v := range list {
		num, _  := strconv.Atoi(v)
		result = append(result, num)
	}
	
	return result
}