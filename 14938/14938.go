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

	list := strings.Split(Input(scanner), " ")
	n, _ := strconv.Atoi(list[0])
	m, _ := strconv.Atoi(list[1])
	r, _ := strconv.Atoi(list[2])

	itemList := make(map[int]int, n)
	
	tempItemList := strings.Split(Input(scanner), " ")
	for i, v := range tempItemList {
		itemList[i+1], _ = strconv.Atoi(v)
	}

	RouteList := make(map[int][]Route, n)
	
	for i:=0 ; i<r ; i++ {
		tempRoute := strings.Split(Input(scanner), " ")
		a, _ := strconv.Atoi(tempRoute[0])
		b, _ := strconv.Atoi(tempRoute[1])
		length, _ := strconv.Atoi(tempRoute[2])
		RouteList[a] = append(RouteList[a], Route{b, length})
		RouteList[b] = append(RouteList[b], Route{a, length})
	}
	max := 0
	for i:=1 ; i<=n ; i++ {
		result := Process(itemList, RouteList, n, m, i)
		if max < result {
			max = result
		}
	}
	fmt.Fprintln(writer, max)
}

func Process(itemList map[int]int, RouteList map[int][]Route, n, m, start int) int {
	dp := make(map[int]int, n)
	for i:=0 ; i<=n ; i++ {
		dp[i] = 1<<30
	}
	areaList := []int{start}
	count := 0
	dp[start] = 0
	for len(areaList) > 0 {
		now := areaList[0]
		areaList = areaList[1:]

		for _, v := range RouteList[now] {
			nextArea := v.to
			Length := v.length
			if dp[nextArea] > dp[now] + Length{
				dp[nextArea] = dp[now] + Length
				areaList = append(areaList, nextArea)
			}
		}
	}
	for i:=1 ; i<=n ; i++ {
		if dp[i] <= m {
			count += itemList[i]
		}
	}
	return count
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}