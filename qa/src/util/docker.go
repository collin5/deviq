package util

import (
	"io"
	"os/exec"

	"github.com/kr/pty"
)

type Container struct {
	name string
	lang string
}

var image = map[string]string{
	"python": "pytest:latest",
}

func NewContainer(name string, lang string) *Container {
	return &Container{name, lang}
}

func (c *Container) Run(code string, tests string) (io.Reader, error) {
	cmd := exec.Command("/usr/local/bin/docker", "run", image[c.lang], code, tests)
	return pty.Start(cmd)
}
