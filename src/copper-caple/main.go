package main

import (
	"log"
	"os"

	"caple"

	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(os.Stdout)

	app := cli.NewApp()
	app.Name = "copper-caple"
	app.Usage = "Backend microservice for managing competences"
	app.Version = caple.Version
	app.Copyright = "(C)2016 B.J.W. Schaap"
	app.Author = "Bastiaan Schaap"
	app.Email = "bastiaan.schaap@gmail.com"
	app.UsageText = `./copper-caple --listen 127.0.0.1:12345 --db 127.0.0.1:3456
    --db (-d) can be omitted if DB_URL environment variable is set.
    --listen (-l) can be omitted, and has a default, or can be set with LISTEN_ADDRESS environment variable`

	// Define the configuration flags the program can/should use
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "address, a",
			Value:  "127.0.0.1:26257",
			Usage:  "Address (host:port) of the database host",
			EnvVar: "DB_ADDRESS",
		},
		cli.StringFlag{
			Name:   "db, d",
			Value:  "capledb",
			Usage:  "Database name",
			EnvVar: "DB_NAME",
		},
		cli.StringFlag{
			Name:   "user, u",
			Value:  "caple",
			Usage:  "Database user",
			EnvVar: "DB_USER",
		},
		cli.StringFlag{
			Name:   "password, p",
			Value:  "1109@Pie",
			Usage:  "Database password",
			EnvVar: "DB_PASSWORD",
		},
		cli.IntFlag{
			Name:   "poolsize, s",
			Value:  20,
			Usage:  "Database connection pool size",
			EnvVar: "DB_POOLSIZE",
		},
		cli.StringFlag{
			Name:   "listen, l",
			Value:  "127.0.0.1:1616",
			Usage:  "Listen address and port",
			EnvVar: "LISTEN_ADDRESS",
		},
		cli.StringFlag{
			Name:   "apikey, k",
			Value:  "^zqpp$!Zv#Ahp=f$#69yFe%h6f_fv5FCK+mJ+PmG",
			Usage:  "API Key needed for using this service",
			EnvVar: "API_KEY",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Set to true for extra debug logging",
			EnvVar: "DEBUG",
		},
	}

	// Set the main program logic
	app.Action = func(c *cli.Context) error {
		return caple.StartServer(c)
	}

	// Now start doing stuff
	app.Run(os.Args)
}
