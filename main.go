package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	AppName    = "archloong"
	AppVersion = "0.1"
	AppConfig  = "archloong.toml"
)

var cfg *Config

func init() {
	cfg, _ = LoadConfig(AppName, AppVersion, AppConfig)
}

func main() {
	al := NewArchLoong(cfg)
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "run web server",
				Action: func(cCtx *cli.Context) error {
					al.Server()
					return nil
				},
			},
			{
				Name:    "update",
				Aliases: []string{"u"},
				Usage:   "update database from the remote repo",
				Action: func(cCtx *cli.Context) error {
					al.Update()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
