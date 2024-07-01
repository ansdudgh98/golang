package cmd

import (
	"fmt"
	"mygoproject/config"
	"mygoproject/network"
)

type Cmd struct {
	config  *config.Config
	network *network.Network
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config:  config.NewConfig(filePath),
		network: network.NewNetwork(),
	}

	fmt.Println(c.config.Server.Port)

	c.network.ServerStar(c.config.Server.Port)

	return c
}
