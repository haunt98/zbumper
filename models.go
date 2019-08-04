package main

import "time"

type Bump struct {
	Project string
	Service string
	Tag     string
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
