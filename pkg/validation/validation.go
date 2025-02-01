package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Validate struct {
	validator *validator.Validate
	db        *sqlx.DB
	log       *zap.Logger
}

func NewValidate(db *sqlx.DB, log *zap.Logger) *Validate {
	validator := validator.New()
	validate := &Validate{
		validator: validator,
		db:        db,
		log:       log,
	}
	validator.RegisterValidation("isUnique", validate.isUnique)
	return validate
}

func (v *Validate) ValidateStruct(s interface{}) error {
	return v.validator.Struct(s)
}

// Function to format validation errors
func (v *Validate) FormatValidationErrors(err error) []map[string]string {
	var errors []map[string]string

	// Check if the error is a ValidationErrors type
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrs {
			// Create a structured error object
			errors = append(errors, map[string]string{
				"field":   fieldErr.Field(),
				"message": getValidationErrorMessage(fieldErr),
			})
		}
	}

	return errors
}

// Helper function to get user-friendly messages
func getValidationErrorMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return err.Field() + " field is required"
	case "email":
		return err.Field() + " must be a valid email address"
	case "min":
		return err.Field() + " must be at least " + err.Param() + " characters long"
	case "max":
		return err.Field() + " must be at most " + err.Param() + " characters long"
	case "gte":
		return err.Field() + " must be greater than or equal to " + err.Param()
	case "lte":
		return err.Field() + " must be less than or equal to " + err.Param()
	case "isUnique":
		return err.Field() + " already exist"
	default:
		return err.Field() + " is invalid"
	}
}
