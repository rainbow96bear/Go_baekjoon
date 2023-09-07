package main

import (
	"bufio"
	"fmt"
	"os"
)

func sieveOfEratosthenes(n int) []bool {
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i*i <= n; i++ {
		if isPrime[i] {
			for j := i * i; j <= n; j += i {
				isPrime[j] = false
			}
		}
	}

	return isPrime
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	const maxN = 1000000
	isPrime := sieveOfEratosthenes(maxN)

	var T int
	fmt.Fscan(reader, &T)

	for i := 0; i < T; i++ {
		var n int
		fmt.Fscan(reader, &n)

		count := 0
		for j := 2; j <= n/2; j++ {
			if isPrime[j] && isPrime[n-j] {
				count++
			}
		}

		fmt.Fprintln(writer, count)
	}
}
