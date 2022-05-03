package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// Sdk messages
	cdc.RegisterConcrete(&MsgCreateDid{}, "mantrachaind/CreateDid", nil)
	cdc.RegisterConcrete(&MsgUpdateDid{}, "mantrachaind/UpdateDid", nil)

	// State value data
	cdc.RegisterInterface((*StateValueData)(nil), nil)
	cdc.RegisterConcrete(&Did{}, "mantrachaind/Did", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// Sdk messages
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDid{},
		&MsgUpdateDid{},
	)

	// State value data
	registry.RegisterInterface("StateValueData", (*StateValueData)(nil))
	registry.RegisterImplementations((*StateValueData)(nil), &Did{})

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
