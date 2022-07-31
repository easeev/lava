package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
)

type PairingKeeper interface {
	UnstakeEntry(ctx sdk.Context, provider bool, chainID string, creator string) error
	CreditStakeEntry(ctx sdk.Context, chainID string, lookUpAddress sdk.AccAddress, creditAmount sdk.Coin, isProvider bool) (bool, error)
	VerifyPairingData(ctx sdk.Context, chainID string, clientAddress sdk.AccAddress, block uint64) (clientStakeEntryRet *epochstoragetypes.StakeEntry, errorRet error)
	JailEntry(ctx sdk.Context, account sdk.AccAddress, isProvider bool, chainID string, jailStartBlock uint64, jailBlocks uint64, bail sdk.Coin) error
	BailEntry(ctx sdk.Context, account sdk.AccAddress, isProvider bool, chainID string, bail sdk.Coin) error
	SlashEntry(ctx sdk.Context, account sdk.AccAddress, isProvider bool, chainID string, percentage sdk.Dec) (sdk.Coin, error)
}

type EpochstorageKeeper interface {
	GetEpochStart(ctx sdk.Context) uint64
	BlocksToSave(ctx sdk.Context, block uint64) (res uint64, err error)
	EpochBlocks(ctx sdk.Context, block uint64) (res uint64, err error)
	GetEpochStartForBlock(ctx sdk.Context, block uint64) (epochStart uint64, blockInEpoch uint64, err error)
	GetNextEpoch(ctx sdk.Context, epoch uint64) (nextEpoch uint64, err error)
	GetStakeEntryForClientEpoch(ctx sdk.Context, chainID string, selectedClient sdk.AccAddress, epoch uint64) (entry *epochstoragetypes.StakeEntry, err error)
	GetStakeEntryForProviderEpoch(ctx sdk.Context, chainID string, selectedProvider sdk.AccAddress, epoch uint64) (entry *epochstoragetypes.StakeEntry, err error)
	GetStakeEntryForAllProvidersEpoch(ctx sdk.Context, chainID string, epoch uint64) (entrys *[]epochstoragetypes.StakeEntry, err error)
}

type SpecKeeper interface {
	IsSpecFoundAndActive(ctx sdk.Context, chainID string) (foundAndActive bool, found bool)
	IsFinalizedBlock(ctx sdk.Context, chainID string, requestedBlock int64, latestBlock int64) bool
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
