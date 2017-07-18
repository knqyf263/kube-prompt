package kube

import (
	"strings"
	"os/exec"
)

func Executor(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}

	args := strings.Split(s, " ")
	cmd := exec.Command("kubectl", args...)
	out, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(out)
}