package main

import "fmt"

// super admin token to create new workspace
const (
	TOKEN                     = "eyJraWQiOiJwdFQxMjk4TWNva3g2aWpGUGwzbm8xQThNeWloMTFJcW1cL25aK2FSd0xkMD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJmMWQzN2RhYS02MDAxLTcwMjktNmNjMS0wMTExZGNkODMxNWYiLCJkZXZpY2Vfa2V5IjoiYXAtc291dGgtMV8xZDk4YjU2Mi02YTY0LTQxMmItOGVkNi0yYmQwNTEwYjAyMDQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtc291dGgtMS5hbWF6b25hd3MuY29tXC9hcC1zb3V0aC0xX21ENVpmb2ZpVCIsImNsaWVudF9pZCI6IjFzNHQ0bDNuMTNyMWcxNWpzNXZxOWh0NnNuIiwib3JpZ2luX2p0aSI6IjE2M2ZhNDgwLTZjYzAtNDY1Mi1hY2MxLWVmMWQ0NWQ5MGUzNiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NDEwMDUzMzEsImV4cCI6MTc0MTA5MTczMSwiaWF0IjoxNzQxMDA1MzMxLCJqdGkiOiJhNzNlYzEwNS0zN2YyLTQwZDMtOTgzOS1hNjM3NzlmOTU0ZDciLCJ1c2VybmFtZSI6ImYxZDM3ZGFhLTYwMDEtNzAyOS02Y2MxLTAxMTFkY2Q4MzE1ZiJ9.JUyiGnT71ZmYA8H8JLzxz6jJ4Sm8mX7nvp9CmSc78hZ6LkEV_tFKzOm8FaCfy1x3x_-U_bNKzRqlqaFXriu8gaBdCJteeB_IomV70BlpbDoMC-FzdasuKhE4zqOFVz0puwMh7sKz9jyI1JaPTBpDxRZTLSr2R20Qs8_GAhBJ-5uE4QQNBtb3H2TDJLb4nm9l1J5ztF6ZrRxDoCX4_HHo9arW2UmAYwBDbtOxckG7HNze-yWtOaWlqBMhBeKEZNMySt97_lJ0WRndOHkgfVVScRIYAW5uuHnX-nvcpHyztrlVO2sxw9OhyiYGv4wjre8TUIKhbJZigqDVf6Z9_UCz1Q"
	BASE_URL                  = "http://localhost:3003"
	TOTAL_RECORDS             = 1000000
	MAX_RECORDS_PER_WORKSPACE = 200000
	MAX_WORKSPACES            = 60 // must be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE
	BATCH_SIZE                = 1000

	WORKSPACE_NAME_PREFIX = "kevit"
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
	customerInsertionCalls()
}
