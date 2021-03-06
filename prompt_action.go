package main

import (
	"errors"
	"github.com/manifoldco/promptui"
	"log"
)

func promptAction(bump *Bump, accessToken string) error {
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
		return doBuildAndPush(bump, accessToken)
	case 1:
		return doBumpRelease(bump)
	default:
		return errors.New("invalid case")
	}
}

func doBumpRelease(bump *Bump) error {
	if err := ensureReleaseBranch(); err != nil {
		return err
	}
	if err := ensureBump(bump); err != nil {
		return err
	}
	msg, err := commitBump(bump)
	log.Println(msg)
	return err
}

func doBuildAndPush(bump *Bump, accessToken string) error {
	if err := ensureBump(bump); err != nil {
		return err
	}
	msg, err := buildAndDeployDocker(bump, accessToken)
	log.Println(msg)
	return err
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
	if len(bump.Version) == 0 {
		bump.Version, err = promptVersion(bump.Project, bump.Service)
	}
	if err != nil {
		return err
	}
	return nil
}
