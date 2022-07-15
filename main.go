package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dimiro1/banner"
	"github.com/mattn/go-colorable"
	"github.com/urfave/cli/v2"
)

func main() {
	var title string
	var path string

	copyright := fmt.Sprintf("Wrecking Ball © %d", time.Now().Year())

	app := &cli.App{
		Name:        "banner",
		Usage:       "A cross-platform server-intended banner generator",
		Description: "Banner intends to be simple, lightweight, and easy to use",
		Copyright:   copyright,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Destination: &title,
				Name:        "title",
				Usage:       "Banner title",
				Value:       "Wrecking Ball",
			},
			&cli.StringFlag{
				Destination: &path,
				Name:        "path",
				Required: true,
				Usage:       "Path to the commands text file",
			},
		},
		Action: func(cCtx *cli.Context) error {
			//////
			// Parse flags.
			//////

			var titleFlag string
			if cCtx.NArg() > 0 {
				titleFlag = cCtx.Args().Get(0)
			}
			if titleFlag != "" {
				title = titleFlag
			}

			//////
			// Open commands text file.
			//////

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			commandsText, err := ioutil.ReadAll(file)
			if err != nil {
				log.Fatal(err)
			}

			//////
			// Add content to the template.
			//////

			templ := `{{ .Title "` + title + `" "" 0 }}
{{ .AnsiColor.BrightCyan }}The following aliases, utils, and commands are available:{{ .AnsiColor.Default }}
`
			templ += string(commandsText)

			//////
			// Print banner.
			//////

			banner.InitString(colorable.NewColorableStdout(), true, true, templ)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
