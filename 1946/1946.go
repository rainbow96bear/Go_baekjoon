package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Applicant struct {
	Document int
	Interview int
}

func main(){
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var T int
	fmt.Fscanf(reader, "%d\n", &T)

	for t:=0 ; t<T ; t++ {
		var N int 
		fmt.Fscanf(reader, "%d\n", &N)
		applicants := make([]Applicant, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(reader, &applicants[i].Document, &applicants[i].Interview)
		}
		fmt.Fscanf(reader, "\n")

		sort.Slice(applicants, func(i, j int) bool {
			return applicants[i].Document < applicants[j].Document
		})

		selected := 1
		interviewRank := applicants[0].Interview

		for i := 1; i < N; i++ {
			if applicants[i].Interview < interviewRank {
				selected++
				interviewRank = applicants[i].Interview
			}
		}
		fmt.Fprintln(writer, selected)
	}
}