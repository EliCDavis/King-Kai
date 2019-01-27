package main

type Input interface {
	Execute(Controller) error
	Hash() string
}
