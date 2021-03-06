package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
)

type pyController struct {
	cmd    *exec.Cmd
	stdin  io.WriteCloser
	stdout io.ReadCloser
}

func newPyController() (*pyController, error) {
	cmd := exec.Command("python", ".\\controller.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	stout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &pyController{cmd, stdin, stout}, nil
}

func (c *pyController) everything() string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.stdout)
	return buf.String()
}

func (c *pyController) control(command string, duration int) error {
	instructions := fmt.Sprintf("%d %s\n", duration, command)
	_, err := c.stdin.Write([]byte(instructions))
	log.Printf("sent: %s", instructions)
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
	return c.controlWithRest(fmt.Sprintf("special %s-%s", attack, direction), 280)
}

func (c *pyController) bar(attack ZBar, direction ZCircle) error {
	return c.controlWithRest(fmt.Sprintf("bar %s-%s", attack, direction), 280)
}

func (c *pyController) command(attack ZCommand) error {
	delay := 100

	switch attack {
	case SuperDash:
		delay = 200
		break

	case DragonRush:
		delay = 1500
		break
	}

	return c.controlWithRest(fmt.Sprintf("command %s", attack), delay)
}

func (c *pyController) jump(attack ZJump) error {
	return c.controlWithRest(fmt.Sprintf("jump %s", attack), 100)
}

func (c *pyController) move(attack ZDirection) error {
	return c.controlWithRest(fmt.Sprintf("move %s", attack), 100)
}

func (c *pyController) reset() error {
	return c.controlWithRest("reset", 500)
}
