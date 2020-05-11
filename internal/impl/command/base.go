package command

type NoArgsCommand struct{}

func (c *NoArgsCommand) HasArgs() bool {
	return false
}

type ArgsCommand struct{}

func (c *ArgsCommand) HasArgs() bool {
	return true
}
