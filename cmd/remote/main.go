package main

import (
	"fmt"
	"os"
	"os/exec"

	PID "github.com/Ribas160/ayanotAnonynousBot/pkg/pid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("[ERROR] Command has not been specified")
		return
	}

	command := os.Args[1]

	switch command {

	case "start":
		err := start()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case "stop":
		err := stop()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	default:
		fmt.Println("Unknown command")
		return
	}
}

func start() error {
	cmd := exec.Command("./bin/ayanotAnonymousBot")
	if err := cmd.Start(); err != nil {
		return err
	}

	if err := PID.Write(cmd.Process.Pid); err != nil {
		proc, _ := os.FindProcess(cmd.Process.Pid)
		proc.Kill()

		fmt.Printf("Unable to write pid file, proccess %d was killed\n", cmd.Process.Pid)

		return err
	}

	return nil
}

func stop() error {
	pid, err := PID.Read()
	if err != nil {
		fmt.Println("Bot is not running")
		return nil
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	if err := PID.Delete(); err != nil {
		return err
	}

	proc.Kill()

	return nil
}
