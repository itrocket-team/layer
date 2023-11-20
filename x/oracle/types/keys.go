package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "oracle"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// TipStoreKey defines the tip store key
	TipStoreKey = "tipStore"

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_oracle"

	ReportsKey = "Reports-value-"

	// TipStoreKey defines the tip store key
	//TipStoreKey = "tip_store"

	// CommitReportStoreKey defines the commit store key
	CommitReportStoreKey = "commit_report_store"

	ReporterStoreKey = "reporter_store"

	AggregateStoreKey = "aggergate_store"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func BlockKey(blockHeight int64) []byte {
	return sdk.Uint64ToBigEndian(uint64(blockHeight))
}
