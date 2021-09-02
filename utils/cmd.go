package utils

import (
	"bufio"
	"os/exec"
)

type cmd struct {
	c *exec.Cmd
}

func Cmd(command string, f func(out string) (exit bool)) (*cmd, error) {
	c := exec.Command(cmdRun, "-c", command)
	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	stdoutScanner := bufio.NewScanner(stdout)
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		for stderrScanner.Scan() {
			if f(stderrScanner.Text()) {
				c.Process.Kill()
			}
		}
	}()
	go func() {
		for stdoutScanner.Scan() {
			if f(stdoutScanner.Text()) {
				c.Process.Kill()
			}
		}
	}()
	return &cmd{c}, c.Start()
}

func (c cmd) Wait() error {
	return c.c.Wait()
}
