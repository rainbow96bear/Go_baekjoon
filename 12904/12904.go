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

	var S, T string
	fmt.Fscanf(reader, "%s\n", &S)
	fmt.Fscanf(reader, "%s\n", &T)
	for len(T) > len(S) {
		if T[len(T)-1] == 'A' {
			T = T[:len(T)-1]
		}else {
			T = T[:len(T)-1]
			T = reverse(T)
		}
	}
	if T == S {
		fmt.Fprintln(writer, 1)
	}else {
		fmt.Fprintln(writer, 0)
	}
}

func reverse(T string) string {
	result := ""
	for _, v := range T {
		result = string(v)+result 
	}
	return result
}