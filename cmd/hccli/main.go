package main

import (
	"github.com/tendermint/tendermint/libs/cli"

	app "github.com/nabakirov/hellochain"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
	"github.com/nabakirov/hellochain/x/greeter"
)

func main() {

	starter.BuildModuleBasics(greeter.AppModuleBasic{})

	rootCmd := starter.NewCLICommand()

	txCmd := starter.TxCmd(starter.Cdc)
	queryCmd := starter.QueryCmd(starter.Cdc)

	// add more Tx and Query commands
	app.ModuleBasics.AddTxCommands(txCmd, starter.Cdc)
	app.ModuleBasics.AddQueryCommands(queryCmd, starter.Cdc)
	rootCmd.AddCommand(txCmd, queryCmd)

	executor := cli.PrepareMainCmd(rootCmd, "HC", starter.DefaultCLIHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
