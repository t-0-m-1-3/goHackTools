package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"time"

	"log"

	"github.com/shavac/gexpect"
)

func main() {
	// where is ssh
	ssh, err := exec.LookPath("ssh")
	if err != nil {
		log.Println(err)
	}

	// new connection
	child, _ := gexpect.NewSubProcess(ssh, "user@127.0.0.1")
	if err := child.Start(); err != nil {
		fmt.Println(err)
	}
	// close the connection
	defer child.Close()
	// password input
	if idx, _ := child.ExpectTimeout(0*time.Second, regexp.MustCompile("password:")); idx >= 0 {
		child.SendLine("pass")
	}

	// command input
	child.SendLine("sudo cat /etc/shadow | grep root")

	// время ожидания
	child.InteractTimeout(3 * time.Second)
}
