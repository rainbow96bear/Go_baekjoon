package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanf(reader, "%d\n", &n)

	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n')
		list := strings.Split(input, " ")
		T, _ := strconv.Atoi(list[0])
		result := mostFrequentNumber(T, list[1:len(list)])
		fmt.Fprintln(writer, result)
	}
}

func mostFrequentNumber(T int, arr []string) string {
	counter := make(map[string]int)

	for _, num := range arr {
		counter[num]++
	}
	maxNum := ""
	for k, v := range counter {
		if v > T/2 {
			maxNum = k
		}
	}
	
	if maxNum != ""{
		return maxNum
	}
	return "SYJKGW"
}