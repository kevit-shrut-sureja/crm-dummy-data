package main

import (
	"fmt"
	"math/rand/v2"
)

var customFieldsType = []string{"text", "select", "date", "multiSelect"}

func CreateRandomCustomFields() {
	for _, w := range workspaceData {
		customFields := probArray(0, customFieldsType, false)

		var customFieldsData []CreateCustomFieldDto

		for _, c := range customFields {
			var dto CreateCustomFieldDto
			dto.InputType = c
			dto.FieldName = "Custom Field " + c
			dto.Mandatory = false

			if c == "select" || c == "multiSelect" {
				n := rand.IntN(5) + 1
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

		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/custom-fields"
		for _, c := range customFieldsData {
			s, r, err := PostRequest(url, c, &empty{}, "module=lead")
			if err != nil {
				fmt.Println("Error creating custom fields")
				panic(err)
			}
			if s != 200 {
				fmt.Println("Error creating custom fields")
				fmt.Printf("Status code: %d\n", s)
				fmt.Printf("Response: %v\n", r)
				panic("something went wrong with custom fields")
			}
		}
	}
}

func GetCustomFieldsForWorkspace() {
	for i, w := range workspaceData {
		var customFields CustomFieldsResponse
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/custom-fields"
		s, r, err := GetRequest(url, nil, &customFields, "page=1&pageSize=100&module=lead")
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
		workspaceData[i].customFields = cf
	}
	fmt.Println("Custom fields fetched successfully")
}
