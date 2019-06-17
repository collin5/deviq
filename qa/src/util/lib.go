package util

import (
	"io"
	"lib/pty"
	"os/exec"
)

type Container struct {
	Name string
	Lang string
}

var image = map[string]string{
	"python": "pytest:latest",
}

func NewContainer(name string, lang string) *Container {
	return &Container{name, lang}
}

func (c *Container) Run(code string, spec string) (io.Reader, error) {
	cmd := exec.Command("/usr/local/bin/docker", "run", "--rm", image[c.Lang], code, spec)
	return pty.Start(cmd)
}
