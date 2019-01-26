package main

type ZCircle string

const (
	QuarterForward ZCircle = "quarterforward"
	QuarterBack    ZCircle = "quarterback"
	HalfForward    ZCircle = "halfforward" // We have to do this cause we never know our orientation -_-
	HalfBack       ZCircle = "halfback"
)

type ZDirection string

const (
	Left    ZDirection = "left"
	Right   ZDirection = "right"
	Down    ZDirection = "down"
	Up      ZDirection = "up"
	Neutral ZDirection = "neutral"
)

type ZAttack string

const (
	Special ZAttack = "special"
	Light   ZAttack = "light"
	Medium  ZAttack = "medium"
	Heavy   ZAttack = "heavy"
)

type ZBar string

const (
	BarOne ZBar = "barone"
	BarTwo ZBar = "bartwo"
)

type ZOtherInput string

const (
	AssistOne  ZOtherInput = "assistone"
	AssistTwo  ZOtherInput = "assisttwo"
	DragonRush ZOtherInput = "dragonrush"
	SuperDash  ZOtherInput = "superdash"
	Vanish     ZOtherInput = "vanish"
	Sparking   ZOtherInput = "sparking"
)

type ZJump string

const (
	JumpUp    ZJump = "up"
	JumpRight ZJump = "right"
	JumpLeft  ZJump = "left"
)

type controller interface {
	attack(ZAttack, ZDirection) error
	specialAttack(ZAttack, ZCircle) error
	bar(ZBar, ZCircle) error
	input(ZOtherInput) error
	jump(ZJump) error
	move(ZDirection) error
	reset() error
}
