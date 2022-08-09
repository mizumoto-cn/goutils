package main

import (
	"errors"
	"os"

	"github.com/mizumoto-cn/goutils/filetools"

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
		},
		Action: func(c *cli.Context) error {
			switch c.String("type") {
			case "rn":
				fallthrough
			case "rename":
				return filetools.RenameFolderObjects(c.String("input"), c.String("output"))
			case "pic":
				fallthrough
			case "pic-spider":
				return filetools.Spider(c.String("input"), c.String("output"))
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
