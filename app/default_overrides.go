package app

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// bankModule defines a custom wrapper around the x/bank module's AppModuleBasic
// implementation to provide custom default genesis state.
type bankModule struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns custom x/bank module genesis state.
func (bankModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	metadata := banktypes.Metadata{
		Description: "The native token of the Tellor Layer.",
		Base:        BondDenom,
		Name:        DisplayDenom,
		Display:     DisplayDenom,
		Symbol:      DisplayDenom,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    BondDenom, // ie Wei
				Exponent: 0,
			},
			{
				Denom:    DisplayDenom, //ie Ether
				Exponent: 6,
			},
		},
	}

	genState := banktypes.DefaultGenesisState()
	genState.DenomMetadata = append(genState.DenomMetadata, metadata)

	return cdc.MustMarshalJSON(genState)
}

// stakingModule wraps the x/staking module in order to overwrite specific
// ModuleManager APIs.
type stakingModule struct {
	staking.AppModuleBasic
}

// DefaultGenesis returns custom x/staking module genesis state.
func (stakingModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	params := stakingtypes.DefaultParams()
	params.BondDenom = BondDenom

	return cdc.MustMarshalJSON(&stakingtypes.GenesisState{
		Params: params,
	})
}

type mintModule struct {
	mint.AppModuleBasic
}

// DefaultGenesis returns custom x/mint module genesis state.
func (mintModule) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	genState := minttypes.DefaultGenesisState()
	genState.Params.MintDenom = BondDenom

	return cdc.MustMarshalJSON(genState)
}
