package main

import (
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	prefixer "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/urfave/cli.v1"
)

var log = logrus.New()
var homedir string

func init() {
	var err error
	log.Formatter = new(prefixer.TextFormatter)
	homedir, err = os.Getwd()
	check(err)
}

func main() {
	app := cli.NewApp()
	app.Name = "Composer - MIT"
	app.Version = "0.1.0"
	app.Author = "mdouchement"
	app.Usage = "Usage:"
	app.Flags = globalFlags
	app.Before = beforeAction

	commands(app)

	err := app.Run(os.Args)
	check(err)
}

var globalFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "D, verbose",
		Usage: "Increase logger level",
	},
}

func beforeAction(context *cli.Context) error {
	if context.Bool("D") || os.Getenv("APP_DEBUG") == "1" {
		log.Level = logrus.DebugLevel
	}
	return nil
}

func check(err error) {
	if err != nil {
		fail(err.Error())
	}
}
func fail(err string) {
	log.Error(err)
	time.Sleep(100 * time.Millisecond) // Wait logger outputing
	os.Exit(1)
}
