package main

import (
	"adventOfCode2025/day01"
	"fmt"
	"strconv"
)

func dayStringBuider(day string) string {
	return fmt.Sprintf("******** DAY %s ********\n**                    **", day)
}

func solutionStringBuilder(solution string) string {
	space := 20
	bookend := "++" 
	var output string
	if len(solution) > space {
		output = bookend + "\n" + solution + "\n" + bookend
	} else {
		diff := space - len(solution)
		buffL := diff / 2
		buffR := buffL
		if diff % 2 != 0 {
			buffR++
		} 	
		output = bookend
		for i := 0; i < buffL; i++ {
			output = output + " "
		}
		output = output + solution
		for i := 0; i < buffR; i++ {
			output = output + " "
		}
		output = output + bookend
	}
	return output + "\n**                    **"
}

func main(){
	day01String := dayStringBuider("01")
	day01Solve := day01.SolveDay01()
	day01SolveStr := strconv.Itoa(day01Solve)
	day01Output := solutionStringBuilder(day01SolveStr)
	fmt.Println(day01String)
	fmt.Println(day01Output)
}	