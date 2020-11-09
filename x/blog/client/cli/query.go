package cli

import (
	"fmt"

	// "strings"

	"github.com/amaurymartiny/bec/x/blog"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group glob queries under a subcommand
	cmd := &cobra.Command{
		Use:                        blog.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", blog.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(cmdListPost())

	return cmd
}

func cmdListPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-post",
		Short: "list all post",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := blog.NewQueryClient(clientCtx)

			params := &blog.QueryAllPostsRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.AllPosts(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
