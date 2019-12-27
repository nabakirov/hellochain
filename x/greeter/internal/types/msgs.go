package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const RouterKey = "greeter"

type MsgGreet struct {
	Body      string
	Sender    sdk.AccAddress
	Recipient sdk.AccAddress
}

func NewMsgGreet(sender sdk.AccAddress, body string, recipient sdk.AccAddress) MsgGreet {
	return MsgGreet{
		Body:      body,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg MsgGreet) Route() string { return RouterKey }

func (msg MsgGreet) Type() string { return "greet" }

func (msg MsgGreet) ValidateBasic() sdk.Error {
	if msg.Recipient.Empty() {
		return sdk.ErrInvalidAddress(msg.Recipient.String())
	}
	if len(msg.Sender) == 0 || len(msg.Body) == 0 || len(msg.Recipient) == 0 {
		return sdk.ErrUnknownRequest("Sender, Recipient and/or Body cannot be empty")
	}
	return nil
}

func (msg MsgGreet) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgGreet) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
