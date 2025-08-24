package utils

import (
	"blog-api/core/models"

	"github.com/gofiber/fiber/v2"
)

// The function `UnknownMethod` returns a JSON response with a "Method Not Allowed" message and status
// code when an unknown method is used.
func MethodNotAllow(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
	// 	"code":    fiber.StatusMethodNotAllowed,
	// 	"status":  false,
	// 	"data":    "",
	// 	"message": "Method Not Allowed",
	// })
	return c.Status(fiber.StatusMethodNotAllowed).JSON(
		ReturnResponse(fiber.StatusMethodNotAllowed, "Method Not Allowed", ""),
	)
}

// The function `BodyParserFail` returns a JSON response with a status code of 400 and an error message
// for incorrect payload format.
func BodyParserFail(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 	"code":    fiber.StatusBadRequest,
	// 	"status":  false,
	// 	"message": "incorrect payload format",
	// 	"data":    "",
	// })
	return c.Status(fiber.StatusBadRequest).JSON(
		ReturnResponse(fiber.StatusBadRequest, "incorrect payload format", ""),
	)
}

// The function `ParamParserFail` returns a JSON response indicating a failure to parse parameters with
// a bad request status.
func ParamParserFail(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 	"code":    fiber.StatusBadRequest,
	// 	"status":  false,
	// 	"message": "Failed to parse params",
	// 	"data":    "",
	// })
	return c.Status(fiber.StatusBadRequest).JSON(
		ReturnResponse(fiber.StatusBadRequest, "Failed to parse params", ""),
	)
}

// The function `QueryParserFail` returns a JSON response indicating a failure to parse query
// parameters with a specific message.
func QueryParserFail(c *fiber.Ctx) error {
	// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 	"code":    fiber.StatusBadRequest,
	// 	"status":  false,
	// 	"message": "Failed to parse query params",
	// 	"data":    "",
	// })
	return c.Status(fiber.StatusBadRequest).JSON(
		ReturnResponse(fiber.StatusBadRequest, "Failed to parse query params", ""),
	)
}

// The function `ErrorFormat` formats an error response with a specified code and message in a JSON
// format using the Fiber framework in Go.
func ErrorFormat(c *fiber.Ctx, code int, msg string) error {
	// return c.Status(code).JSON(fiber.Map{
	// 	"code":    code,
	// 	"status":  false,
	// 	"message": msg,
	// 	"data":    "",
	// })
	return c.Status(code).JSON(
		ReturnResponse(code, msg, ""),
	)
}

// The function `SuccessFormat` formats a success response with a status code, message, and data in a
// Fiber context.
func SuccessFormat(c *fiber.Ctx, code int, msg string, data interface{}) error {
	tmp := data
	if tmp == nil {
		tmp = ""
	}
	// return c.Status(code).JSON(fiber.Map{
	// 	"code":    code,
	// 	"status":  true,
	// 	"message": msg,
	// 	"data":    tmp,
	// })
	return c.Status(code).JSON(
		ReturnResponse(code, msg, tmp),
	)
}

func ReturnResponse(code int, message string, data interface{}) (result models.Response) {
	success := false
	if code >= 200 && code < 300 {
		success = true
	}
	// คืนค่า response
	result = models.Response{
		Status:  success,
		Code:    code,
		Message: message,
		Data:    data,
	}
	return result
}
