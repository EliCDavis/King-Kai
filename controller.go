package main

import (
	"io"
	"log"
	"os/exec"
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

func (c *controller) left() error {
	_, err := c.stdin.Write([]byte("left\n"))
	log.Println("written")
	return err
}
