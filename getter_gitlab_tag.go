package main

import (
	"errors"
	"github.com/blang/semver"
	"strings"
)

func getLatestVersion(repositoryID uint64, accessToken string) (*semver.Version, error) {
	tags, err := getRepositoryTags(repositoryID, accessToken)
	if err != nil {
		return nil, err
	}
	if len(tags) == 0 {
		return nil, errors.New("no tag found")
	}

	versions := make([]semver.Version, 0)
	for _, t := range tags {
		v, err := semver.Make(strings.TrimPrefix(t.Name, "v"))
		if err != nil {
			continue
		}
		versions = append(versions, v)
	}
	semver.Sort(versions)
	return &(versions[len(versions)-1]), nil
}
