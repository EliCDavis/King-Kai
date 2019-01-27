package main

type JumpInput struct {
	jump ZJump
}

func NewJumpInput(jump ZJump) *JumpInput {
	return &JumpInput{jump}
}

func (a JumpInput) Execute(c Controller) error {
	return c.jump(a.jump)
}

func (a JumpInput) Hash() string {
	return string(a.jump)
}
