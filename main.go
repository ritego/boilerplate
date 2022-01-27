package main

import (
	"github.com/aellacredit/jara/cli"
	"github.com/aellacredit/jara/config"
	"github.com/aellacredit/jara/store"
)

func main() {
	config.Init()
	store.Init()
	cli.Init()
}
