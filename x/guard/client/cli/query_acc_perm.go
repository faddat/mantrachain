package cli

import (
    "context"
	
    "github.com/spf13/cobra"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/LimeChain/mantrachain/x/guard/types"
)

func CmdListAccPerm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-acc-perm",
		Short: "list all acc_perm",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllAccPermRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.AccPermAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowAccPerm() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-acc-perm [cat]",
		Short: "shows a acc_perm",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

             argCat := args[0]
            
            params := &types.QueryGetAccPermRequest{
                Cat: argCat,
                
            }

            res, err := queryClient.AccPerm(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
