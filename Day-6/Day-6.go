package main

import (
	"fmt"
	"github.com/Shuddown/Advent-of-Code/utils"
	"io"
)

func main() {
	inputFile, err := utils.GetInput(6)
	utils.HandleError(err)
	defer inputFile.Close()
	b, err := io.ReadAll(inputFile)
	utils.HandleError(err)
	fmt.Println(string(b))
}

