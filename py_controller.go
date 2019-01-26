package main

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

type pyController struct {
	cmd   *exec.Cmd
	stdin io.WriteCloser
}

func newPyController() (*pyController, error) {
	cmd := exec.Command("python", ".\\controller.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &pyController{cmd, stdin}, nil
}

func (c *pyController) control(command string, duration int) error {
	instructions := fmt.Sprintf("%d %s\n", duration, command)
	_, err := c.stdin.Write([]byte(instructions))
	//log.Printf("sent: %s", instructions)
	return err
}

func (c *pyController) controlWithRest(command string, duration int) error {
	err := c.control(command, duration)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(duration) * time.Millisecond)
	return nil
}

func (c *pyController) attack(attack ZAttack, direction ZDirection) error {
	return c.controlWithRest(fmt.Sprintf("attack %s-%s", attack, direction), 200)
}

func (c *pyController) specialAttack(attack ZAttack, direction ZCircle) error {
	return c.controlWithRest(fmt.Sprintf("special %s-%s", attack, direction), 100)
}

func (c *pyController) bar(attack ZBar, direction ZCircle) error {
	return c.controlWithRest(fmt.Sprintf("bar %s-%s", attack, direction), 100)
}

func (c *pyController) input(attack ZOtherInput) error {
	return c.controlWithRest(fmt.Sprintf("other %s", attack), 100)
}

func (c *pyController) jump(attack ZJump) error {
	return c.controlWithRest(fmt.Sprintf("jump %s", attack), 100)
}

func (c *pyController) move(attack ZDirection) error {
	return c.controlWithRest(fmt.Sprintf("move %s", attack), 100)
}

func (c *pyController) reset() error {
	return c.controlWithRest("reset", 100)
}
