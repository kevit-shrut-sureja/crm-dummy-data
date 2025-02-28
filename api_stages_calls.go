package main

import (
	"fmt"

	"github.com/google/uuid"
)

const WORKSPACE_URL = BASE_URL + "/workspace"

func apiCalls() {
	// 1. make n workspace calls
	for i := 0; i < len(records); i++ {
		createSingleWorkspace(i, records[i])
	}
	// 2. get n new workspace and map them with records count
	getWorkspacesAndMapData()
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
	CreateRandomCustomFields()
	fmt.Println("Custom fields created")

	// 8. map them with workspace
	GetCustomFieldsForWorkspace()
	fmt.Printf("workspaceData %+v\n", workspaceData[0])
}

func createSingleWorkspace(number, records int) {
	dto := WorkspaceDto{
		Name: WORKSPACE_NAME_PREFIX + fmt.Sprintf("-%d-%d", number, records),
	}

	statusCode, res, err := PostRequest(WORKSPACE_URL, dto, &empty{})
	if err != nil {
		fmt.Printf("Error creating workspace %d with records %d\n", number, records)
		panic(err)
	}
	if statusCode != 200 {
		fmt.Printf("Error creating workspace %d with records %d\n", number, records)
		fmt.Printf("Status code: %d\n", statusCode)
		fmt.Printf("Response: %v\n", res)
		panic("something went wrong")
	}

	workspaceData = append(workspaceData, workspaceInfo{
		workspaceName: dto.Name,
		records:       records,
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
	if statusCode != 200 {
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
		if s != 200 {
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
