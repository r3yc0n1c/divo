package util

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("powershell", "-Command", "Clear-Host")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
	}
}
