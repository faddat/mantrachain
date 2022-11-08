package cli

import (
	"context"
	"strconv"
	"strings"

	"github.com/LimeChain/mantrachain/x/vault/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdGetNftStake() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-stake [marketplace_creator] [marketplace_id] [collection_creator] [collection_id] [id]",
		Short: "Query a nft stake",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			reqMarketplaceCreator := args[0]
			reqMarketplaceId := args[1]
			reqCollectionCreator := args[2]
			reqCollectionId := args[3]
			reqNftId := args[4]

			if strings.TrimSpace(reqMarketplaceId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidMarketplaceId, "empty marketplace id")
			}

			if strings.TrimSpace(reqCollectionId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidCollectionId, "empty collection id")
			}

			if strings.TrimSpace(reqNftId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidNftId, "empty nft id")
			}

			marketplaceCreator, err := sdk.AccAddressFromBech32(reqMarketplaceCreator)
			if err != nil {
				return err
			}

			collectionCreator, err := sdk.AccAddressFromBech32(reqCollectionCreator)
			if err != nil {
				return err
			}

			params := &types.QueryGetNftStakeRequest{
				MarketplaceCreator: marketplaceCreator.String(),
				MarketplaceId:      reqMarketplaceId,
				CollectionCreator:  collectionCreator.String(),
				CollectionId:       reqCollectionId,
				NftId:              reqNftId,
			}

			res, err := queryClient.NftStake(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdGetNftBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nft-balances [marketplace_creator] [marketplace_id] [collection_creator] [collection_id] [id] [staking_chain] [staking_validator]",
		Short: "Query a nft balance",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			reqMarketplaceCreator := args[0]
			reqMarketplaceId := args[1]
			reqCollectionCreator := args[2]
			reqCollectionId := args[3]
			reqNftId := args[4]
			reqStakingChain := args[5]
			reqStakingValidator := args[6]

			if strings.TrimSpace(reqMarketplaceId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidMarketplaceId, "empty marketplace id")
			}

			if strings.TrimSpace(reqCollectionId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidCollectionId, "empty collection id")
			}

			if strings.TrimSpace(reqNftId) == "" {
				return sdkerrors.Wrap(types.ErrInvalidNftId, "empty nft id")
			}

			marketplaceCreator, err := sdk.AccAddressFromBech32(reqMarketplaceCreator)
			if err != nil {
				return err
			}

			collectionCreator, err := sdk.AccAddressFromBech32(reqCollectionCreator)
			if err != nil {
				return err
			}

			params := &types.QueryGetNftBalancesRequest{
				MarketplaceCreator: marketplaceCreator.String(),
				MarketplaceId:      reqMarketplaceId,
				CollectionCreator:  collectionCreator.String(),
				CollectionId:       reqCollectionId,
				NftId:              reqNftId,
				StakingChain:       reqStakingChain,
				StakingValidator:   reqStakingValidator,
			}

			res, err := queryClient.NftBalances(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
