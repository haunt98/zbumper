package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func ensureReleaseBranch() error {
	// git branch | grep \* | cut -d ' ' -f2
	output, err := exec.
		Command("git","branch",
			"|","grep","\\*",
			"|","cut","-d","'","'","-f2").CombinedOutput()
	if err != nil {
		return err
	}
	msg := string(output)
	if strings.HasPrefix(msg, "releases/") || strings.HasPrefix(msg, "release/") {
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s is not a releasable branch", msg))
	}
}

func commitBump(b *Bump) (string, error) {
	output, err := exec.
		Command("git", "commit",
			"-m", b.composeGitCommitMessage(),
			"--allow-empty").
		CombinedOutput()
	return string(output), err
}
