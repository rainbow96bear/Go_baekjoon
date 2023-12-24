package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var G, P int
	fmt.Fscanf(reader, "%d\n", &G)
	fmt.Fscanf(reader, "%d\n", &P)

	PlaneList := make([]int, P)
	for i := 0; i < P; i++ {
		fmt.Fscanf(reader, "%d\n", &PlaneList[i])
	}

	result := FindMaxPlaneCount(G, P, PlaneList)
	fmt.Fprintln(writer, result)
}

func FindMaxPlaneCount(G, P int, PlaneList []int) int {
	parents := make([]int, G+1)
	for i := 0; i <= G; i++ {
		parents[i] = i
	}

	count := 0
	for i := 0; i < P; i++ {
		gi := find(parents, PlaneList[i])
		if gi == 0 {
			break
		}
		union(parents, gi, gi-1)
		count++
	}
	return count
}

func find(parents []int, x int) int {
	if parents[x] == x {
		return x
	}
	parents[x] = find(parents, parents[x])
	return parents[x]
}

func union(parents []int, a, b int) {
	a = find(parents, a)
	b = find(parents, b)

	if a != b {
		parents[a] = b
	}
}
