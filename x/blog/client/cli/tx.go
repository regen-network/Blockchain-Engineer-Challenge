package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/amaurymartiny/bec/x/blog"
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

	cmd.AddCommand(cmdCreatePost())

	return cmd
}

func cmdCreatePost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-post [title] [body]",
		Short: "Creates a new post",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := cmd.Flags().Set(flags.FlagFrom, args[0])
			if err != nil {
				return err
			}

			argsTitle := string(args[0])
			argsBody := string(args[1])

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

			return nil
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
