package main

import (
	"fmt"

	"github.com/jaswdr/faker/v2"
)

func NewCustomer(w *workspaceInfo, F *faker.Faker) *Customer {
	s := probArray(0.3, sourceList, true)
	var source *string
	if len(s) == 0 {
		s = nil
	} else {
		source = &s[0]
	}
	return &Customer{
		OwnerId:          probArray(0, w.users, true)[0],
		Name:             F.Person().Name(),
		Email:            F.Internet().SafeEmail(),
		PhoneCountryCode: ptr("+91"),
		Phone:            safePtr[string](probSingle(0.5, randomPhoneNumber)),
		CompanyName:      safePtr[string](probSingle(0.4, F.Company().Name())),
		JobTitle:         safePtr[string](probSingle(0.4, F.Company().JobTitle())),
		Website:          safePtr[string](probSingle(0.5, F.Internet().Domain())),
		Source:           source,
		Tags:             probArray(0.2, w.tags, false),
		CustomFields:     GetRandomCustomFields(w, F, w.customerCustomFields),
		CreatedAt:        randomTimePicker(),
	}
}

func CreateCustomerApi(w *workspaceInfo, F *faker.Faker) <-chan int {

	resultCh := make(chan int)
	go func() {
		customer := NewCustomer(w, F)
		url := WORKSPACE_URL + "/" + w.workspaceID.String() + "/customer"
		s, r, err := PostRequest(url, customer, &empty{})
		if err != nil {
			fmt.Println("Error creating customer")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %+v\n", r)
			fmt.Printf("Customer data %+v\n", customer)
			panic(err)
		}
		if s > 300 {
			fmt.Println("Error creating customer")
			fmt.Printf("Status code: %d\n", s)
			fmt.Printf("Response: %+v\n", r)
			fmt.Printf("customer data %+v\n", customer)
			panic("something went wrong with customer creation")
		}
		resultCh <- s
	}()
	return resultCh

}
