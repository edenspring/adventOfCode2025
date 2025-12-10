package main

import (
	"adventOfCode2025/day01"
	"adventOfCode2025/day02"
	"fmt"
	"strconv"
)

func dayStringBuider(day string) string {
	return fmt.Sprintf("******** DAY %s ********\n**                    **", day)
}

func dayStringBuilderPtTwo(day string) string {
	return fmt.Sprintf("****** DAY %s PT 2 ******\n**                    **", day)
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
		if diff%2 != 0 {
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

func main() {

	// Day 01
	day01String := dayStringBuider("01")
	day01Solve := day01.SolveDay01()
	day01SolveStr := strconv.Itoa(day01Solve)
	day01Output := solutionStringBuilder(day01SolveStr)
	fmt.Println(day01String)
	day01StringPartTwo := dayStringBuilderPtTwo("01")
	fmt.Println(day01Output)
	day01SolvePartTwo := day01.SolveDay01PartTwo()
	day01SolvePartTwoStr := strconv.Itoa(day01SolvePartTwo)
	day01PartTwoOutput := solutionStringBuilder(day01SolvePartTwoStr)
	fmt.Println(day01StringPartTwo)
	fmt.Println(day01PartTwoOutput)

	// Day 02
	day02String := dayStringBuider("02")
	day02Solve := day02.SolveDay02()
	day02SolveStr := strconv.Itoa(day02Solve)
	day02Output := solutionStringBuilder(day02SolveStr)
	fmt.Println(day02String)
	fmt.Println(day02Output)
	// Day 02 -- Part Two
	day02StringPartTwo := dayStringBuilderPtTwo("02")
	day02SolvePartTwo := day02.SolveDay02PartTwo()
	day02SolvePartTwoStr := strconv.Itoa(day02SolvePartTwo)
	day02PartTwoOutput := solutionStringBuilder(day02SolvePartTwoStr)
	fmt.Println(day02StringPartTwo)
	fmt.Println(day02PartTwoOutput)


}
