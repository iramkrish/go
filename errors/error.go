package errors

import "fmt"

type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid field '%s': %s", e.Field, e.Msg)
}

func ValidateAge(age int) error {
	if age < 18 {
		return &ValidationError{Field: "age", Msg: "must be at least 18"}
	}
	return nil
}
