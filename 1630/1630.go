package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N int
	fmt.Fscanf(reader, "%d\n", &N)
	mod := 987654321
	prime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		for j := i * 2; j <= N; j += i {
			prime[j] = true
		}
	}

	var primes []int
	for i := 2; i <= N; i++ {
		if !prime[i] {
			primes = append(primes, i)
		}
	}

	ret := int64(1)
	for _, p := range primes {
		if p > N {
			break
		}
		k := int64(p)
		for k*int64(p) <= int64(N) {
			k *= int64(p)
		}
		ret = (ret * k) % int64(mod)
	}
	fmt.Fprintln(writer, ret)
}