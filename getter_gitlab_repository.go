package main

import (
	"errors"
	"fmt"
)

func getGitlabRepository(project, service, accessToken string) (*Repository, error) {
	repos, err := getRepositories(accessToken)
	if err != nil {
		return nil, err
	}
	toCompare := fmt.Sprintf("%s/%s", project, service)
	for _, repo := range repos {
		if repo.Name == toCompare {
			return repo, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("repository not found %s", toCompare))
}
