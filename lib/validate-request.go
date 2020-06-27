package lib

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
)

type request struct {
	err []string
}

func (r *request) Validate(err error) []string {
	log.Printf("%#v", err)
	if e, ok := err.(validator.ValidationErrors); ok {
		for _, err := range e {
			var field string = strings.ToLower(err.Field())
			switch err.Tag() {
			case "required":
				r.err = append(r.err, fmt.Sprintf("%s is required", field))
			case "email":
				r.err = append(r.err, fmt.Sprintf("%s is not valid email", field))
			case "gte":
				r.err = append(r.err, fmt.Sprintf("%s value must be greater than %s", field, err.Param()))
			case "lte":
				r.err = append(r.err, fmt.Sprintf("%s value must be lower than %s", field, err.Param()))
			}
		}
	}

	return r.err
}
