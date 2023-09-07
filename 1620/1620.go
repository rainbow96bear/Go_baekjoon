package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var N, M int
	fmt.Fscan(reader, &N, &M)
	PokemonListNum := make([]string,N)
	PokemonListStr := make(map[string]int,N)
	for i:=0 ; i<N ; i++{
		fmt.Fscan(reader, &PokemonListNum[i])
		PokemonListStr[PokemonListNum[i]]=i
	}
	for i:=0 ; i<M ; i++{
		var input string
		fmt.Fscan(reader,&input)
		index,err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintln(writer, PokemonListStr[input]+1)
		}else {
			fmt.Fprintln(writer, PokemonListNum[index-1])
		}
	}
}