package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"

	"golang.org/x/term"
)

func main() {
	if err := cTest(); err != nil {
		fmt.Println(err)
	}
}

func cTest() error {

	if !term.IsTerminal(int(os.Stdin.Fd())) {
		if err := readLines(os.Stdin); err != nil {
			return err
		}
		return nil
	}
	return errors.New("must be pipe input")
}

// reference: https://www.yunabe.jp/tips/golang_readlines.html
func readLines(f *os.File) error {
	s := bufio.NewScanner(f)
	for s.Scan() {
		colorString(s.Text())
	}
	if s.Err() != nil {
		return s.Err()
	}
	return nil
}

var c *color.Color
var (
	success = color.New(color.FgGreen)
	fail    = color.New(color.FgHiRed)
)

func colorString(line string) {
	trimmed := strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(trimmed, "=== RUN"):
		fallthrough
	case strings.HasPrefix(trimmed, "?"):
		c = nil

	// success
	case strings.HasPrefix(trimmed, "--- PASS"):
		fallthrough
	case strings.HasPrefix(trimmed, "ok"):
		fallthrough
	case strings.HasPrefix(trimmed, "PASS"):
		c = success

	// failure
	case strings.HasPrefix(trimmed, "--- FAIL"):
		fallthrough
	case strings.HasPrefix(trimmed, "FAIL"):
		c = fail
	}

	if c == nil {
		fmt.Printf("%s\n", line)
		return
	}
	c.Printf("%s\n", line)
}

func Example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
