package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	input1 := strings.Split(Input(scanner), " ")
	N, _ := strconv.Atoi(input1[0])
	M, _ := strconv.Atoi(input1[1])

	input2 := strings.Split(Input(scanner), " ")
	Num, _ := strconv.Atoi(input2[0])

	list := make([]int, Num)
	for i:=0; i<Num ; i++ {
		list[i], _ = strconv.Atoi(input2[i+1])
	}

	whereVisit := make(map[int][]int)
	partyList := make(map[int][]int)
	for i:=0 ; i<M ; i++ {
		input3 := strings.Split(Input(scanner), " ")
		peopleNum, _ := strconv.Atoi(input3[0])
		for j:=0 ; j<peopleNum ; j++ {
			who, _ := strconv.Atoi(input3[j+1])
			whereVisit[who] = append(whereVisit[who], i)
			partyList[i] = append(partyList[i], who)
		}
	}
	answer := Process(list, whereVisit, partyList, N, M)
	fmt.Fprintln(writer, answer)
}

func Process(list []int, whereVisit, partyList map[int][]int, N, M int) int {
	answer := M
	isCheck := make(map[int]bool)
	cantLie := make(map[int]bool)
	for len(list) > 0 {
		now := list[0]
		list = list[1:]
		for i:=0 ; i<len(whereVisit[now]) ; i++ {
			party := whereVisit[now][i]
			if cantLie[party] == false {
				cantLie[party] = true
				answer--
			}
			for j:=0 ; j<len(partyList[party]) ; j++ {
				if isCheck[partyList[party][j]] == false {
					isCheck[partyList[party][j]] = true
					list = append(list, partyList[party][j])
				}
			}
		}
	}
	return answer
}

func Input(scanner *bufio.Scanner) string {
	scanner.Scan()

	return scanner.Text()
}