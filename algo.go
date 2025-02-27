package main

import (
	"fmt"
	"math/rand/v2"
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
			fmt.Println("last record :: ", x, "  sum exceed :: ", sum+x)
			x = TOTAL_RECORDS - sum
		}

		records = append(records, x)
		sum += x
	}

	fmt.Println("sum :: ", sum)

	if len(records) < TOTAL_RECORDS/MAX_RECORDS_PER_WORKSPACE {
		fmt.Println("Number of workspaces less than expected")
		panic("something wrong in algorithm")
	}
	return records
}
