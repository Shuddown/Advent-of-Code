package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/Shuddown/Advent-of-Code/utils"
)

func performOp(operator rune, nums ...int) int {
	if operator == '*' {
		product := 1
		for _, num := range nums {
			product *= num
		}
		return product
	} else {
		sum := 0
		for _, num := range nums {
			sum += num
		}
		return sum
	}
}

func indexOf(row, col, numCols int) int {
	return row*(numCols+1) + col
}

func part2(s string) [][]int {
	numCols := strings.Index(s, "\n")
	numRows := len(s)/numCols - 1
	numString := strings.Builder{}
	numStrings := make([]string, 0)
	for i := range numCols {
		for j := range numRows {
			numString.WriteByte(s[indexOf(j, i, numCols)])
		}
		numStrings = append(numStrings, numString.String())
		numString.Reset()
	}
	problems := make([][]int, 0)
	problem := make([]int, 0)
	for _, numString := range numStrings {
		cleanedUpNumString := strings.TrimSpace(numString)
		if cleanedUpNumString == "" {
			problems = append(problems, problem)
			problem = nil
			continue
		}
		num, err := strconv.Atoi(cleanedUpNumString)
		utils.HandleError(err)
		problem = append(problem, num)
	}
	problems = append(problems, problem)
	return problems
}

func main() {
	inputFile, err := utils.GetInput(6)
	utils.HandleError(err)
	defer inputFile.Close()
	b, err := io.ReadAll(inputFile)
	utils.HandleError(err)

	s := string(b)
	problems := part2(s)
	numsPerProblem := strings.Count(s, "\n") - 1
	delimiter := " \n"
	numsAndOperators := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(delimiter, r)
	})

	l := len(numsAndOperators)
	numProblems := l / (numsPerProblem + 1)
	stringNums := numsAndOperators[:numProblems*numsPerProblem]
	operators := numsAndOperators[numProblems*numsPerProblem:]
	sum := 0
	fmt.Println(operators)
	fmt.Println(len(operators))
	fmt.Println(len(problems))
	for i := range problems {
		sum += performOp([]rune(operators[i])[0], problems[i]...)
	}
	fmt.Println(sum)

	nums := make([]int, len(stringNums))
	for index := range stringNums {
		num, err := strconv.Atoi(stringNums[index])
		utils.HandleError(err)
		nums[index] = num
	}
	sum = 0
	for i := range numProblems {
		problem := make([]int, 0, numsPerProblem)
		for j := range numsPerProblem {
			problem = append(problem, nums[i+numProblems*j])
		}
		sum += performOp([]rune(operators[i])[0], problem...)
	}
	fmt.Println(sum)

}
