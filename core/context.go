package core

type Context interface {
	Prompt() string

	Alive() bool
}
