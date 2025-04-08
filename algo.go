package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"sort"
	"strings"
)

func getRandomRecordsPerWorkspace() []int {
	var records []int

	maximumRecords := MAX_RECORDS_PER_WORKSPACE
	minimumRecords := TOTAL_RECORDS / MAX_WORKSPACES

	sum := 0
	for sum < TOTAL_RECORDS {
		var x int
		r := rand.IntN(100)

		if r < 37 {
			x = rand.IntN(minimumRecords/2) + 1
		} else if r < 60 {
			x = rand.IntN(maximumRecords-minimumRecords) + minimumRecords
		} else {
			x = rand.IntN(maximumRecords-minimumRecords) + minimumRecords/2
		}

		// ensure we dont exceed total records
		if sum+x > TOTAL_RECORDS {
			x = TOTAL_RECORDS - sum
		}

		records = append(records, x)
		sum += x
	}

	if len(records) < TOTAL_RECORDS/MAX_RECORDS_PER_WORKSPACE || len(records) > MAX_WORKSPACES {
		panic("something wrong in algorithm")
	}
	return records
}

func generateRandomRecords() []int {
	for true {
		y := getRandomRecordsPerWorkspace()
		sort.Ints(y)
		fmt.Println("Records generated for all workspaces ::")
		fmt.Println(y, len(y))
		fmt.Printf("\n\n")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Do you want to continue? (yes/no): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "yes" || input == "y" {
			fmt.Println("Moving ahead with the generated records...")
			fmt.Printf("\n\n")
			return y
		} else if input == "no" || input == "n" {
			fmt.Println("Regenerating random records...")
		} else {
			fmt.Println("Invalid input! Please enter yes or no.")
		}
	}
	return []int{}
}
