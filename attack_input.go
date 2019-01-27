package main

import "fmt"

type AttackInput struct {
	attack    ZAttack
	direction ZDirection
}

func NewAttackInput(attack ZAttack, direction ZDirection) *AttackInput {
	return &AttackInput{attack, direction}
}

func (a AttackInput) Execute(c Controller) error {
	return c.attack(a.attack, a.direction)
}

func (a AttackInput) Hash() string {
	return fmt.Sprintf("%s-%s", a.attack, a.direction)
}
