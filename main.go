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

// if you want to change text color, modify this function
func colorString(line string) {
	trimmed := strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(trimmed, "=== RUN"):
		fallthrough
	case strings.HasPrefix(trimmed, "?"):
		fmt.Println(line)

	case strings.HasPrefix(trimmed, "ok"):
		fallthrough
	case strings.HasPrefix(trimmed, "PASS"):
		fallthrough
	case strings.HasPrefix(trimmed, "--- PASS"):
		color.Green(line)

	case strings.HasPrefix(trimmed, "--- FAIL"):
		fallthrough
	case strings.HasPrefix(trimmed, "FAIL"):
		color.Red(line)
	}
}
