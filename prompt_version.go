package main

import (
	"errors"
	"fmt"
	"github.com/blang/semver"
	"github.com/manifoldco/promptui"
	"regexp"
)

const semverRegEx = `^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(-(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\+[0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*)?$`

func promptVersion(project, service string) (string, error) {
	accessToken, err := ensureAccessTokenAvailableAndValid()
	if err != nil {
		return promptInputVersionOrAccessToken(project, service)
	} else {
		return promptInputVersionOrAutoVersion(project, service, accessToken)
	}
}

func promptInputVersionOrAccessToken(project, service string) (string, error) {
	i, result, err := (&promptui.SelectWithAdd{
		Label: "New version?",
		Items: []string{
			"Fetch from Container Registry",
		},
		AddLabel: "Use my own input",
		Validate: validateVersionInput,
	}).Run()
	if err != nil {
		return "", err
	}
	switch i {
	case 0:
		if err := promptAccessToken(); err != nil {
			return "", err
		}
		return promptVersion(project, service)
	case -1:
		return result, nil
	default:
		return "", errors.New("invalid case")
	}
}

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

func promptInputVersionOrAutoVersion(project, service, accessToken string) (string, error) {
	repo, err := getGitlabRepository(project, service, accessToken)
	if err != nil {
		return "", err
	}
	latestVer, err := getLatestVersion(repo.ID, accessToken)
	if err != nil {
		return promptForUserInputOnly()
	} else {
		return promptAsIfLatestVersionValid(latestVer)
	}
}

func promptForUserInputOnly() (string, error) {
	i, result, err := (&promptui.SelectWithAdd{
		Label: "New version?",
		Items: []string{
		},
		AddLabel: "Use my own input",
		Validate: validateVersionInput,
	}).Run()
	if err != nil {
		return "", err
	}
	switch i {
	case -1:
		return result, nil
	default:
		return "", errors.New("invalid case")
	}
}

func promptAsIfLatestVersionValid(latestVer *semver.Version) (string, error) {
	patchVer, err := semver.Make(latestVer.String())
	if err != nil {
		return "", err
	}
	if err := patchVer.IncrementPatch(); err != nil {
		return "", err
	}

	minorVer, err := semver.Make(latestVer.String())
	if err != nil {
		return "", err
	}
	if err := minorVer.IncrementMinor(); err != nil {
		return "", err
	}

	majorVer, err := semver.Make(latestVer.String())
	if err != nil {
		return "", err
	}
	if err := majorVer.IncrementMajor(); err != nil {
		return "", err
	}

	i, result, err := (&promptui.SelectWithAdd{
		Label: fmt.Sprintf("%s is the latest, what is the new one?", latestVer.String()),
		Items: []string{
			fmt.Sprintf("%s - increase patch", patchVer.String()),
			fmt.Sprintf("%s - increase minor", minorVer.String()),
			fmt.Sprintf("%s - increase major", majorVer.String()),
		},
		AddLabel: "Use my own input",
		Validate: validateVersionInput,
	}).Run()
	if err != nil {
		return "", err
	}
	switch i {
	case 0:
		return patchVer.String(), nil
	case 1:
		return minorVer.String(), nil
	case 2:
		return majorVer.String(), nil
	case -1:
		return result, nil
	default:
		return "", errors.New("invalid case")
	}
}

func validateVersionInput(ver string) error {
	match, err := regexp.MatchString(semverRegEx, ver)
	if err != nil {
		return err
	}
	if !match {
		return errors.New("invalid semver")
	}
	return nil
}

func validateAccessToken(accessToken string) error {
	return ensureAccessTokenValid(accessToken)
}
