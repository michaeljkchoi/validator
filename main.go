package main

import (
	"fmt"
	"omnomsun/validator/cmd/validator"
)

type Params struct {
	Amount   int
	Resource string
}

func main() {
	params := Params{
		Amount:   10,
		Resource: "",
	}

	ruleset := validator.Ruleset{
		"Resource": []validator.Validator{
			validator.Required,
		},
	}

	fmt.Println("Testing validation")

	resp, err := validator.Validate(params, ruleset)

	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)
	}

	fmt.Printf("%+v", resp)
}
