package response

import "github.com/gofiber/fiber/v2"

// StatusCode represents application-specific status codes
type StatusCode string

const (
	// Success codes (2xx)
	GeneralSuccess  StatusCode = "S_2000"
	GeneralCreated  StatusCode = "S_2001"
	GeneralAccepted StatusCode = "S_2002"

	// Client error codes (4xx)
	GeneralBadRequest    StatusCode = "E_4000"
	ValidationError      StatusCode = "E_4001"
	GeneralUnauthorized  StatusCode = "E_4010"
	GeneralForbidden     StatusCode = "E_4030"
	GeneralNotFound      StatusCode = "E_4040"
	GeneralConflict      StatusCode = "E_4090"
	GeneralUnprocessable StatusCode = "E_4220"

	// Server error codes (5xx)
	GeneralServerError StatusCode = "E_5000"
	GeneralUnavailable StatusCode = "E_5030"
)

// statusCodeMap maps application status codes to HTTP codes and messages
var statusCodeMap = map[StatusCode]struct {
	HTTPCode int
	Message  string
}{
	GeneralSuccess:  {fiber.StatusOK, "Success"},
	GeneralCreated:  {fiber.StatusCreated, "Created"},
	GeneralAccepted: {fiber.StatusAccepted, "Accepted"},

	GeneralBadRequest:    {fiber.StatusBadRequest, "Bad Request"},
	ValidationError:      {fiber.StatusBadRequest, "Validation Error"},
	GeneralUnauthorized:  {fiber.StatusUnauthorized, "Unauthorized"},
	GeneralForbidden:     {fiber.StatusForbidden, "Forbidden"},
	GeneralNotFound:      {fiber.StatusNotFound, "Not Found"},
	GeneralConflict:      {fiber.StatusConflict, "Conflict"},
	GeneralUnprocessable: {fiber.StatusTeapot, "Unable to process since I'm a Teapot"},

	GeneralServerError: {fiber.StatusInternalServerError, "Internal Server Error"},
	GeneralUnavailable: {fiber.StatusServiceUnavailable, "Service Unavailable"},
}

// GetHTTPCode returns the corresponding HTTP status code
func (s StatusCode) GetHTTPCode() int {
	if info, exists := statusCodeMap[s]; exists {
		return info.HTTPCode
	}
	return 500 // Default to internal server error
}

// GetMessage returns the default message for the status code
func (s StatusCode) GetMessage() string {
	if info, exists := statusCodeMap[s]; exists {
		return info.Message
	}
	return "Internal Server Error"
}
