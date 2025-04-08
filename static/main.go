package main

// 1. create a workspace
const workspaceId = ""

// 2. Invite some users to the workspace
type userId string

var userIds = []userId{
	"cf3ab647-1ea1-4408-bd92-de07b6bf7ef3",
	"e645f63b-be8e-4e53-bb27-4412147a1329",
	"c4be0798-07e6-41a7-b747-56f4faa5d98d",
	"3a3dfe72-3d60-4902-bfe7-7e6e0101f349",
}

// 3. Stage list
type stageId string

var stageLeadMapping = map[stageId]int{
	"2e631ab7-def1-4b42-a27c-2d3f28359009": 13320,
	"af1e64a8-ee6a-4d3c-a1a9-94c93ce832fa": 8004,
	"13fffb9a-a852-4c5f-b285-ee205519be64": 5312,
	"ab22600f-20f2-4015-b6af-af8f96eaea94": 4352,
	"f0ba081a-707f-453f-8a70-f64f72bf0614": 2122,
	"d2d65aac-0aa8-41b9-aca7-00b0b755db51": 3432,
	"22be514b-2dbf-47d7-8102-e9d7ca4e2043": 9321,
	"3a9336cc-ed6a-4c94-9766-acdc6202abbe": 1232,
}

// 4. Tags List
type tagId string

var tagIds = []tagId{}

// Real Existing phone number
var phoneNumbers = []string{}

var sourceList = []string{}
var leadSourceMapping = map[string]int{
	"Google":   3000,
	"Facebook": 100,
	"LinkedIn": 100,
}

const (
	BASE_URL      = "http://localhost:3000"
	TOTAL_RECORDS = 50000
)

func main() {

}
