package main

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
)

func main() {
	err := KillApp("server.pid")
	if err != nil {
		fmt.Println("err:", err)

		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("not found")
		}

		for e := err; e != nil; e = errors.Unwrap(e) {
			fmt.Printf("> %s\n", e)
		}
	}
}

func KillApp(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			slog.Warn("Failed to close file", "file", fileName, "err", err)
		}
	}()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return fmt.Errorf("%q - bad pid: %w", fileName, err)
	}

	slog.Info("killing", "pid", pid)
	if err := os.Remove(fileName); err != nil {
		slog.Warn("Failed to kill pid", "pid", pid, "err", err)
	}

	return nil
}
