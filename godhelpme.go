package main

// God help me
var allPosibleActions []Input = []Input{
	NewAttackInput(Light, Down),
	NewAttackInput(Light, Up),
	NewAttackInput(Light, Neutral),

	NewAttackInput(Medium, Down),
	NewAttackInput(Medium, Right),
	NewAttackInput(Medium, Up),
	NewAttackInput(Medium, Left),
	NewAttackInput(Medium, Neutral),

	NewAttackInput(Heavy, Down),
	NewAttackInput(Heavy, Neutral),

	NewAttackInput(Special, Down),
	NewAttackInput(Special, Right),
	NewAttackInput(Special, Up),
	NewAttackInput(Special, Left),
	NewAttackInput(Special, Neutral),

	NewSpecialAttackInput(Light, QuarterForward),
	NewSpecialAttackInput(Light, QuarterBack),

	NewSpecialAttackInput(Medium, QuarterForward),
	NewSpecialAttackInput(Medium, QuarterBack),

	NewSpecialAttackInput(Heavy, QuarterForward),
	NewSpecialAttackInput(Heavy, QuarterBack),

	NewSpecialAttackInput(Special, QuarterForward),
	NewSpecialAttackInput(Special, QuarterBack),

	NewJumpInput(JumpUp),
	NewJumpInput(JumpLeft),
	NewJumpInput(JumpRight),

	NewMoveInput(Down),
	NewMoveInput(Left),
	NewMoveInput(Right),

	NewBarInput(BarOne, QuarterForward),
	NewBarInput(BarOne, QuarterBack),
	NewBarInput(BarOne, HalfForward),
	NewBarInput(BarOne, HalfBack),

	NewBarInput(BarTwo, QuarterForward),
	NewBarInput(BarTwo, QuarterBack),

	NewCommandInput(AssistOne),
	NewCommandInput(AssistTwo),
	NewCommandInput(DragonRush),
	NewCommandInput(SuperDash),
	NewCommandInput(Vanish),
	NewCommandInput(Sparking),
}
