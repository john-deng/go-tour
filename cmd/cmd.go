package main

import (
	"fmt"
	"errors"
	"github.com/manifoldco/promptui"
)

func main() {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
		}
		return nil
	}

	username := promptui.Prompt{
		Label:    "Username",
	}
	u, err := username.Run()

	password := promptui.Prompt{
		Label:    "Password",
		Validate: validate,
		Mask:     '*',
	}

	p, err := password.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("Your username is %q\n", u)
	fmt.Printf("Your password is %q\n", p)
}
