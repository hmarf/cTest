package ctest

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var GoTestOptions = []string{"test"}

type COption struct {
	Run  bool
	Pass bool
	Fail bool
}

func CTest(o COption) error {
	cmd := exec.Command("go", GoTestOptions...)
	// standard output
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	// error output
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// run test
	if err := cmd.Start(); err != nil {
		return err
	}

	// output
	scan := bufio.NewScanner(stdout)
	scanErr := bufio.NewScanner(stderr)
	go scanner(scan, o)
	go scanner(scanErr, o)

	// wait to
	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func scanner(r *bufio.Scanner, o COption) {
	for r.Scan() {
		colorString(r.Text(), o)
	}
}

// if you want to change text color, modify this function
func colorString(line string, o COption) {
	trimmed := strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(trimmed, "=== PAUSE"):
		fallthrough
	case strings.HasPrefix(trimmed, "=== RUN"):
		fallthrough
	case strings.HasPrefix(trimmed, "?"):
		if !o.Run {
			fmt.Println(line)
		}

	case strings.HasPrefix(trimmed, "ok"):
		fallthrough
	case strings.HasPrefix(trimmed, "PASS"):
		fallthrough
	case strings.HasPrefix(trimmed, "--- PASS"):
		if !o.Pass {
			color.Green(line)
		}

	case strings.HasPrefix(trimmed, "--- FAIL"):
		fallthrough
	case strings.HasPrefix(trimmed, "FAIL"):
		if !o.Fail {
			color.Red(line)
		}
	}
}
