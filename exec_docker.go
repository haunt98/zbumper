package main

import (
	"bufio"
	"errors"
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
	cmdStr := strings.Join([]string{
		"docker", "build", "-t", imageTag, "-f", dockerfileName, ".", "&&",
		"docker", "login", dockerHost, "&&",
		"docker", "push", imageTag},
		" ")


	cmd := exec.Command("bash", "-c", cmdStr)


	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()
	if err := cmd.Start(); err != nil {
		err = errors.New("cannot start command: " + err.Error())
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		err = errors.New("cannot wait command: " + err.Error())
		return "", err
	}
	return string("success"), err
}
