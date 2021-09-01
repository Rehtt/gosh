package utils

import (
	"bufio"
	"os/exec"
)

type com struct {
	c             *exec.Cmd
	stdoutScanner *bufio.Scanner
	stderrScanner *bufio.Scanner
}

func Cmd(cmd string) (com2 *com) {
	com2 = &com{}
	com2.c = exec.Command(cmdRun, "-c", cmd)
	stdout, _ := com2.c.StdoutPipe()
	stderr, _ := com2.c.StderrPipe()
	com2.stdoutScanner = bufio.NewScanner(stdout)
	com2.stderrScanner = bufio.NewScanner(stderr)
	return
}
func (c *com) Run() error {
	return c.c.Wait()
}
func (c *com) OutPut(f func(out string) (exit bool)) error {
	go func() {
		go func() {
			for c.stderrScanner.Scan() {
				if f(c.stderrScanner.Text()) {
					c.c.Process.Kill()
				}
			}
		}()
		for c.stdoutScanner.Scan() {
			if f(c.stdoutScanner.Text()) {
				c.c.Process.Kill()
			}
		}
	}()
	return c.c.Start()
}
