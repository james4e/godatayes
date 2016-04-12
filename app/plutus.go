package main

import (
	"fmt"
	"godatayes"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	// init datayes
	envInit()

	app := cli.NewApp()
	app.Name = "Plutus"
	app.Usage = "get Shanghai & Shenzhen stock market data"
	app.Version = "0.1.0"
	app.Author = "majian"
	app.Email = "james4e@163.com"
	commonFlag := []cli.Flag{
		// cli.StringFlag{Name: "ticker, t", Value: "", Usage: "the ticker of stock"},
		cli.StringFlag{Name: "output, o", Value: "", Usage: "the file name to store the data"},
	}
	app.Commands = []cli.Command{
		{
			Name:        "pull",
			Usage:       "pull <ticker> <start date> <end date> <options>",
			Description: "Pull the stock data of specific number and save them into a specific file. ticker and start date are required.",
			Flags:       commonFlag,
			Action: func(c *cli.Context) {
				if len(c.Args()) >= 2 {
					ticker := c.Args().First()
					start := c.Args().Get(1)
					end := c.Args().Get(2)
					fileName := c.String("output")
					if end == "" {
						end = time.Now().Format("20060102")
					}
					if fileName == "" {
						fileName = strings.Join([]string{ticker, "_", start, "_", end, ".txt"}, "")
					}
					data, err := godatayes.GetMktEqudByTicker(ticker, start, end)
					if err != nil {
						fmt.Println(err.Error())
					}
					ioutil.WriteFile(fileName, []byte(data), 0644)
					fmt.Println("done!")
				} else {
					cli.ShowCommandHelp(c, "pull")
				}

			},
		},
	}
	app.Run(os.Args)
}

func envInit() {
	godatayes.Init("b4be14a835d1550257a0423a58f1ec27b60ddd17e3ae7dade54e8976c6675697")
}
