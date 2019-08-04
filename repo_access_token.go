package main

import (
	"github.com/zalando/go-keyring"
	"os/user"
)

const serviceName = "zbumper gitlab access token"

func getUsername() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.Username, nil
}

func getAccessToken() (string, error) {
	username, err := getUsername()
	if err != nil {
		return "", err
	}
	result, err := keyring.Get(serviceName, username)
	if err != nil {
		return "", err
	}
	return result, nil
}

func saveAccessToken(accessToken string) error {
	username, err := getUsername()
	if err != nil {
		return err
	}
	return keyring.Set(serviceName, username, accessToken)
}
