package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func PostRequest[T any](url string, data interface{}, t *T, params ...string) (int, *response[T], error) {
	c := fiber.Client{}
	a := c.Post(url)
	a.Add("Authorization", "Bearer "+TOKEN)
	a.JSON(data)

	if len(params) > 0 {
		a.QueryString(params[0])
		a.QueryString(params[0])
	}

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return 0, nil, errs[0]
	}

	var wrapper response[T]
	err := json.Unmarshal(body, &wrapper)
	if err != nil {
		return statusCode, nil, err
	}

	if !wrapper.Success {
		if wrapper.Message == "" {
			wrapper.Message = "Something went wrong"
		}
		return statusCode, nil, fiber.NewError(statusCode, wrapper.Message)
	}

	*t = wrapper.Data

	return statusCode, &wrapper, nil
}

func GetRequest[T any](url string, data interface{}, t *T, params ...string) (int, *response[T], error) {
	c := fiber.Client{}
	a := c.Get(url)
	a.Add("Authorization", "Bearer "+TOKEN)

	if len(params) > 0 {
		a.QueryString(params[0])
	}

	statusCode, body, errs := a.Bytes()
	if len(errs) > 0 {
		return 0, nil, errs[0]
	}

	var wrapper response[T]
	err := json.Unmarshal(body, &wrapper)
	if err != nil {
		return statusCode, nil, err
	}

	if !wrapper.Success {
		if wrapper.Message == "" {
			wrapper.Message = "Something went wrong"
		}
		return statusCode, nil, fiber.NewError(statusCode, wrapper.Message)
	}

	*t = wrapper.Data

	return statusCode, &wrapper, nil
}
