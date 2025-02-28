package main

// super admin token to create new workspace
const (
	TOKEN                     = "eyJraWQiOiJwdFQxMjk4TWNva3g2aWpGUGwzbm8xQThNeWloMTFJcW1cL25aK2FSd0xkMD0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJmMWQzN2RhYS02MDAxLTcwMjktNmNjMS0wMTExZGNkODMxNWYiLCJkZXZpY2Vfa2V5IjoiYXAtc291dGgtMV8xZDk4YjU2Mi02YTY0LTQxMmItOGVkNi0yYmQwNTEwYjAyMDQiLCJpc3MiOiJodHRwczpcL1wvY29nbml0by1pZHAuYXAtc291dGgtMS5hbWF6b25hd3MuY29tXC9hcC1zb3V0aC0xX21ENVpmb2ZpVCIsImNsaWVudF9pZCI6IjFzNHQ0bDNuMTNyMWcxNWpzNXZxOWh0NnNuIiwib3JpZ2luX2p0aSI6ImFiNzNhNGZlLWRmMGEtNGEwNy05ZmVkLTgzOTc0OGVjMzBjMSIsImV2ZW50X2lkIjoiMjExZTg1YTItNDE2ZS00N2NhLTliNDUtMWEwMjJkMTVmNDdhIiwidG9rZW5fdXNlIjoiYWNjZXNzIiwic2NvcGUiOiJhd3MuY29nbml0by5zaWduaW4udXNlci5hZG1pbiIsImF1dGhfdGltZSI6MTc0MDU3Njk3NCwiZXhwIjoxNzQwODA2MDI1LCJpYXQiOjE3NDA3MTk2MjUsImp0aSI6Ijc2NGRlZDRkLTU4ZTgtNGU5OC04NDBiLTZhMzk0MDNjYTM1NyIsInVzZXJuYW1lIjoiZjFkMzdkYWEtNjAwMS03MDI5LTZjYzEtMDExMWRjZDgzMTVmIn0.I5xBt-ZwzslG7GwpWWWSQ9g8VXsRfs95PrWCgHBPGP1GeJLJdC68jVqRbrgA2Jp4UF5eZHY7Drs9LP0V6NWGi-AXM8QtizZHO4VvaIWKDo1AKqbFH-vRktuWtL-JzESrBuYSr6KE4jSDaE3xG8qXu4EQDBaHmagRVIbXRxtCW-ZSJ6hToIDd0OvyOoJcnErl9klubEJqVhweIuLg0CVGGASnqc5aNofh8eFzxx4IMWOkyCUcWkwvgayUZB68lSH9VAzdZyCPLw0sz7VlVr5mznd1bjLfm2ExZ_bnPNSrZIzN_wyjmfQaVr_0SyOb_sWOyZY9M5qvezuE-78wqDBGiw"
	BASE_URL                  = "http://localhost:3002"
	TOTAL_RECORDS             = 1000000
	MAX_RECORDS_PER_WORKSPACE = 200000
	MAX_WORKSPACES            = 60 // must be greater than TOTAL_RECORDS/MAX_PER_WORKSPACE
	MAX_REQ_PER_WORKER        = 100

	WORKSPACE_NAME_PREFIX = "aaaaaaaaaaaaaaa"
)

var records []int
var workspaceData data

var tagsNames = []string{
	"High-Value", "Cold Lead", "Hot Lead", "Follow-Up", "VIP Client", "Pending Approval", "New Customer", "Lost Lead", "Returning Customer", "Referral",
}
var tagColors = []string{
	"#E0E9C3", "#CECDCA", "#E9DAE2", "#CAE8FD", "#E4BDDD", "#D3D5ED", "#9AEDD1", "#ECDBA8",
}

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
	records = generateRandomRecords()

	apiCalls()
}
