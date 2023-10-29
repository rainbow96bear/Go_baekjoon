package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin) 
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	N, _ := strconv.Atoi(Input(scanner))

	song := Input(scanner)

	answer := 0
	for i:=0 ; i<N ; i++ {
		if song[0:N-i] == song[i:] {
			answer++
		}
	}
	fmt.Fprint(writer, answer)
}

func Input(scanner *bufio.Scanner)string{
	scanner.Scan()

	input := scanner.Text()

	return input
}