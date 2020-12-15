package main

import (
	"github.com/Goscord/Bot/command"
	"github.com/Goscord/Bot/config"
	"github.com/Goscord/Bot/event"
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
)

var (
	Client *gateway.Session
	Config *config.Config
	CmdMgr *command.Manager
)

func main() {
	Config, _ = config.GetConfig()
	Client = goscord.New(&gateway.Options{Token: Config.Token})
	CmdMgr = command.Init()

	_ = Client.On("ready", event.OnReady(Client, Config))
	_ = Client.On("message", CmdMgr.Handler(Client, Config))

	if err := Client.Login(); err != nil {
		panic(err)
	}

	select{}
}