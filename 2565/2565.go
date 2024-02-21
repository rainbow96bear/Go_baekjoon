package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Line struct {
	A, B int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	lines := make([]Line, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %d\n", &lines[i].A, &lines[i].B)
	}

	sort.Slice(lines, func(i, j int) bool {
		return lines[i].A < lines[j].A
	})

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if lines[i].B > lines[j].B && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	result := n - max(dp...)
	fmt.Fprintln(writer, result)
}

func max(nums ...int) int {
	maximum := nums[0]
	for _, num := range nums[1:] {
		if num > maximum {
			maximum = num
		}
	}
	return maximum
}
