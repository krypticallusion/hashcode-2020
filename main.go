package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bufioR := bufio.NewScanner(os.Stdin)
	var firstLine string
	var secondLine string

	count := 1

	for bufioR.Scan() {
		if count == 1 {
			firstLine = bufioR.Text()
			count--
		} else {
			secondLine = bufioR.Text()
			break
		}
	}

	firstArray := strings.Split(firstLine, " ")
	secondArray := strings.Split(secondLine, " ")

	maxPizza, _ := strconv.Atoi(firstArray[0])
	//numberPizza, _ := strconv.Atoi(firstArray[1])

	//convert secondArray of strings to ints
	var secondIntArray []int
	for _, item := range secondArray {
		itemInt, _ := strconv.Atoi(item)
		secondIntArray = append(secondIntArray, itemInt)
	}

	var finalRes []int
	res := true

	for i := len(secondIntArray) - 1; i >= 0; i-- {
		var tempSolution []int

		sum := 0

		if res {
			for j := i; j >= 0; j-- {
				value := secondIntArray[j]
				if sum+value == maxPizza {
					sum = sum + value
					tempSolution = append(tempSolution, j)
					res = false
				} else if sum+value < maxPizza {
					sum = sum + value
					tempSolution = append(tempSolution, j)
				}
			}

			if sum > 0  {
				finalRes = tempSolution
			}
		}
	}

	var finalStr string
	for _, item := range Reverse(finalRes) {
		finalStr = finalStr + fmt.Sprintf("%v", item) + " "
	}


	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println(len(finalRes))
	fmt.Println(finalStr)

}

func Reverse(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}

	return slice
}
