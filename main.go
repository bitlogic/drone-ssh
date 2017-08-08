package main

import (
	"os"

	"github.com/appleboy/easyssh-proxy"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"
)

// Version set at compile-time
var Version = "v1.1.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "Gitlab SSH"
	app.Usage = "Executing remote ssh commands"
	app.Copyright = "Copyright (c) 2017 bitlogic.io"
	app.Authors = []cli.Author{
		{
			Name:  "Bo-Yi Wu",
			Email: "appleboy.tw@gmail.com",
		},
		{
			Name:  "Federico Aguirre",
			Email: "federico.aguirre@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "ssh-key",
			Usage:  "private ssh key",
			EnvVar: "SSH_KEY",
		},
		cli.StringFlag{
			Name:   "key-path,i",
			Usage:  "ssh private key path",
			EnvVar: "SSH_KEY_PATH",
		},
		cli.StringFlag{
			Name:   "username,user,u",
			Usage:  "connect as user",
			EnvVar: "SSH_USERNAME",
			Value:  "root",
		},
		cli.StringFlag{
			Name:   "password,P",
			Usage:  "user password",
			EnvVar: "SSH_PASSWORD",
		},
		cli.StringSliceFlag{
			Name:   "host,H",
			Usage:  "connect to host",
			EnvVar: "SSH_HOST",
		},
		cli.IntFlag{
			Name:   "port,p",
			Usage:  "connect to port",
			EnvVar: "SSH_PORT",
			Value:  22,
		},
		cli.DurationFlag{
			Name:   "timeout,t",
			Usage:  "connection timeout",
			EnvVar: "SSH_TIMEOUT",
		},
		cli.IntFlag{
			Name:   "command.timeout,T",
			Usage:  "command timeout",
			EnvVar: "SSH_COMMAND_TIMEOUT",
			Value:  60,
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
		cli.StringFlag{
			Name:   "proxy.ssh-key",
			Usage:  "private ssh key of proxy",
			EnvVar: "PROXY_SSH_KEY",
		},
		cli.StringFlag{
			Name:   "proxy.key-path",
			Usage:  "ssh private key path of proxy",
			EnvVar: "PROXY_SSH_KEY_PATH",
		},
		cli.StringFlag{
			Name:   "proxy.username",
			Usage:  "connect as user of proxy",
			EnvVar: "PROXY_SSH_USERNAME",
			Value:  "root",
		},
		cli.StringFlag{
			Name:   "proxy.password",
			Usage:  "user password of proxy",
			EnvVar: "PROXY_SSH_PASSWORD",
		},
		cli.StringFlag{
			Name:   "proxy.host",
			Usage:  "connect to host of proxy",
			EnvVar: "PROXY_SSH_HOST",
		},
		cli.StringFlag{
			Name:   "proxy.port",
			Usage:  "connect to port of proxy",
			EnvVar: "PROXY_SSH_PORT",
			Value:  "22",
		},
		cli.DurationFlag{
			Name:   "proxy.timeout",
			Usage:  "proxy connection timeout",
			EnvVar: "PROXY_SSH_TIMEOUT",
		},
		cli.StringSliceFlag{
			Name:   "secrets",
			Usage:  "plugin secret",
			EnvVar: "PLUGIN_SECRETS",
		},
		cli.StringSliceFlag{
			Name:   "envs",
			Usage:  "Pass envs",
			EnvVar: "PLUGIN_ENVS",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "PLUGIN_DEBUG",
		},
	}

	// Override a template
	cli.AppHelpTemplate = `
  ________.__  __  .__        ___.               _________ _________ ___ ___  
 /  _____/|__|/  |_|  | _____ \_ |__            /   _____//   _____//   |   \ 
/   \  ___|  \   __\  | \__  \ | __ \   ______  \_____  \ \_____  \/    ~    \
\    \_\  \  ||  | |  |__/ __ \| \_\ \ /_____/  /        \/        \    Y    /
 \______  /__||__| |____(____  /___  /         /_______  /_______  /\___|_  / 
        \/                   \/    \/                  \/        \/       \/
                                                    version: {{.Version}}
NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
REPOSITORY:
    Github: https://github.com/appleboy/bitlogic/gitlab-ssh
`

	app.Run(os.Args)
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}
	if len(c.Args()) == 0 {
		return nil
	}
	plugin := Plugin{
		Config: Config{
			Key:            c.String("ssh-key"),
			KeyPath:        c.String("key-path"),
			UserName:       c.String("user"),
			Password:       c.String("password"),
			Host:           c.StringSlice("host"),
			Port:           c.Int("port"),
			Timeout:        c.Duration("timeout"),
			CommandTimeout: c.Int("command.timeout"),
			Script:         c.Args(),
			Secrets:        c.StringSlice("secrets"),
			Envs:           c.StringSlice("envs"),
			Debug:          c.Bool("debug"),
			Proxy: easyssh.DefaultConfig{
				Key:      c.String("proxy.ssh-key"),
				KeyPath:  c.String("proxy.key-path"),
				User:     c.String("proxy.username"),
				Password: c.String("proxy.password"),
				Server:   c.String("proxy.host"),
				Port:     c.String("proxy.port"),
				Timeout:  c.Duration("proxy.timeout"),
			},
		},
	}

	return plugin.Exec()
}
