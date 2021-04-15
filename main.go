package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
	"golang.org/x/term"
)

type Option struct {
	Run  bool
	Pass bool
	Fail bool
}

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "cTest"
	app.Usage = "Give color to the output according to the test result."
	app.Version = "0.0.1"
	app.Author = "hmarf"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "run, r",
			Usage: "Do not output '=== RUN: ~'",
		},
		cli.BoolFlag{
			Name:  "pass, p",
			Usage: "Do not output '--- PASS: ~'",
		},
		cli.BoolFlag{
			Name:  "fail, f",
			Usage: "Do not output '--- FAIL: ~'",
		},
	}
	return app
}

func Action(c *cli.Context) {
	app := App()
	if c.String("input") == "None" {
		app.Run(os.Args)
		return
	}
	option := Option{
		Run:  c.Bool("run"),
		Pass: c.Bool("pass"),
		Fail: c.Bool("fail"),
	}
	if err := cTest(option); err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := App()
	app.Action = Action
	app.Run(os.Args)
}

func cTest(o Option) error {
	if !term.IsTerminal(int(os.Stdin.Fd())) {
		if err := readLines(os.Stdin, o); err != nil {
			return err
		}
		return nil
	}
	return errors.New("must be pipe input")
}

// reference: https://www.yunabe.jp/tips/golang_readlines.html
func readLines(f *os.File, o Option) error {
	s := bufio.NewScanner(f)
	for s.Scan() {
		colorString(s.Text(), o)
	}
	if s.Err() != nil {
		return s.Err()
	}
	return nil
}

// if you want to change text color, modify this function
func colorString(line string, o Option) {
	trimmed := strings.TrimSpace(line)
	switch {
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
