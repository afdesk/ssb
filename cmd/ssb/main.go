/*
Copyright © 2022 afdesk <amf@afdesk.com>
*/
package main

import (
	"log"

	"github.com/afdesk/ssb/pkg/commands"
)

var (
	version = "dev"
)

func main() {
	app := commands.NewApp(version)
	if err := app.Execute(); err != nil {
		log.Fatal(err)
	}
}
