package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jniltinho/go-pwgen"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "pwgen"
	app.Version = "0.3.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "jniltinho",
			Email: "jniltinho@gmail.com",
		},
	}
	app.Copyright = "(c) 2019 jniltinho"
	app.Usage = "`pwgen` password generator by Golang 1.12"
	app.UsageText = "pwgen [options]"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "length, l",
			Usage: "password length (range is [8...64])",
			Value: 10,
		},
		cli.IntFlag{
			Name:  "count, c",
			Usage: "number of counts",
			Value: 10,
		},
		cli.BoolTFlag{
			Name:  "digit",
			Usage: "add password characters 0-9",
		},
		cli.BoolTFlag{
			Name:  "alphabetlarge",
			Usage: "add password characters A-Z",
		},
		cli.BoolTFlag{
			Name:  "alphabetsmall",
			Usage: "add password characters a-z",
		},
		cli.BoolFlag{
			Name:  "underscore, u",
			Usage: "add password character Underscore(_)",
		},
		cli.BoolFlag{
			Name:  "specialchars, s",
			Usage: "add password characters special characters, exclude from Space, Backslash, Underscore and Delete",
		},
	}

	app.Action = func(c *cli.Context) error {
		writer := os.Stdout
		err := pwgen.Pwgen(
			writer,
			c.Int("length"),
			c.Int("count"),
			c.Bool("digit"),
			c.Bool("alphabetlarge"),
			c.Bool("alphabetsmall"),
			c.Bool("underscore"),
			c.Bool("specialchars"),
		)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		return nil
	}

	app.Run(os.Args)
}
