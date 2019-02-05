package validator

import (
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
		r := reflect.ValueOf(params)
		val := reflect.Indirect(r).FieldByName(k)

		for _, rule := range rules {
			resp := rule(val.Interface(), params)

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

func Required(value interface{}, params interface{}) []string {
	if value == nil || value == "" {
		return []string{"cannot be blank"}
	}

	return nil
}
