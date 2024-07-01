package main

import (
	"flag"
	"mygoproject/init/cmd"
)

var configPathFlag = flag.String("config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd("./config.toml")
}
