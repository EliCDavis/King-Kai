package main

type MoveInput struct {
	move ZDirection
}

func NewMoveInput(move ZDirection) *MoveInput {
	return &MoveInput{move}
}

func (a MoveInput) Execute(c Controller) error {
	return c.move(a.move)
}

func (a MoveInput) Hash() string {
	return string(a.move)
}
