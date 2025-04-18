package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/jaswdr/faker/v2"
)

func randomPhoneNumber() string {
	num := rand.Int64N(9000000000) + 1000000000 // Generate a 10-digit number
	return fmt.Sprintf("%010d", num)            // Format as 10-digit string
}

func getRandomRealPhoneNumber() string {
	s := probArray(0.2, []string{"8160460050", "9909138575"}, true)
	if len(s) == 0 {
		return ""
	}
	return s[0]
}

func NewLead(w *workspaceInfo, F *faker.Faker) *Lead {
	s := probArray(0.3, sourceList, true)
	var source *string
	if len(s) == 0 {
		s = nil
	} else {
		source = &s[0]
	}
	return &Lead{
		LeadStageId:           probArray(0, w.stages, true)[0],
		OwnerId:               probArray(0, w.users, true)[0],
		Name:                  F.Person().Name(),
		Email:                 F.Internet().SafeEmail(),
		PhoneCountryCode:      ptr("+91"),
		Phone:                 safePtr[string](probSingle(0.5, randomPhoneNumber)),
		CompanyName:           safePtr[string](probSingle(0.4, F.Company().Name())),
		JobTitle:              safePtr[string](probSingle(0.4, F.Company().JobTitle())),
		Website:               safePtr[string](probSingle(0.5, F.Internet().Domain())),
		Source:                source,
		DealCountryCode:       ptr("INR"),
		DealSize:              safePtr[int](probSingle(0.4, rand.IntN(1000000))),
		ConversionProbability: safePtr[int](probSingle(0.5, rand.IntN(100))),
		Tags:                  probArray(0.2, w.tags, false),
		CustomFields:          GetRandomCustomFields(w, F, w.leadCustomFields),
		CreatedAt:             randomTimePicker(),
	}
}

func GetRandomCustomFields(w *workspaceInfo, F *faker.Faker, c []customField) []CustomFieldPayload {
	var customFields []CustomFieldPayload
	// c := probArray(0, w.customFields, false)

	for _, v := range c {
		x := CustomFieldPayload{
			ID:        v.ID,
			InputType: v.Type,
		}
		switch v.Type {
		case "text":
			if v.Name == "Custom_Field_Phone" {
				x.Value = ptr(getRandomRealPhoneNumber())
			} else {
				x.Value = ptr(F.Lorem().Sentence(5))
			}
		case "date":
			x.Value = ptr(randomTimePicker())
		case "select":
			z := probArray(0, v.Options, true)[0]
			x.Value = ptr(z)
		case "multiSelect":
			z := probArray(0, v.Options, false)
			x.MultipleValue = z
		}
		customFields = append(customFields, x)
	}

	return customFields
}

func CreateLeadApi(w *workspaceInfo, F *faker.Faker) <-chan int {

	resultCh := make(chan int)
	go func() {
		lead := NewLead(w, F)
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/lead"
		s, r, err := PostRequest(url, lead, &empty{})
		if err != nil {
			fmt.Println("Error creating lead")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %+v\n", r)
			fmt.Printf("Lead data %+v\n", lead)
			panic(err)
		}
		if s > 300 {
			fmt.Println("Error creating lead")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %+v\n", r)
			fmt.Printf("Lead data %+v\n", lead)
			panic("something went wrong with lead creation")
		}
		resultCh <- s
	}()
	return resultCh

}
