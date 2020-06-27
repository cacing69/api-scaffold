package lib

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func Validate(err error) []string {
	var res []string

	if e, ok := err.(validator.ValidationErrors); ok {
		for _, err := range e {
			var field string = strings.ToLower(err.Field())
			switch err.Tag() {
			case "required":
				res = append(res, fmt.Sprintf("%s is required", field))
			case "email":
				res = append(res, fmt.Sprintf("%s is not valid email", field))
			case "gte":
				res = append(res, fmt.Sprintf("%s value must be greater than %s", field, err.Param()))
			case "lte":
				res = append(res, fmt.Sprintf("%s value must be lower than %s", field, err.Param()))
			}
		}
	}

	return res
}
