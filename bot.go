package main

import (
	"github.com/syfaro/finch"
	_ "github.com/syfaro/finch-example/commands/minecraft"
	_ "github.com/syfaro/finch/commands/help"
	_ "github.com/syfaro/finch/commands/info"
	_ "github.com/syfaro/finch/commands/stats"
	"os"
)

func main() {
	f := finch.NewFinch(os.Getenv("TELEGRAM_APITOKEN"))

	f.Start()
}
