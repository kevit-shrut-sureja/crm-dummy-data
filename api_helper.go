package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type response[T any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func PostRequest[T any](url string, data interface{}, t *T) (int, *T, error) {
	c := fiber.Client{}
	a := c.Post(url)
	a.Add("Authorization", "Bearer "+TOKEN)
	a.JSON(data)
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

	return statusCode, t, nil
}

func GetRequest[T any](url string, data interface{}, t *T) (int, *T, error) {
	c := fiber.Client{}
	a := c.Get(url)
	a.Add("Authorization", "Bearer "+TOKEN)

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

	return statusCode, t, nil
}
