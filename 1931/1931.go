package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Meeting struct {
	start, end int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	var N int
	fmt.Fscan(reader, &N)

	meetings := make([]Meeting, N)
	for i:=0; i<N; i++ {
		fmt.Fscan(reader, &meetings[i].start, &meetings[i].end)
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i].end == meetings[j].end {
			return meetings[i].start < meetings[j].start
		}
		return meetings[i].end < meetings[j].end
	})

	count := 0
	endTime := 0

	for i:=0; i<N; i++ {
		if meetings[i].start >= endTime {
			endTime = meetings[i].end
			count++
		}
	}

	fmt.Fprint(writer, count)
}