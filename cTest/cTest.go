package cTest

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

var GoTestOptions = []string{"test"}

type COption struct {
	Run          bool
	Pass         bool
	Fail         bool
	PassColor    string
	FailColor    string
	passColorFmt string
	failColorFmt string
}

func createPassColorFmt(c string) string {
	return "\x1b[" + c + "m%s\x1b[0m"
}

func colorIdentification(color string, kind string) string {
	switch color {
	case "black":
		return createPassColorFmt("30")
	case "red":
		return createPassColorFmt("31")
	case "green":
		return createPassColorFmt("32")
	case "yellow":
		return createPassColorFmt("33")
	case "blue":
		return createPassColorFmt("34")
	case "magenta":
		return createPassColorFmt("35")
	case "cyan":
		return createPassColorFmt("36")
	case "white":
		return createPassColorFmt("37")
	}
	if kind == "pass" {
		return createPassColorFmt("32") // pass default color: green
	} else if kind == "fail" {
		return createPassColorFmt("31") // fail default color: red
	}
	return createPassColorFmt("30")
}

func CTest(o COption) error {
	o.passColorFmt = colorIdentification(o.PassColor, "pass")
	o.failColorFmt = colorIdentification(o.FailColor, "fail")
	fmt.Println(o.PassColor, o.FailColor)
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
		return nil
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
	default:
		fmt.Println(line)
	}
}
