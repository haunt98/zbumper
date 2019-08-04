package main

import (
	"errors"
	"github.com/manifoldco/promptui"
)

func promptAction(bump *Bump) error {
	i, _, err := (&promptui.Select{
		Label: "What do you want?",
		Items: []string{
			"Build image and push",
			"Bump release",
		},
		Size: 10,
	}).Run()
	if err != nil {
		return err
	}
	switch i {
	case 0:
		return doBumpRelease(bump)
	case 1:
		return doBuildAndPush(bump)
	default:
		return errors.New("invalid case")
	}
}

func doBumpRelease(bump *Bump) error {
	return ensureBump(bump)
}

func doBuildAndPush(bump *Bump) error {
	return ensureBump(bump)
}

func ensureBump(bump *Bump) error {
	var err error
	if len(bump.Project) == 0 {
		bump.Project, err = promptProjectName()
	}
	if err != nil {
		return err
	}
	if len(bump.Service) == 0 {
		bump.Service, err = promptServiceName(bump.Project)
	}
	if err != nil {
		return err
	}
	if len(bump.Tag) == 0 {
		bump.Tag, err = promptTag(bump.Project, bump.Service)
	}
	if err != nil {
		return err
	}
	return nil
}
