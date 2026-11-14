package ref_task_status

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type CreateRequest struct {
	UUID       string `json:"uuid" validate:"required,uuid"`
	TaskStatus string `json:"task_status" validate:"required"`
}

type ErrorResponse map[string]map[string]string

func newValidator() *validator.Validate {
	validate := validator.New()

	// Use the json tag as the field name (e.g. TaskStatus -> task_status)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" || name == "" {
			return fld.Name
		}
		return name
	})

	return validate
}

// Used for validate create task status
func CreateRequestValidation(request CreateRequest) ErrorResponse {
	validate := newValidator()
	err := validate.Struct(request)
	if err == nil {
		return nil
	}

	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return ErrorResponse{
			"error": {
				"validation": err.Error(),
			},
		}
	}

	result := make(ErrorResponse)
	for _, e := range verrs {
		if result[e.Field()] == nil {
			result[e.Field()] = map[string]string{}
		}
		result[e.Field()][e.Tag()] = message(e)
	}

	return result
}

// Used to generare message
func message(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s wajib diisi", label(e.Field()))
	case "min":
		return fmt.Sprintf("%s minimal %s karakter", label(e.Field()), e.Param())
	default:
		return fmt.Sprintf("%s tidak valid", label(e.Field()))
	}
}

// Used to generate label inside message
func label(field string) string {
	s := strings.ReplaceAll(field, "_", " ")
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
