package main

import (
	"fmt"

	"github.com/google/uuid"
)

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
			if s != 200 {
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
		if s != 200 {
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
