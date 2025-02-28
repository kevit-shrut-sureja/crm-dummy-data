package main

import (
	"fmt"

	"github.com/google/uuid"
)

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
	var invitedString []string

	for _, w := range workspaceData {
		users := probArray(0, Users, false)

		for _, u := range users {
			url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/users/invite"
			var invitedurl string
			fmt.Println("Inviting user to workspace", url)
			s, res, err := PostRequest(url, u, &invitedurl)
			if err != nil {
				fmt.Println("Error inviting user to workspace", w.workspaceName)
				panic(err)
			}
			if s != 200 {
				fmt.Println("Error inviting user to workspace", w.workspaceName)
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", res)
				panic("something went wrong")
			}
			invitedString = append(invitedString, invitedurl)
		}
	}
	fmt.Println("Invited users to workspace", invitedString)

	for _, i := range invitedString {
		url := WORKSPACE_URL + "/invited-user/" + i

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
		if s != 200 {
			fmt.Println("Error confirming user invite")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %v\n", r)
			panic("something went wrong")
		}
	}

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
		if s != 200 {
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
