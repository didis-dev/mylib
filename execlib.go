package mylib

import (
	"fmt"
	"os/exec"
)

func Myexec(a, b int) int {
	cmd := exec.Command("whoami")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		fmt.Println("a=", a, "b=", b)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
	return a
}
