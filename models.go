package main

import (
	"fmt"
	"time"
)

type Bump struct {
	Project string
	Service string
	Version string
}

func (b *Bump) composeGitCommitMessage() string {
	return fmt.Sprintf("bump(%s): %s - version v%s", b.Project, b.Service, b.Version)
}

type Repository struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
}

type Tag struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Location string `json:"location"`
}
