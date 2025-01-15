package sources

import (
	"errors"
	"io"
	"os/exec"
	"strings"
)

type CommandSource struct {
	command string
	args    []string
}

func (c *CommandSource) Configure(config map[string]string) error {
	cmd, ok := config["command"]
	if !ok {
		return errors.New("command configuration required")
	}
	c.command = cmd

	// Optional args as comma separated string
	if args, ok := config["args"]; ok {
		c.args = strings.Split(args, ",")
	}

	return nil
}

func (c *CommandSource) Print(writer io.Writer) {
	cmd := exec.Command(c.command, c.args...)
	cmd.Stdout = writer
	err := cmd.Run()
	if err != nil {
		writer.Write([]byte("Error running command: " + err.Error()))
	}
}
