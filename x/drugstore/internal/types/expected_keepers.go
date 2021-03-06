package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SupplyKeeper is required for mining coin
type AdminKeeper interface {
	// SaleDrugs is used for drugstore to sale drugs on blockchain.
	SaleDrugs(ctx sdk.Context, patient string, id string, drugstore string) error
}
