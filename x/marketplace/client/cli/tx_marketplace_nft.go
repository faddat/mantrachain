package cli

import (
	"fmt"
	"strconv"

	"github.com/LimeChain/mantrachain/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdBuyNft() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-nft [nft-id]",
		Short: "Broadcast message buy-nft",
		Long: "Buy a NFT. " +
			"[nft-id] is the NFT id.",
		Example: fmt.Sprintf(
			"$ %s tx marketplace buy-nft <nft-id> "+
				"--from=<from> "+
				"--marketplace-creator=<marketplace-creator> "+
				"--marketplace-id=<marketplace-id> "+
				"--collection-creator=<collection-creator> "+
				"--collection-id=<collection-id> "+
				"--cw20-contract-address=<cw20-contract-address> "+
				"--chain-id=<chain-id> "+
				"--chain=<chain> "+
				"--validator=<validator> ",
			version.AppName,
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argNftId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			marketplaceCreator, err := cmd.Flags().GetString(FlagMarketplaceCreator)
			if err != nil {
				return err
			}

			marketplaceId, err := cmd.Flags().GetString(FlagMarketplaceId)
			if err != nil {
				return err
			}

			collectionCreator, err := cmd.Flags().GetString(FlagCollectionCreator)
			if err != nil {
				return err
			}

			collectionId, err := cmd.Flags().GetString(FlagCollectionId)
			if err != nil {
				return err
			}

			cw20ContractAddress, err := cmd.Flags().GetString(FlagCw20ContractAddress)
			if err != nil {
				return err
			}

			stakingChain, err := cmd.Flags().GetString(FlagStakingChain)
			if err != nil {
				return err
			}

			stakingValidator, err := cmd.Flags().GetString(FlagStakingValidator)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyNft(
				clientCtx.GetFromAddress().String(),
				marketplaceCreator,
				marketplaceId,
				collectionCreator,
				collectionId,
				argNftId,
				cw20ContractAddress,
				stakingChain,
				stakingValidator,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FsBuyNft)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
