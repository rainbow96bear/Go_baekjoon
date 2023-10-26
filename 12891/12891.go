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

	var N, M int
	fmt.Fscan(reader, &N, &M)

	var DNA string
	fmt.Fscan(reader, &DNA)

	var A, C, G, T int
	fmt.Fscan(reader, &A, &C, &G, &T)

	password := DNA[0:M]
	check := []int{0, 0, 0, 0}
	answer := 0
	for i:=0 ; i<M ; i++ {
		if password[i] == 'A' {
			check[0]++
		}else if password[i] == 'C' {
			check[1]++
		}else if password[i] == 'G' {
			check[2]++
		}else if password[i] == 'T' {
			check[3]++
		}
	}
	if isPossible(check, A, C, G, T) {
		answer++ 
	}
	for i:=M ; i<N ; i++ {
		if DNA[i-M] == 'A' {
			check[0]--
		}else if DNA[i-M] == 'C' {
			check[1]--
		}else if DNA[i-M] == 'G' {
			check[2]--
		}else if DNA[i-M] == 'T' {
			check[3]--
		}
		if DNA[i] == 'A' {
			check[0]++
		}else if DNA[i] == 'C' {
			check[1]++
		}else if DNA[i] == 'G' {
			check[2]++
		}else if DNA[i] == 'T' {
			check[3]++
		}
		if isPossible(check, A, C, G, T) {
			answer++ 
		}
	}
	fmt.Fprintln(writer, answer)
}

func isPossible(check []int, A, C, G, T int) bool {
	if check[0] < A {
		return false
	}
	if check[1] < C {
		return false
	}
	if check[2] < G {
		return false
	}
	if check[3] < T {
		return false
	}
	return true
}