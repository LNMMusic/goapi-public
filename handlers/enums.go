package handlers

import (
	"fmt"
	"errors"
)

type Enum struct {
	Name	string
	Keys	[]string
}
// methods
func (e *Enum) Validate(str string) error {
	for _, key := range e.Keys {
		if key == str {
			return nil
		}
	}
	return errors.New(fmt.Sprintf("[Enum %s Error] -> <%s> is invalid. Try with [%#v]",e.Name, str, e.Keys))
}

func ValidateEnums(enums map[string]*Enum) []string {
	var errors []string
	
	for str, enum := range enums {
		if err := enum.Validate(str); err != nil {
			errors = append(errors, err.Error())
		}
	}

	return errors
}