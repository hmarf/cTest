package main

import (
	"fmt"
	"os"

	"github.com/hmarf/ctest/cTest"
	"github.com/urfave/cli"
)

func App() *cli.App {
	app := cli.NewApp()
	app.Name = "cTest"
	app.UsageText = "ctest [global options] [go test options]"
	app.Usage = "Give color to the output according to the test result."
	app.Version = "0.1.0"
	app.Author = "hmarf"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "r", // run
			Usage: "Do not output '=== RUN: ~'",
		},
		cli.BoolFlag{
			Name:  "p", // pass
			Usage: "Do not output '--- PASS: ~'",
		},
		cli.BoolFlag{
			Name:  "f", // fail
			Usage: "Do not output '--- FAIL: ~'",
		},
		cli.StringFlag{
			Name:  "passColor", // pass color
			Value: "green",
			Usage: "Change '--- PASS: ~' color. Color list:[black, white, green, yellow, blue, magenta, cyan, white]",
		},
		cli.StringFlag{
			Name:  "failColor", // fail color
			Value: "red",
			Usage: "Change '--- FAIL: ~' color. Color list:[black, white, green, yellow, blue, magenta, cyan, white]",
		},
	}
	return app
}

func Action(c *cli.Context) {
	co := cTest.COption{
		Run:       c.Bool("r"), // run
		Pass:      c.Bool("p"), // pass
		Fail:      c.Bool("f"), // fail
		PassColor: c.String("passColor"),
		FailColor: c.String("failColor"),
	}
	if err := cTest.CTest(co); err != nil {
		fmt.Println(err)
	}
}

func getGoTestOption(args []string) []string {
	all_ctestOptions := []string{"-r", "-p", "-f", "-h", "--help", "help"}
	ctestOption := []string{args[0]}
	for _, arg := range args[1:] {
		opTF := true
		for _, cOption := range all_ctestOptions {
			if arg == cOption {
				ctestOption = append(ctestOption, arg)
				opTF = false
				break
			}
		}
		if opTF {
			cTest.GoTestOptions = append(cTest.GoTestOptions, arg)
		}
	}
	return ctestOption
}

func main() {
	app := App()
	app.Action = Action
	app.HideVersion = true
	app.Run(getGoTestOption(os.Args))
}
