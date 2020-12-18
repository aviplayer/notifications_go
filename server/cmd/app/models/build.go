package models

type BuildStatus int

const (
	Success = iota
	Fail
	Unstable
)

type Build struct {
	Name       string
	Version    string
	Status     BuildStatus
	LastCommit string
}
