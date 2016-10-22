package utils

import (
	"bytes"
	"os/exec"
)

type commandError struct {
	message string
	err     error
}

func (e *commandError) Error() string {
	return e.message
}

type Command struct {
	Binary string
	Args   []string
}

type CommandUtils struct {
	output bytes.Buffer
}

func (utils *CommandUtils) CommandOutput() string {
	return utils.output.String()
}

func (utils *CommandUtils) RunCommand(c Command) (bool, error) {
	stringUtils := StringUtils{}

	path, err := exec.LookPath(c.Binary)
	if err != nil {
		return false, &commandError{stringUtils.StringConcat([]string{"[", c.Binary, "] command not found"}), err}
	}

	cmd := exec.Command(path, stringUtils.StringConcat(c.Args))
	cmd.Stdout = &utils.output

	err = cmd.Start()
	err = cmd.Wait()
	if err != nil {
		return false, &commandError{utils.output.String(), err}
	}
	return true, nil
}
