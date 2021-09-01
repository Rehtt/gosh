package utils

import (
	"bufio"
	"os/exec"
)

type cmd struct {
	C     *exec.Cmd
	Input *bufio.Writer
}

func Cmd(command string, f func(out string) (exit bool)) (*cmd, error) {
	c := exec.Command(cmdRun, "-c", command)
	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()
	stdin, _ := c.StdinPipe()
	stdoutScanner := bufio.NewScanner(stdout)
	stderrScanner := bufio.NewScanner(stderr)
	go func() {
		go func() {
			for stderrScanner.Scan() {
				if f(stderrScanner.Text()) {
					c.Process.Kill()
				}
			}
		}()
		for stdoutScanner.Scan() {
			if f(stdoutScanner.Text()) {
				c.Process.Kill()
			}
		}
	}()
	return &cmd{c, bufio.NewWriter(stdin)}, c.Start()
}
