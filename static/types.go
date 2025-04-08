package main

import (
	"github.com/google/uuid"
)

// for api calls
type response[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type empty struct{}

type Lead struct {
	LeadStageId           uuid.UUID            `json:"leadStageId"`
	OwnerId               uuid.UUID            `json:"ownerId"`
	Name                  string               `json:"name"`
	Email                 string               `json:"email"`
	PhoneCountryCode      *string              `json:"phoneCountryCode"`
	Phone                 *string              `json:"phone"`
	CompanyName           *string              `json:"companyName"`
	JobTitle              *string              `json:"jobTitle"`
	Website               *string              `json:"website"`
	Source                *string              `json:"source"`
	DealCountryCode       *string              `json:"dealCountryCode"`
	DealSize              *int                 `json:"dealSize"`
	ConversionProbability *int                 `json:"conversionProbability"`
	CustomFields          []CustomFieldPayload `json:"customFields"`
	Tags                  []uuid.UUID          `json:"tags"`
	CreatedAt             string               `json:"createdAt"`
}

type Customer struct {
	OwnerId          uuid.UUID            `json:"ownerId"`
	Name             string               `json:"name"`
	Email            string               `json:"email"`
	PhoneCountryCode *string              `json:"phoneCountryCode"`
	Phone            *string              `json:"phone"`
	CompanyName      *string              `json:"companyName"`
	JobTitle         *string              `json:"jobTitle"`
	Website          *string              `json:"website"`
	Source           *string              `json:"source"`
	CustomFields     []CustomFieldPayload `json:"customFields"`
	Tags             []uuid.UUID          `json:"tags"`
	CreatedAt        string               `json:"createdAt"`
}

type CustomFieldPayload struct {
	ID            uuid.UUID `json:"id"`
	InputType     string    `json:"inputType"`
	Value         *string   `json:"value"`
	MultipleValue []string  `json:"multipleValue"`
}
