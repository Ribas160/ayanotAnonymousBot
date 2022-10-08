package pid

import (
	"errors"
	"os"
	"strconv"
)

const pidFile = "run/ayanotAnonymousBot.pid"

func Write(pid int) error {
	oldPid, _ := Read()

	if oldPid != 0 {
		return errors.New("Bot has already been running")
	}

	fp, err := os.OpenFile(pidFile, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}

	defer fp.Close()

	fp.WriteString(strconv.Itoa(pid))

	return nil
}

func Read() (int, error) {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return pid, nil
}

func Delete() error {
	err := os.Remove(pidFile)
	if err != nil {
		return err
	}

	return nil
}
