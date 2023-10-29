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

	input := Input(scanner)

	answer := Process(N, input)

	fmt.Fprintln(writer, answer)
}
func Process(N int, input string) string {
	checkSmall := make([]bool, 7)
	checkBig := make([]bool, 7)
	for i:=0 ; i<N ; i++ {
		switch input[i] {
			case 'r' :
				checkSmall[0] = true
			case 'R' :
				checkBig[0] = true
			case 'o' :
				checkSmall[1] = true
			case 'O' :
				checkBig[1] = true
			case 'y' :
				checkSmall[2] = true
			case 'Y' :
				checkBig[2] = true
			case 'g' :
				checkSmall[3] = true
			case 'G' :
				checkBig[3] = true
			case 'b' :
				checkSmall[4] = true
			case 'B' :
				checkBig[4] = true
			case 'i' :
				checkSmall[5] = true
			case 'I' :
				checkBig[5] = true
			case 'v' :
				checkSmall[6] = true
			case 'V' :
				checkBig[6] = true
		}
	}
	isSmallPossible := true
	isBigPossible := true
	for i:=0 ; i<7 ; i++ {
		if checkSmall[i] == false {
			isSmallPossible = false
		}
		if checkBig[i] == false {
			isBigPossible = false
		}
	}
	if isSmallPossible == true && isBigPossible == true {
		return "YeS"
	}
	if isSmallPossible == true {
		return "yes"
	}
	if isBigPossible == true {
		return "YES"
	}
	return "NO!"
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	input := scanner.Text()

	return input
}