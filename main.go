package main

import (
	"fmt"
	"github.com/omnisyle/validator/cmd/validator"
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
		"Amount": []validator.Validator{
			validator.InBetweenIntExclusive(11, 50),
		},
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
