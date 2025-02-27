package main

import "fmt"

// super admin token to create new workspace
const (
	TOKEN                     = "asdasd"
	BASE_URL                  = "http://localhost:3000"
	TOTAL_RECORDS             = 1000000
	MAX_RECORDS_PER_WORKSPACE = 200000
	MAX_WORKSPACES            = 20 // must be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE
	MAX_REQ_PER_WORKER        = 100
)

func init() {
	// simple checks
	if TOKEN == "" {
		panic("TOKEN is required")
	}
	if BASE_URL == "" {
		panic("BASE_URL is required")
	}
	if TOTAL_RECORDS == 0 || TOTAL_RECORDS < 0 {
		panic("TOTAL_RECORDS is required")
	}
	if MAX_RECORDS_PER_WORKSPACE == 0 || MAX_RECORDS_PER_WORKSPACE < 0 {
		panic("MAX_PER_WORKSPACE is required")
	}
	if MAX_WORKSPACES == 0 || MAX_WORKSPACES < 0 {
		panic("MAX_WORKSPACES is required")
	}
	if MAX_REQ_PER_WORKER == 0 || MAX_REQ_PER_WORKER < 0 {
		panic("MAX_REQ_PER_WORKER is required")
	}

	// Logical checks
	if MAX_RECORDS_PER_WORKSPACE > TOTAL_RECORDS {
		panic("MAX_PER_WORKSPACE should be less than TOTAL_RECORDS")
	}
	if MAX_WORKSPACES < TOTAL_RECORDS/MAX_RECORDS_PER_WORKSPACE {
		panic("MAX_WORKSPACES should be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE")
	}
}

func main() {
	// Generate random workspace count
	y := getRandomRecordsPerWorkspace()
	fmt.Println(y, len(y))
}
