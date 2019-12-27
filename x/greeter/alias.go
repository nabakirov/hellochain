package greeter

import (
	"github.com/nabakirov/hellochain/x/greeter/internal/types"
	"github.com/nabakirov/hellochain/x/greeter/internal/keeper"
)

const (
	ModuleName 	= types.ModuleName
	RouterKey 	= types.RouterKey
	StoreKey 	= types.StoreKey
)

var (
	NewKeeper 	= keeper.NewKeeper
	NewQuerier 	= keeper.NewQuerier
	NewMsgGreet = types.NewMsgGreet
	NewGreeting = types.NewGreeting
)

type (
	Keeper	= keeper.Keeper
	MsgGreet	= types.MsgGreet
	Greeting 	= types.Greeting
	QueryResGreetings = types.QueryResGreetings
	GreetingsList = types.GreetingList
)