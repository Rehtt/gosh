package utils

import (
	"fmt"
	"os"
)

var cmdRun string

func init() {
	if _, err := os.Open("C:\\Program Files\\PowerShell\\7\\pwsh.exe"); err == nil {
		cmdRun = "pwsh"
	} else if _, err = os.Open("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"); err == nil {
		cmdRun = "powershell"
	} else {
		fmt.Println("not run")
		os.Exit(1)
	}
}
