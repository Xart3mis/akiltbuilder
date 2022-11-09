package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type CommandAttrs struct {
  Starter func() error
  Waiter func() error
}

func CloneAKILT() {
	exec.Command("git", "clone", "https://github.com/Xart3mis/AKILT").Run()
}

func BuildwithServerIP(server_ip string) (*bufio.Scanner, CommandAttrs) {
	exec.Command("cd", "AKILT/").Run()
	cmd := exec.Command("go", "build", "-v", "-x", "--ldflags", fmt.Sprintf("\"-X main.ServerIp=%s\"", server_ip), "Client/")

  cmd.Stdout = os.Stdout 

	pipe, err := cmd.StdoutPipe()
  if err != nil {
    fmt.Println(err)
  }
    
	scanner := bufio.NewScanner(pipe)

	return scanner, CommandAttrs{Starter: cmd.Run, Waiter: cmd.Wait}
}
