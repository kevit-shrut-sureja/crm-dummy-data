package main

import (
	"time"

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

// Helper types
type data []workspaceInfo

type workspaceInfo struct {
	workspaceID   uuid.UUID
	workspaceName string
	maxRecords    int
	records       int
	users         []uuid.UUID
	tags          []uuid.UUID
	stages        []uuid.UUID
	customFields  []customField
}

type customField struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Options []string  `json:"options"`
}

// Types requied for api calls
type WorkspaceDto struct {
	Name string `json:"name"`
}

type WorkspaceResponse struct {
	Workspaces []struct {
		ID        uuid.UUID `json:"workspaceId"`
		Name      string    `json:"name"`
		UserCount int       `json:"userCount"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"workspaces"`
}

type InviteUserDto struct {
	Email       string              `json:"email"`
	Permissions PermissionsAndRoles `json:"permissions"`
}

type PermissionsAndRoles struct {
	RoleName string                 `json:"roleName" validate:"required,oneof=owner admin manager"`
	Modules  []DtoModulePermissions `json:"modules" validate:"omitempty,required,dive"`
}

// This is a duplicate of the models.ModulePermissions struct with validation tags
type DtoModulePermissions struct {
	Create     bool   `json:"create"`
	Read       bool   `json:"read"`
	Update     bool   `json:"update"`
	Delete     bool   `json:"delete"`
	ModuleName string `json:"moduleName" validate:"required"`
}

type UserResponse struct {
	UserId uuid.UUID `json:"userId"`
}

type TagDto struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Tag struct {
	ID uuid.UUID `json:"id"`
}

type Stage struct {
	StageId uuid.UUID `json:"stageId"`
}

type CreateCustomFieldDto struct {
	FieldName string              `json:"fieldName"`
	InputType string              `json:"inputType"`
	Mandatory bool                `json:"mandatory"`
	Data      *CustomFieldDataDto `json:"data"`
}

type CustomFieldDataDto struct {
	Options []CustomFieldOptionValueDto `json:"options"`
}

type CustomFieldOptionValueDto struct {
	Value string `json:"value"`
}

type CustomFieldsResponse struct {
	CF []struct {
		CustomFieldId uuid.UUID `json:"customFieldId"`
		FieldName     string    `json:"fieldName"`
		InputType     string    `json:"inputType"`
		Data          struct {
			Options []string `json:"options"`
		}
	} `json:"customFields"`
}

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

type CustomFieldPayload struct {
	ID            uuid.UUID `json:"id"`
	InputType     string    `json:"inputType"`
	Value         *string   `json:"value"`
	MultipleValue []string  `json:"multipleValue"`
}
