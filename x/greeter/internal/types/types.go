package types

import (
	"fmt"
	"strings"

	codec "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "greeter"
	StoreKey  = ModuleName
)

var (
	ModuleCdc = codec.New()
)

type Greeting struct {
	Sender    sdk.AccAddress `json:"sender" yaml:"sender"`
	Recipient sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Body      string         `json:"body" yaml:"body"`
}

type GreetingList []Greeting

func NewGreeting(sender sdk.AccAddress, body string, receiver sdk.AccAddress) Greeting {
	return Greeting{
		Recipient: receiver,
		Sender:    sender,
		Body:      body,
	}
}

func (g Greeting) String() string {
	return strings.TrimSpace(
		fmt.Sprintf(`Sender: %s Recipient: %s Body: %s`, g.Sender.String(), g.Recipient.String(), g.Body),
	)
}

type QueryResGreetings map[string][]Greeting

func (q QueryResGreetings) String() string {
	b := ModuleCdc.MustMarshalJSON(q)
	return string(b)
}

func NewQueryResGreetings() QueryResGreetings {
	return make(map[string][]Greeting)
}
