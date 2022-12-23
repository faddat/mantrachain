package vault

import (
	"math/rand"

	"github.com/LimeChain/mantrachain/testutil/sample"
	vaultsimulation "github.com/LimeChain/mantrachain/x/vault/simulation"
	"github.com/LimeChain/mantrachain/x/vault/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = vaultsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateChainValidatorBridge = "op_weight_msg_chain_validator_bridge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateChainValidatorBridge int = 100

	opWeightMsgUpdateChainValidatorBridge = "op_weight_msg_chain_validator_bridge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateChainValidatorBridge int = 100

	opWeightMsgDeleteChainValidatorBridge = "op_weight_msg_chain_validator_bridge"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteChainValidatorBridge int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vaultGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		ChainValidatorBridgeList: []types.ChainValidatorBridge{
			{
				Creator:  sample.AccAddress(),
				BridgeId: "0",
			},
			{
				Creator:  sample.AccAddress(),
				BridgeId: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vaultGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateChainValidatorBridge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateChainValidatorBridge, &weightMsgCreateChainValidatorBridge, nil,
		func(_ *rand.Rand) {
			weightMsgCreateChainValidatorBridge = defaultWeightMsgCreateChainValidatorBridge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateChainValidatorBridge,
		vaultsimulation.SimulateMsgCreateChainValidatorBridge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateChainValidatorBridge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateChainValidatorBridge, &weightMsgUpdateChainValidatorBridge, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateChainValidatorBridge = defaultWeightMsgUpdateChainValidatorBridge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateChainValidatorBridge,
		vaultsimulation.SimulateMsgUpdateChainValidatorBridge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteChainValidatorBridge int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteChainValidatorBridge, &weightMsgDeleteChainValidatorBridge, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteChainValidatorBridge = defaultWeightMsgDeleteChainValidatorBridge
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteChainValidatorBridge,
		vaultsimulation.SimulateMsgDeleteChainValidatorBridge(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
