package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/regen-network/bec/x/blog"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        blog.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", blog.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePost())

	return cmd
}

func CmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [author] [title] [body]",
		Short: "Creates a new post",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			argsTitle := string(args[1])
			argsBody := string(args[2])

			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err = client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := blog.MsgCreatePostRequest{
				Author: clientCtx.GetFromAddress().String(),
				Title:  argsTitle,
				Body:   argsBody,
			}
			svcMsgClientConn := &ServiceMsgClientConn{}
			msgClient := blog.NewMsgClient(svcMsgClientConn)
			_, err = msgClient.CreatePost(cmd.Context(), &msg)
			if err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), svcMsgClientConn.Msgs...)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
