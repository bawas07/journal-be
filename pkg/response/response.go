package response

import (
	"mindscribe-be/pkg/validation"

	"github.com/gofiber/fiber/v2"
)

// Response represents the standard API response structure
type Response struct {
	Code     StatusCode           `json:"code"`
	Message  string               `json:"message"`
	Data     interface{}          `json:"data,omitempty"`
	Validate *validation.Validate `json:"-"`
}

func New(validate *validation.Validate) *Response {
	return &Response{
		Validate: validate,
	}
}

func (r *Response) Ok(c *fiber.Ctx, code StatusCode, data interface{}) error {
	return c.Status(code.GetHTTPCode()).JSON(Response{
		Code:    code,
		Message: code.GetMessage(),
		Data:    data,
	})
}

func (r *Response) OkWithMessage(c *fiber.Ctx, code StatusCode, message string, data interface{}) error {
	return c.Status(code.GetHTTPCode()).JSON(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (r *Response) Fail(c *fiber.Ctx, code StatusCode, data interface{}) error {
	return c.Status(code.GetHTTPCode()).JSON(Response{
		Code:    code,
		Message: code.GetMessage(),
		Data:    data,
	})
}

func (r *Response) FailWithMessage(c *fiber.Ctx, code StatusCode, message string, data interface{}) error {
	return c.Status(code.GetHTTPCode()).JSON(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func (r *Response) ValidationError(c *fiber.Ctx, err error) error {
	data := r.Validate.FormatValidationErrors(err)
	code := ValidationError
	return c.Status(code.GetHTTPCode()).JSON(Response{
		Code:    code,
		Message: "Validation Error",
		Data:    data,
	})
}
