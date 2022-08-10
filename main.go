package main

import (
	"errors"
	"os"

	"github.com/mizumoto-cn/goutils/file"
	"github.com/mizumoto-cn/goutils/text"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Rename",
		Usage: "Semi-automatic renaming of files",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "The source file or folder",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "The destination file or folder",
			},
			&cli.StringFlag{
				Name:    "type",
				Aliases: []string{"t"},
				Usage:   "The type of the tools from the utilities library",
			},
			&cli.StringFlag{
				Name:    "pattern",
				Aliases: []string{"p"},
				Usage:   "The pattern of the tools from the utilities library",
			},
		},
		Action: func(c *cli.Context) error {
			switch c.String("type") {
			case "rn":
				fallthrough
			case "rename":
				return file.RenameFolderObjects(c.String("input"), c.String("output"))
			case "pic":
				fallthrough
			case "pic-spider":
				return file.Spider(c.String("input"), c.String("output"))
			case "srv":
				fallthrough
			case "split-reverse":
				return text.Srv(c.String("input"), c.String("output"), c.String("pattern"))
			case "rr":
				fallthrough
			case "remove-return":
				return text.RR(c.String("input"))
			default:
				return errors.New("unknown tool type selected")
			}
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		println(err.Error())
	}
}
