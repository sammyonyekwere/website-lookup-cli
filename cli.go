package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Saamz Website Lookup CLI"
	app.Usage = "easily query IPs, MX records, Name Servers and CNames"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "saamz.com",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Get the Name Servers for a particular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))

				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
