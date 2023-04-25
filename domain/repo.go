package domain

import (
	"time"
)

type Repo struct {
	Id           int64
	RepoId       int64
	FullRepoName string
	RepoName     string
	CreateTime   time.Time
	UpdateTime   time.Time
}

type RepoName struct {
	Id     int64
	RepoId int64
	Name   string
}

type RepoWith struct {
	Repo
	RepoNames []RepoName
}
