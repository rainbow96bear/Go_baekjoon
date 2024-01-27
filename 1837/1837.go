package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()
	
	var p, k string
	fmt.Fscanf(reader , "%s %s\n", &p, &k)

	pBigInt := new(big.Int)
	kBigInt := new(big.Int)

	pBigInt.SetString(p, 10)
	kBigInt.SetString(k, 10)

	smallestPrime := findSmallestPrime(pBigInt, kBigInt)

	if smallestPrime != nil {
		fmt.Println("BAD", smallestPrime)
	} else {
		fmt.Println("GOOD")
	}
}

// 두 소수의 곱으로 이루어진 수에서 특정한 조건을 만족하는 가장 작은 소수 찾기
func findSmallestPrime(p, k *big.Int) *big.Int {
	one := big.NewInt(1)
	zero := big.NewInt(0)

	for i := big.NewInt(2); i.Cmp(k) < 0; i.Add(i, one) {
			if new(big.Int).Mod(p, i).Cmp(zero) == 0 {
				return i
			}
	}

	return nil
}