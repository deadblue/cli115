package terminal

type Context interface {

	// Is context alive
	Alive() bool
	// Prompt text
	Prompt() string
	// Startup context
	Startup() error
	// Shutdown context
	Shutdown() error
}
