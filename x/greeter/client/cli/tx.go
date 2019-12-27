package cli

import (
	"github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	gtypes "github.com/nabakirov/hellochain/x/greeter/internal/types"
)

func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	greetingTxCmd := &cobra.Command{
		Use: "greeter",
		Short: "greeter transaction subcommands",
		DisableFlagParsing: true,
		SuggestionsMinimumDistance: 2,
		RunE: client.ValidateCmd,
	}
	greetingTxCmd.AddCommand(client.PostCommands(
		GetCmdSayHello(cdc),
	)...)

	return greetingTxCmd
}

func GetCmdSayHello(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "say [body] [addr]",
		Short: "send a greeting to another user. Usage: say [body] [address]",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			sender := cliCtx.GetFromAddress()
			body := args[0]

			recipient, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := gtypes.NewMsgGreet(sender, body, recipient)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},

	}
}