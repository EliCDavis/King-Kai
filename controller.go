package main

import (
	"fmt"
	"io"
	"os/exec"
	"time"
)

type controller struct {
	cmd   *exec.Cmd
	stdin io.WriteCloser
}

func newController() (*controller, error) {
	cmd := exec.Command("python", ".\\controller.py")

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return &controller{cmd, stdin}, nil
}

func (c *controller) control(command string, duration int) error {
	instructions := fmt.Sprintf("%s %d\n", command, duration)
	_, err := c.stdin.Write([]byte(instructions))
	//log.Printf("sent: %s", instructions)
	return err
}

func (c *controller) controlWithRest(command string, duration int, rest int) error {
	err := c.control(command, duration)
	if err != nil {
		return err
	}
	time.Sleep(time.Duration(rest) * time.Millisecond)
	return nil
}

func (c *controller) left() error {
	return c.control("left", 10)
}

func (c *controller) right() error {
	return c.control("right", 10)
}

func (c *controller) down() error {
	return c.control("down", 10)
}

func (c *controller) up() error {
	return c.control("up", 10)
}

func (c *controller) special() error {
	return c.controlWithRest("special", 100, 100)
}

func (c *controller) light() error {
	return c.controlWithRest("light", 100, 100)
}

func (c *controller) medium() error {
	return c.controlWithRest("medium", 100, 100)
}

func (c *controller) heavy() error {
	return c.controlWithRest("heavy", 100, 100)
}
