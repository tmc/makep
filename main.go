// Program makep invokes the nearest Makefile
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	for cwd != "" && cwd != "/" {
		cwd, err = tryMake(cwd)
		if err != nil {
			return err
		}

	}
	return nil
}

func tryMake(wd string) (string, error) {
	p := filepath.Join(wd, "Makefile")
	_, err := os.Stat(p)
	if err == nil {
		mk, err := exec.LookPath("make")
		if err != nil {
			return "", err
		}
		cmd := exec.Command(mk, os.Args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Dir = wd
		return "", cmd.Run()
	} else if os.IsNotExist(err) {
		return filepath.Clean(filepath.Join(wd, "..")), nil
	}
	return "", err
}
