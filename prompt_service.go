package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"strings"
)

func promptServiceName(project string) (string, error) {
	services, err := getServices(project)
	if err != nil {
		return "", err
	}
	prompt := promptui.Select{
		Label: "Select Service",
		Items: services,
		Size:  10,
	}

	prompt.StartInSearchMode = true
	prompt.Searcher = func(input string, index int) bool {
		return strings.HasPrefix(services[index], input)
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getServices(project string) ([]string, error) {
	files, err := ioutil.ReadDir(fmt.Sprintf("./%s/cmd/", project))
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, f := range files {
		if f.IsDir() {
			result = append(result, f.Name())
		}
	}
	return result, nil
}
