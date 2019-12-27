package main

import (
	app "github.com/nabakirov/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/tendermint/tendermint/libs/cli"
)

func main() {
	params := starter.NewServerCommandParams(
		"hcd",
		"hellochain AppDaemon",
		starter.NewAppCreator(app.NewHelloChainApp),
		starter.NewAppExporter(app.NewHelloChainApp),
	)
	serverCmd := starter.NewServerCommand(params)

	executor := cli.PrepareBaseCmd(serverCmd, "HC", starter.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
