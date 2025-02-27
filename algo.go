package main

import (
	"fmt"
	"math/rand/v2"
)

func getRandomRecordsPerWorkspace() []int {
	var records []int

	maximumRecords := MAX_RECORDS_PER_WORKSPACE
	minimumRecords := TOTAL_RECORDS / MAX_WORKSPACES
	fmt.Println(minimumRecords)

	sum := 0

	for sum < TOTAL_RECORDS {
		x := rand.IntN(maximumRecords-minimumRecords) + minimumRecords

		// ensure we dont exceed total records
		if sum+x > TOTAL_RECORDS {
			x = TOTAL_RECORDS - sum
		}
		records = append(records, x)
		sum += x
	}

	if len(records) < TOTAL_RECORDS/MAX_RECORDS_PER_WORKSPACE {
		fmt.Println("Number of workspaces less than expected")
		panic("something wrong in algorithm")
	}
	return records
}
