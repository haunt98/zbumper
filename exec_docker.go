package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func buildAndDeployDocker(b *Bump, accessToken string) (string, error) {
	repo, err := getGitlabRepository(b.Project, b.Service, accessToken)
	if err != nil {
		return "", err
	}
	imageTag := fmt.Sprintf("%s:v%s", repo.Location, b.Version)
	dockerfileName := fmt.Sprintf("%s/docker/Dockerfile_%s", b.Project, b.Service)
	dockerHost := strings.Split(repo.Location, "/")[0]
	output, err := exec.Command(
		"docker", "build", "-t", imageTag, "-f", dockerfileName, ".", "&&",
		"docker", "login", dockerHost, "&&",
		"docker", "push", imageTag).
		CombinedOutput()
	return string(output), err
}
