package main

import "fmt"

// super admin token to create new workspace
const (
	TOKEN                     = "eyJraWQiOiJwdFQxMjk4TWNva3g2aWpGUGwzbm8xQThNeWloMTFJcW1cL25aK2FSd0xkMD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJmMWQzN2RhYS02MDAxLTcwMjktNmNjMS0wMTExZGNkODMxNWYiLCJkZXZpY2Vfa2V5IjoiYXAtc291dGgtMV82NjAxMzNkNi1jODIzLTRlMTAtODhhNi1hYjgxZWZiMjExMjgiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtc291dGgtMS5hbWF6b25hd3MuY29tXC9hcC1zb3V0aC0xX21ENVpmb2ZpVCIsImNsaWVudF9pZCI6IjFzNHQ0bDNuMTNyMWcxNWpzNXZxOWh0NnNuIiwib3JpZ2luX2p0aSI6IjFkZjMwYzA1LTA0MTQtNDJhNC05ZDM2LWM1MzY4M2I3NThhMiIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4iLCJhdXRoX3RpbWUiOjE3NDI4MjA2MTQsImV4cCI6MTc0MjkwNzAxNCwiaWF0IjoxNzQyODIwNjE0LCJqdGkiOiI0Njc2NzgxMC1jNTZiLTQ0ZTAtYTc3YS0xMWJiMGE0ODUxZTkiLCJ1c2VybmFtZSI6ImYxZDM3ZGFhLTYwMDEtNzAyOS02Y2MxLTAxMTFkY2Q4MzE1ZiJ9.J8cRqEzcOEPaxFs9FPNJ9o4-Y5aw20uWajVNkjPHtAjDb7J2zW-K0W7i0gHRl1Z2Tm4qpV94crDRoLIsxMJZA9M4KYWR8ohdOCeXEh83loEC_G7qST17A5FKhOQCw2efOQaPRSITmtbfbv-R-yCuhabhna-kp6B24lotL_wPcwHb8rX1rsTvn69pVursMcQP6T8236j49NLiKvCqF60eH_0ueuLNlHCE1iGnOa3V5eUmO7l76mAG0EdPV20nHueeRuVrvg0v50G4piRcnvGmqPhRBg_YJpwhKFtZ1Xfp3UPrb-qyIOpqNvkDD8esZbgXGUSOgWQQT8wOfLgSCAxDtw"
	BASE_URL                  = "http://localhost:3000"
	TOTAL_RECORDS             = 1000000
	MAX_RECORDS_PER_WORKSPACE = 200000
	MAX_WORKSPACES            = 30 // must be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE
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
