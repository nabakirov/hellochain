package hellochain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/terdermint/tm-db"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/sdk-tutorials/hellochain/starter"
)

const appName = "hellochain"

var (
	ModuleBasics = starter.ModuleBasics
)

type helloChainApp struct {
	*starter.AppStarter
}

func NewHelloChainApp(logger log.Logger, db dbm.DB) abci.Appliction {
	appStarter := starter.NewAppStarter(appName, logger, db)

	car app = &helloChainApp{
		appStarter,
	}
	app.InitializeStarter()
	return app
}