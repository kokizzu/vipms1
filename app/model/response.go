package model

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Data  interface{}
	Error error
}


func NotValid(c *fiber.Ctx, notValid bool, errMsg string) bool {
	if notValid {
		c.JSON(Response{
			Error: errors.New(errMsg),
		})
		return true
	}
	return false
}

func CreateResponse(c *fiber.Ctx, data interface{}, err error) error {
	return c.JSON(Response{
		Data: data,
		Error: err,
	})
}


