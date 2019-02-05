package validator

import (
	"fmt"
	"reflect"
)

type Validator func(value interface{}, params interface{}) []string

type Ruleset map[string][]Validator

type ValidatorResponse struct {
	errors map[string][]string
}

func Validate(params interface{}, ruleset Ruleset) (ValidatorResponse, error) {
	result := map[string][]string{}

	for k, rules := range ruleset {
		val := getFieldValue(k, params)

		for _, rule := range rules {
			resp := rule(val, params)

			if resp != nil {
				curErrs, ok := result[k]

				if !ok {
					curErrs = []string{}
				}

				newErrs := append(curErrs, resp...)

				result[k] = newErrs
			}
		}
	}

	return ValidatorResponse{result}, nil
}

func getFieldValue(fieldName string, struc interface{}) interface{} {
	r := reflect.ValueOf(struc)
	val := reflect.Indirect(r).FieldByName(fieldName)

	return val.Interface()
}

func Required(value interface{}, params interface{}) []string {
	if value == nil || value == "" {
		return []string{"cannot be blank"}
	}

	return nil
}

func InBetweenIntExclusive(min int, max int) Validator {

	return func(value interface{}, params interface{}) []string {
		castVal := value.(int)
		if castVal < min || castVal > max {
			return []string{fmt.Sprintf("must be in between %d and %d", min, max)}
		}

		return nil
	}
}
