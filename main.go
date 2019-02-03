package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "GoSome"
	app.Usage = "Utility to manage your routine actions for you"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello it's still work in progress")
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
