package main

import "fmt"

// super admin token to create new workspace
const (
	TOKEN                     = "eyJraWQiOiJwdFQxMjk4TWNva3g2aWpGUGwzbm8xQThNeWloMTFJcW1cL25aK2FSd0xkMD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJmMWQzN2RhYS02MDAxLTcwMjktNmNjMS0wMTExZGNkODMxNWYiLCJkZXZpY2Vfa2V5IjoiYXAtc291dGgtMV8xZDk4YjU2Mi02YTY0LTQxMmItOGVkNi0yYmQwNTEwYjAyMDQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtc291dGgtMS5hbWF6b25hd3MuY29tXC9hcC1zb3V0aC0xX21ENVpmb2ZpVCIsImNsaWVudF9pZCI6IjFzNHQ0bDNuMTNyMWcxNWpzNXZxOWh0NnNuIiwib3JpZ2luX2p0aSI6ImZlMTVkYzdlLTUxZTMtNDgwYy1iMjQ4LWNmMWQ0YTU1NzM1ZSIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NDA3MjAyNDEsImV4cCI6MTc0MDgwNjY0MSwiaWF0IjoxNzQwNzIwMjQxLCJqdGkiOiIzNGFkODIwYi1lOTA1LTQzMjMtODQ0ZS0wYWJkZDA5YTczMDAiLCJ1c2VybmFtZSI6ImYxZDM3ZGFhLTYwMDEtNzAyOS02Y2MxLTAxMTFkY2Q4MzE1ZiJ9.IgdRFhoSkoCd7iTrm_qtZt6CNc6orZxVS9mrHLI86LZX7AffB9ygEJ7tLrO296Gagm_H78Vs8KsUFYRGxtrCBS_O0R5eaTeyDAiv_VlB82MmOScL5Oj1EApY6tEwtHck-mJSByyogaNtDjwUqQf3NXX6xUnGrKyTA3YlniydufRslKUcVzmbXmfWxnisRginwxyuJ90RX2y_3-2rD-_pv7cMAeTlfqujtdPdpt_8yw3UZ0_i5vARqSfdh81RXj0bPifF3IAZ6BieY5RN_WJD-SVZ1UZAa3gUzvde6PvFMSbDqzIN_xE3fGvzrcAbNPgesy2jBT9o4ddj0G1qxI2NTQ"
	BASE_URL                  = "http://localhost:3002"
	TOTAL_RECORDS             = 1000000
	MAX_RECORDS_PER_WORKSPACE = 200000
	MAX_WORKSPACES            = 60 // must be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE
	BATCH_SIZE                = 1000

	WORKSPACE_NAME_PREFIX = "kevit-"
)

var records []int
var workspaceData data

var tagsNames = []string{
	"High-Value", "Cold Lead", "Hot Lead", "Follow-Up", "VIP Client", "Pending Approval", "New Customer", "Lost Lead", "Returning Customer", "Referral",
}
var tagColors = []string{
	"#E0E9C3", "#CECDCA", "#E9DAE2", "#CAE8FD", "#E4BDDD", "#D3D5ED", "#9AEDD1", "#ECDBA8",
}
var sourceList = []string{"Social Media", "Website", "LinkedIn", "WhatsApp", "Referral"}

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
	records = generateRandomRecords()

	workspaceSetupCalls()

	fmt.Println("All workspace setup calls completed successfully")

	leadInsertionCalls()
}
