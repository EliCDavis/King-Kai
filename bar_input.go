package main

import "fmt"

type BarInput struct {
	bar    ZBar
	circle ZCircle
}

func NewBarInput(bar ZBar, circle ZCircle) *BarInput {
	return &BarInput{bar, circle}
}

func (a BarInput) Execute(c Controller) error {
	return c.bar(a.bar, a.circle)
}

func (a BarInput) Hash() string {
	return fmt.Sprintf("%s-%s", a.bar, a.circle)
}
