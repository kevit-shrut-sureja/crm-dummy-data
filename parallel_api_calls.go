package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/jaswdr/faker/v2"
)

func leadInsertionCalls() {

	overallStart := time.Now() // Track total time for all workspaces

	for i := range workspaceData {
		ws := &workspaceData[i]
		fmt.Printf("\nStarting processing for %s (max records: %d)\n", ws.workspaceName, ws.maxRecords)

		var totalAPITime time.Duration
		var maxAPICallTime time.Duration
		totalAPICount := 0
		totalSuccessCount := 0
		workspaceStart := time.Now()

		for ws.records < ws.maxRecords {
			remaining := ws.maxRecords - ws.records
			currentBatch := BATCH_SIZE
			if remaining < BATCH_SIZE {
				currentBatch = remaining
			}

			var wg sync.WaitGroup
			batchDurationCh := make(chan time.Duration, currentBatch)
			statusCh := make(chan int, currentBatch)

			for j := 0; j < currentBatch; j++ {
				wg.Add(1)
				go func(w *workspaceInfo) {
					start := time.Now()
					var F = faker.New()
					status := <-CreateLeadApi(w, &F)
					elapsed := time.Since(start)
					batchDurationCh <- elapsed
					statusCh <- status
					wg.Done()
				}(ws)
			}

			wg.Wait()
			close(batchDurationCh)
			close(statusCh)

			batchSuccessCount := 0
			for s := range statusCh {
				if s == 200 {
					batchSuccessCount++
				}
			}

			for d := range batchDurationCh {
				totalAPITime += d
				if d > maxAPICallTime {
					maxAPICallTime = d
				}
				totalAPICount++
			}

			ws.records += currentBatch
			totalSuccessCount += batchSuccessCount

			fmt.Printf("Workspace %s: Batch completed with %d API calls, success count: %d, total records: %d\n",
				ws.workspaceName, currentBatch, batchSuccessCount, ws.records)
		}

		workspaceTotalTime := time.Since(workspaceStart)
		averageAPICallTime := time.Duration(0)
		if totalAPICount > 0 {
			averageAPICallTime = totalAPITime / time.Duration(totalAPICount)
		}

		fmt.Printf("\n--- Stats for %s ---\n", ws.workspaceName)
		fmt.Printf("Total API calls made: %d\n", ws.maxRecords)
		fmt.Printf("Overall success count: %d\n", totalSuccessCount)
		fmt.Printf("Average API call time: %v\n", averageAPICallTime)
		fmt.Printf("Maximum API call time: %v\n", maxAPICallTime)
		fmt.Printf("Total time taken for workspace: %v\n", workspaceTotalTime)
	}

	overallTotalTime := time.Since(overallStart)
	fmt.Printf("\n=== Overall Stats ===\n")
	fmt.Printf("Total time taken for all workspaces: %v\n", overallTotalTime)
}
