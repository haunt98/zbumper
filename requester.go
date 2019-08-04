package main

import (
	"errors"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"os"
)

const baseUrlEnv = "ZBUMPER_GITLAB_BASEURL"
const projectIDEnv = "ZBUMPER_PROJECT_ID"
const tokenHeader = "PRIVATE-TOKEN"

func getBaseUrl() string {
	return os.Getenv(baseUrlEnv)
}

func getProjectID() string {
	return os.Getenv(projectIDEnv)
}

func getRepositoriesUrl() string {
	return fmt.Sprintf("%s/api/v4/projects/%s/registry/repositories", getBaseUrl(), getProjectID())
}

func getRepositoryTagsUrl(repositoryID uint64) string {
	return fmt.Sprintf("%s/%d/tags", getRepositoriesUrl(), repositoryID)
}

func getRepositories(accessToken string) ([]*Repository, error) {
	var result []*Repository
	rsp, _, err := gorequest.New().Get(getRepositoriesUrl()).
		Set(tokenHeader, accessToken).
		EndStruct(&result)
	if len(err) != 0 {
		return nil, err[0]
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d - %s", rsp.StatusCode, http.StatusText(rsp.StatusCode)))
	}
	return result, nil
}

func getRepositoryTags(repositoryID uint64, accessToken string) ([]*Tag, error) {
	var result []*Tag
	rsp, _, err := gorequest.New().Get(getRepositoryTagsUrl(repositoryID)).
		Set(tokenHeader, accessToken).
		EndStruct(&result)
	if len(err) != 0 {
		return nil, err[0]
	}
	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("%d - %s", rsp.StatusCode, http.StatusText(rsp.StatusCode)))
	}
	return result, nil
}
