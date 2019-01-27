package main

import "fmt"

type SpecialAttackInput struct {
	attack ZAttack
	circle ZCircle
}

func NewSpecialAttackInput(attack ZAttack, circle ZCircle) *SpecialAttackInput {
	return &SpecialAttackInput{attack, circle}
}

func (a SpecialAttackInput) Execute(c Controller) error {
	return c.specialAttack(a.attack, a.circle)
}

func (a SpecialAttackInput) Hash() string {
	return fmt.Sprintf("%s-%s", a.attack, a.circle)
}
