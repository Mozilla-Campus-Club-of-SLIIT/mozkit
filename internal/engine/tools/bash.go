package tools

import (
	"os/exec"
	"strings"
)

func RunBash(args []string) (string, error) {
	cmd := exec.Command("bash", "-c", strings.Join(args, " "))
	output, err := cmd.CombinedOutput()
	return string(output), err
}
