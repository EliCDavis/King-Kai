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
	BarOne ZBar = "one"
	BarTwo ZBar = "two"
)

type ZCommand string

const (
	AssistOne  ZCommand = "assistone"
	AssistTwo  ZCommand = "assisttwo"
	DragonRush ZCommand = "dragonrush"
	SuperDash  ZCommand = "superdash"
	Vanish     ZCommand = "vanish"
	Sparking   ZCommand = "sparking"
)

type ZJump string

const (
	JumpUp    ZJump = "up"
	JumpRight ZJump = "right"
	JumpLeft  ZJump = "left"
)

type Controller interface {
	attack(ZAttack, ZDirection) error
	specialAttack(ZAttack, ZCircle) error
	bar(ZBar, ZCircle) error
	command(ZCommand) error
	jump(ZJump) error
	move(ZDirection) error
	everything() string
	reset() error
}
