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

func main() {
	inputFile, err := utils.GetInput(6)
	utils.HandleError(err)
	defer inputFile.Close()
	b, err := io.ReadAll(inputFile)
	utils.HandleError(err)

	s := string(b)
	numsPerProblem := strings.Count(s, "\n") - 1
	delimiter := " \n"
	numsAndOperators := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(delimiter, r)
	})

	l := len(numsAndOperators)
	numProblems := l / (numsPerProblem + 1)
	stringNums := numsAndOperators[:numProblems*numsPerProblem]
	operators := numsAndOperators[numProblems*numsPerProblem:]

	nums := make([]int, len(stringNums))
	for index := range stringNums {
		num, err := strconv.Atoi(stringNums[index])
		utils.HandleError(err)
		nums[index] = num
	}
	sum := 0
	for i := range numProblems {
		problem := make([]int, 0, numsPerProblem)
		for j := range numsPerProblem {
			problem = append(problem, nums[i+numProblems*j])
		}
		sum += performOp([]rune(operators[i])[0], problem...)
	}
	fmt.Println(sum)

}
