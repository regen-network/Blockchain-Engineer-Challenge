package cli

import (
	"fmt"

	// "strings"

	"github.com/regen-network/bec/x/blog"
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

	cmd.AddCommand(
		CmdAllPosts(),
		CmdPostComments(),
	)

	return cmd
}

func CmdAllPosts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-post",
		Short: "list all post",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
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

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "blog")

	return cmd
}

func CmdPostComments() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-comments [slug]",
		Short: "list all comments of the post",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			argsSlug := string(args[0])

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := blog.NewQueryClient(clientCtx)

			params := &blog.QueryPostCommentsRequest{
				Pagination: pageReq,
				Slug:       argsSlug,
			}

			res, err := queryClient.PostComments(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "blog")

	return cmd
}
