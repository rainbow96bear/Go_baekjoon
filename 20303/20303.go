package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M, K int
	fmt.Fscanf(reader, "%d %d %d\n", &N, &M, &K)

	candyList := make([]int, N+1)
	for i:=0 ; i<N ; i++ {
		fmt.Fscanf(reader, "%d", &candyList[i+1])
	}
	fmt.Fscanf(reader, "\n")
	
	FriendList := make([][]int, N+1)
	for i:=0 ; i<M ; i++ {
		var a, b int
		fmt.Fscanf(reader, "%d %d\n", &a, &b)
		FriendList[a] = append(FriendList[a], b)
		FriendList[b] = append(FriendList[b], a)
	}
	result := FindMaxCandy(N, M, K, candyList, FriendList)
	fmt.Fprintln(writer, result)
}

func FindMaxCandy(N, M, K int, candyList []int, FriendList [][]int) int {
	// FriendList로 그룹을 만들고 그룹의 인원 수와 사탕의 개수 확인
	// 
	Group := makeGroup(N, candyList, FriendList)
	sort.Slice(Group, func(i, j int) bool {return Group[i][0] < Group[j][0]})
	dp := make([][]int, 2)
	dp[0] = make([]int,K)
	dp[1] = make([]int,K)
	for i:=0 ; i<len(Group) ; i++ {
		for j:=0 ; j<K ; j++ {
			if j>=Group[i][0] {
				dp[1][j] = max(dp[0][j], dp[0][j-Group[i][0]]+Group[i][1])
			}else {
				dp[1][j] = dp[0][j]
			}
		}
		for i:=0 ; i<K ; i++ {
			dp[0][i] = dp[1][i]
		}
	}
	return dp[1][K-1]
}

func makeGroup(N int, candyList []int, FriendList [][]int) [][]int {
	isVisit := make([]bool, N+1)
	Group := [][]int{}
	for i:=1 ; i<=N ; i++ {
		if isVisit[i] == false {
			isVisit[i] = true
			queue := []int{i}
			count := 1
			totalCandy := candyList[i]
			for len(queue) > 0 {
				now := queue[0]
				queue= queue[1:]

				for _, num := range FriendList[now] {
					if isVisit[num] == false {
						isVisit[num]=true
						count++
						totalCandy += candyList[num]
						queue = append(queue, num)
					}
				}
			}
			Group = append(Group, []int{count, totalCandy})
		}
	}
	return Group
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}