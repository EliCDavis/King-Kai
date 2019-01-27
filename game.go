package main

type Game interface {
	GetDamage() (int, error)
}
