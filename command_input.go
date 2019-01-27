package main

type CommandInput struct {
	command ZCommand
}

func NewCommandInput(command ZCommand) *CommandInput {
	return &CommandInput{command}
}

func (a CommandInput) Execute(c Controller) error {
	return c.command(a.command)
}

func (a CommandInput) Hash() string {
	return string(a.command)
}
