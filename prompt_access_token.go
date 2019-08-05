package main

import "github.com/manifoldco/promptui"

func promptAccessToken() error {
	prompt := promptui.Prompt{
		Label:    "Input you Gitlab access token",
		Validate: validateAccessToken,
	}
	result, err := prompt.Run()
	if err != nil {
		return err
	}
	return saveAccessToken(result)
}
