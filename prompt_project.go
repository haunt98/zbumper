package main

import (
	"github.com/manifoldco/promptui"
	"io/ioutil"
	"strings"
)

var ignoreDirs = []string{
	"docs",
	"dynomite",
	"examples",
	"pkg",
	"project_template",
	"referralcode",
	"setup",
	"tools",
	"vendor",
}

func promptProjectName() (string, error) {
	projects, err := getProjects()
	if err != nil {
		return "", err
	}
	prompt := promptui.Select{
		Label: "Select Project",
		Items: projects,
		Size:  10,
	}

	prompt.StartInSearchMode = true
	prompt.Searcher = func(input string, index int) bool {
		return strings.HasPrefix(projects[index], input)
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}

func getProjects() ([]string, error) {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for _, f := range files {
		if f.IsDir() &&
			!strings.HasPrefix(f.Name(), ".") &&
			!shouldBeIgnored(f.Name()) {
			result = append(result, f.Name())
		}
	}
	return result, nil
}

func shouldBeIgnored(project string) bool {
	for _, d := range ignoreDirs {
		if project == d {
			return true
		}
	}
	return false
}
