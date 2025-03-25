package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
)

const WORKSPACE_URL = BASE_URL + "/workspace"

var customFieldsType = []string{"text", "select", "date", "multiSelect"}

func workspaceSetupCalls() {
	// 1. make n workspace calls
	for i := 0; i < len(records); i++ {
		createSingleWorkspace(i, records[i])
	}
	// shuffe records
	rand.Shuffle(len(records), func(i, j int) {
		records[i], records[j] = records[j], records[i]
	})

	for i := range workspaceData {
		workspaceData[i].maxCustomersRecords = records[i]
	}

	// 2. get n new workspace and map them with records count
	getWorkspacesAndMapData()

	for _, w := range workspaceData {
		fmt.Printf("Workspace Name :: %s\n", w.workspaceName)
		fmt.Printf("Id :: %s\n", w.workspaceID)
		fmt.Printf("LeadsRecords :: %d\n", w.maxLeadsRecords)
		fmt.Println("CustomersRecords ::", w.maxCustomersRecords)
		fmt.Println("---------------------------------------------------")
	}

	fmt.Println("Workspaces created and mapped successfully")

	// each workspace
	// 3. insert static users data in all of them but random number of users
	inviteUserToWorkspace()
	fmt.Println("Users invited to workspace")

	// mapped users
	getWorkspaceUsers()
	fmt.Println("Users mapped with workspace")

	// 4. create new tags with a random count of 5~10
	CreateRandonTags()
	fmt.Println("Tags created")

	// 5. get tags and map them with workspace
	GetTagsForWorkspace()
	fmt.Println("Tags mapped with workspace")

	// 6. get stages and map them with workspace
	GetWorkspaceStages()
	fmt.Println("Stages mapped with workspace")

	// 7. create random custom field
	CreateRandomCustomFields("lead")
	fmt.Println("Custom fields created leads")
	CreateRandomCustomFields("customer")
	fmt.Println("Custom fields created customers")

	// 8. map them with workspace
	GetCustomFieldsForWorkspace("lead")
	GetCustomFieldsForWorkspace("customer")

	for _, w := range workspaceData {
		fmt.Printf("Workspace Name :: %s\n", w.workspaceName)
		fmt.Printf("Id :: %s\n", w.workspaceID)
		fmt.Printf("Records :: %d\n", w.maxLeadsRecords)
		fmt.Printf("Users :: %v\n", w.users)
		fmt.Printf("Tags :: %v\n", w.tags)
		fmt.Printf("Stages  :: %v\n", w.stages)

		for _, c := range w.leadCustomFields {
			fmt.Printf("Custom Field")
			fmt.Printf("ID :: %s\n", c.ID)
			fmt.Printf("Name :: %s\n", c.Name)
			fmt.Printf("Type:: %s\n", c.Type)
			fmt.Printf("options  :: %v\n", c.Options)
		}
		fmt.Println("---------------------------------------------------")
	}
}

func createSingleWorkspace(number, records int) {
	dto := WorkspaceDto{
		Name: WORKSPACE_NAME_PREFIX + fmt.Sprintf("-%d-%d", number, records),
	}

	statusCode, res, err := PostRequest(WORKSPACE_URL, dto, &empty{})
	if err != nil {
		fmt.Printf("Error creating workspace %d with records %d\n", number, records)
		fmt.Println("Error creating workspace", err)
		panic(err)
	}
	if statusCode > 299 {
		fmt.Printf("Error creating workspace %d with records %d\n", number, records)
		fmt.Printf("Status code: %d\n", statusCode)
		fmt.Printf("Response: %v\n", res)
		fmt.Println("Error creating workspace", err)
		panic("something went wrong")
	}

	workspaceData = append(workspaceData, workspaceInfo{
		workspaceName:   dto.Name,
		maxLeadsRecords: records,
	})

	fmt.Printf("Workspace %d created with records %d\n", number, records)
}

func getWorkspacesAndMapData() {
	var wk WorkspaceResponse
	statusCode, res, err := GetRequest(WORKSPACE_URL, nil, &wk, "page=1&pageSize=100")
	if err != nil {
		fmt.Println("Error getting workspaces")
		panic(err)
	}
	if statusCode > 299 {
		fmt.Println("Error getting workspaces")
		fmt.Printf("Status code: %d\n", statusCode)
		fmt.Printf("Response: %v\n", res)
		panic("something went wrong")
	}

	for i, w := range workspaceData {
		for _, ws := range wk.Workspaces {
			if ws.Name == w.workspaceName {
				fmt.Println("Matched", w.workspaceName, ws.Name)
				workspaceData[i].workspaceID = ws.ID
				fmt.Println("Workspace ID", ws.ID, "mapped with", workspaceData[i].workspaceID)
				break
			}
		}
	}

	fmt.Println("Workspaces mapped successfully")
	fmt.Printf("\n\nWorkspaces data: %+v\n\n", workspaceData)
}

func GetWorkspaceStages() {
	for i, w := range workspaceData {
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/stages"
		var stages []Stage
		s, r, err := GetRequest(url, nil, &stages, "page=1&pageSize=100")
		if err != nil {
			fmt.Println("Error getting stages")
			panic(err)
		}
		if s > 299 {
			fmt.Println("Error getting stages")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %v\n", r)
			panic("something went wrong")
		}
		var x []uuid.UUID
		for _, t := range stages {
			x = append(x, t.StageId)
		}
		workspaceData[i].stages = x

	}
}

func CreateRandomCustomFields(module string) {
	for _, w := range workspaceData {
		// customFields := probArray(0, customFieldsType, false)
		customFields := customFieldsType

		var customFieldsData []CreateCustomFieldDto

		for _, c := range customFields {
			var dto CreateCustomFieldDto
			dto.InputType = c
			dto.FieldName = "Custom_Field_" + c
			dto.Mandatory = false

			if c == "select" || c == "multiSelect" {
				n := rand.IntN(5) + 1
				var F = faker.New()
				options := F.Lorem().Words(n)
				var y []CustomFieldOptionValueDto
				for _, o := range options {
					y = append(y, CustomFieldOptionValueDto{
						Value: o,
					})
				}
				if dto.Data == nil {
					dto.Data = &CustomFieldDataDto{}
				}
				dto.Data.Options = y
			}
			customFieldsData = append(customFieldsData, dto)
		}

		customFieldsData = append(customFieldsData, CreateCustomFieldDto{
			FieldName: "Custom_Field_Phone",
			InputType: "text",
			Mandatory: false,
		})

		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/custom-fields"
		for _, c := range customFieldsData {
			s, r, err := PostRequest(url, c, &empty{}, "module="+module)
			if err != nil {
				fmt.Println("Error creating custom fields")
				panic(err)
			}
			if s > 299 {
				fmt.Println("Error creating custom fields")
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", r)
				panic("something went wrong with custom fields")
			}
		}
	}
}

func GetCustomFieldsForWorkspace(module string) {
	for i, w := range workspaceData {
		var customFields CustomFieldsResponse
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/custom-fields"
		s, r, err := GetRequest(url, nil, &customFields, "page=1&pageSize=100&module="+module)
		if err != nil {
			fmt.Println("Error getting custom fields")
			panic(err)
		}
		if s != 200 {
			fmt.Println("Error getting custom fields")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %v\n", r)
			panic("something went wrong with custom fields")
		}

		var cf []customField
		for _, c := range customFields.CF {
			cf = append(cf, customField{
				ID:      c.CustomFieldId,
				Name:    c.FieldName,
				Type:    c.InputType,
				Options: c.Data.Options,
			})
		}
		if module == "lead" {
			workspaceData[i].leadCustomFields = cf
		} else {
			workspaceData[i].customerCustomFields = cf
		}
	}
	fmt.Println("Custom fields fetched successfully")
}

func CreateRandonTags() {
	for _, w := range workspaceData {
		t := probArray(0, tagsNames, false)
		for _, name := range t {
			dto := TagDto{
				Name:  name,
				Color: probArray(0, tagColors, true)[0],
			}
			url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/tags"
			s, r, err := PostRequest(url, dto, &empty{})
			if err != nil {
				fmt.Println("Error creating tags")
				panic(err)
			}
			if s > 299 {
				fmt.Println("Error creating tags")
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", r)
				panic("something went wrong")
			}
		}
	}
}

func GetTagsForWorkspace() {
	for i, w := range workspaceData {
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/tags/all"
		var tags []Tag
		s, r, err := GetRequest(url, nil, &tags)
		if err != nil {
			fmt.Println("Error getting tags")
			panic(err)
		}
		if s > 299 {
			fmt.Println("Error getting tags")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %v\n", r)
			panic("something went wrong")
		}
		var x []uuid.UUID
		for _, t := range tags {
			x = append(x, t.ID)
		}
		workspaceData[i].tags = x
	}
}

/*
NOTE : as of now keeping the user data static, making dynamic is harder because aws cognito is involved

NOTE : make sure to only use the emails that are verified in aws cognito only for DEVELOPMENT purposes

*/

var Users = []InviteUserDto{
	{
		Email: "shrut.sureja@kevit.io",
		Permissions: PermissionsAndRoles{
			RoleName: "owner",
		},
	},
	{
		Email: "shrut.sureja+manager@kevit.io",
		Permissions: PermissionsAndRoles{
			RoleName: "admin",
		},
	},
	{
		Email: "shrut.sureja+1@kevit.io",
		Permissions: PermissionsAndRoles{
			RoleName: "admin",
		},
	},
	{
		Email: "shrut.sureja+2@kevit.io",
		Permissions: PermissionsAndRoles{
			RoleName: "admin",
		},
	},
	{
		Email: "shrut.sureja+4@kevit.io",
		Permissions: PermissionsAndRoles{
			RoleName: "admin",
		},
	},
}

func inviteUserToWorkspace() {
	// var invitedString []string

	for _, w := range workspaceData {
		users := probArray(0, Users, false)

		for _, u := range users {
			url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/users/invite"
			var invitedurl string
			fmt.Println("Inviting user to workspace", url)
			s, res, err := PostRequest(url, u, &invitedurl)
			// time.Sleep(500 * time.Millisecond)
			if err != nil {
				fmt.Println("Error inviting user to workspace", w.workspaceName)
				panic(err)
			}
			if s > 299 {
				fmt.Println("Error inviting user to workspace", w.workspaceName)
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", res)
				panic("something went wrong")
			}

			url = WORKSPACE_URL + "/invited-user/" + invitedurl

			data := struct {
				Status string `json:"status"`
			}{
				Status: "InitiateInvite",
			}

			s, r, err := PostRequest(url, data, &empty{})
			if err != nil {
				fmt.Println("Error confirming user invite")
				panic(err)
			}
			if s > 299 {
				fmt.Println("Error confirming user invite")
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", r)
				panic("something went wrong")
			}
		}
	}

	// for _, i := range invitedString {
	// 	url := WORKSPACE_URL + "/invited-user/" + i

	// 	data := struct {
	// 		Status string `json:"status"`
	// 	}{
	// 		Status: "InitiateInvite",
	// 	}

	// 	s, r, err := PostRequest(url, data, &empty{})
	// 	time.Sleep(1000 * time.Millisecond)
	// 	fmt.Println("Confirming user invite", url)
	// 	if err != nil {
	// 		fmt.Println("Error confirming user invite")
	// 		panic(err)
	// 	}
	// 	if s != 200 {
	// 		fmt.Println("Error confirming user invite")
	// 		fmt.Printf("Status code: %d\n", s)
	// 		fmt.Printf("Response: %v\n", r)
	// 		panic("something went wrong")
	// 	}
	// }

	fmt.Println("Confirmed user invite")
}

func getWorkspaceUsers() {
	for i, w := range workspaceData {
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/users/all"
		var users []UserResponse
		s, r, err := GetRequest(url, nil, &users, "page=1&limit=100")
		if err != nil {
			fmt.Println("Error getting workspace users")
			panic(err)
		}
		if s > 299 {
			fmt.Println("Error getting workspace users")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %v\n", r)
			panic("something went wrong")
		}

		var userIDs []uuid.UUID
		for _, u := range users {
			userIDs = append(userIDs, u.UserId)
		}

		workspaceData[i].users = userIDs
	}
}
