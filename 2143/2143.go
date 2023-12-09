package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T, N, M int
	fmt.Fscan(reader, &T)
	fmt.Fscan(reader, &N)

	Aarr := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &Aarr[i])
	}

	fmt.Fscan(reader, &M)
	Barr := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Fscan(reader, &Barr[i])
	}

	Asum := make([]int, 0)
	Bsum := make([]int, 0)

	// 부분합 계산
	for i := 0; i < N; i++ {
		s := Aarr[i]
		Asum = append(Asum, s)
		for j := i + 1; j < N; j++ {
			s += Aarr[j]
			Asum = append(Asum, s)
		}
	}

	for i := 0; i < M; i++ {
		s := Barr[i]
		Bsum = append(Bsum, s)
		for j := i + 1; j < M; j++ {
			s += Barr[j]
			Bsum = append(Bsum, s)
		}
	}

	sort.Ints(Asum)
	sort.Ints(Bsum)

	answer := 0

	for i := 0; i < len(Asum); i++ {
		l := sort.Search(len(Bsum), func(j int) bool { return Bsum[j] >= T-Asum[i] })
		r := sort.Search(len(Bsum), func(j int) bool { return Bsum[j] > T-Asum[i] })
		answer += r - l
	}

	fmt.Fprintln(writer, answer)
}
