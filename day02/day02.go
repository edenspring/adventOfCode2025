package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DayTwoSolution interface {
	ProcessInput(inputPath string) (string, error)
	GetBadIDs() []int
	GetTotal() int
}

type IDChecker struct {
	IDs string
	BadIDS []int
	Total int
}

func (id *IDChecker) ProcessInput(inputPath string) (string, error) {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		return "", err
	}
	ids := string(bytes)
	id.IDs = ids
	return ids, nil
}

func (id *IDChecker) GetBadIDs() []int {
	idSlice := strings.Split(id.IDs, ",")
	for _, idRange := range idSlice {
		rangeNums := strings.Split(idRange, "-")
		start, startErr := strconv.Atoi(rangeNums[0])
		if startErr != nil {
			msg := fmt.Sprintf("++++++ Encountered unexpected value, %s", rangeNums[0])
			fmt.Println(msg)
		}
		end, endErr := strconv.Atoi(rangeNums[1])
		if endErr != nil {
			msg := fmt.Sprintf("++++++ Encountered unexpected value, %s", rangeNums[1])
			fmt.Println(msg)
		}
		id.processIDs(start, end)
	}
	return id.BadIDS
}

func (id *IDChecker) GetBadIDsPartTwo() []int {
	idSlice := strings.Split(id.IDs, ",")
	for _, idRange := range idSlice {
		rangeNums := strings.Split(idRange, "-")
		start, startErr := strconv.Atoi(rangeNums[0])
		if startErr != nil {
			msg := fmt.Sprintf("++++++ Encountered unexpected value, %s", rangeNums[0])
			fmt.Println(msg)
		}
		end, endErr := strconv.Atoi(rangeNums[1])
		if endErr != nil {
			msg := fmt.Sprintf("++++++ Encountered unexpected value, %s", rangeNums[1])
			fmt.Println(msg)
		}
		id.processIDsPartTwo(start, end)
	}
	return id.BadIDS
}

func (id *IDChecker) processIDs(start int, end int) {
	fmt.Printf("Start: %d\n", start)
	fmt.Printf("End: %d\n", end)

	for start <= end {
		curr := strconv.Itoa(start)
		currLength := len(curr)
		if currLength % 2 != 0 {
			start++
			continue
		}
		begin := curr[:currLength/2]
		end := curr[currLength/2:]
		if begin == end {
			fmt.Printf("Found BadID: %d\n", start)
			id.BadIDS = append(id.BadIDS, start)
			id.Total = id.Total + start
		}
		start++
	}
	fmt.Println("****************************************")
}

func (id *IDChecker) processIDsPartTwo(start int, end int) {
	fmt.Printf("Start: %d\n", start)
	fmt.Printf("End: %d\n", end)

	for start <= end {
		curr := strconv.Itoa(start)
		currLength := len(curr)
		for i := 0; i < currLength/2; i++ {
			pattern := curr[:i+1]
			patternLength := len(pattern)
			if currLength % patternLength != 0 {
				continue;
			}
			found := true
			for j := i +1; j < currLength; {
				toCheck := curr[j:j+patternLength]
				if pattern != toCheck {
					found = false
				}
				j = j + patternLength
			}
			if found {
				fmt.Printf("Found BadID: %d\n", start)
				id.BadIDS = append(id.BadIDS, start)
				id.Total = id.Total + start
				break;
			}
		}
		start++
	}
	fmt.Println("****************************************")
}

func (id *IDChecker) GetTotal() int {
	total := 0
	for _, num := range id.BadIDS {
		total += num
	}
	return total
}

func NewIDChecker(input string) IDChecker {
	checker := IDChecker{ Total: 0 } 
	if input == "" {
		checker.ProcessInput("./day02/input.txt")
	} else {
		checker.IDs = input
	}
	return checker
}

func SolveDay02() int {
	checker := NewIDChecker("")
	badIds := checker.GetBadIDs()
	fmt.Println(fmt.Sprintf("Bad ids: %v", badIds))
	fmt.Println("&&&& Total: +", checker.Total)
	return checker.GetTotal()
}

func SolveDay02PartTwo() int {
checker := NewIDChecker("")
	badIds := checker.GetBadIDsPartTwo()
	fmt.Println(fmt.Sprintf("Bad ids: %v", badIds))
	fmt.Println("&&&& Total: +", checker.Total)
	return checker.GetTotal()
}