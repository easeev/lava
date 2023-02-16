package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	common "github.com/lavanet/lava/common"
	"github.com/lavanet/lava/utils"
	"github.com/lavanet/lava/x/packages/types"
)

// SetPackage sets a specific package in its fixated form
func (k Keeper) SetPackage(ctx sdk.Context, packageIndex string, packageToSet types.Package) error {
	// get the entry to set
	entry, found := common.GetEntry(ctx, k.storeKey, types.PackageKeyPrefix, k.cdc, packageIndex)
	if !found {
		return utils.LavaError(ctx, k.Logger(ctx), "SetPackage_package_not_found", map[string]string{"packageIndex": packageIndex}, "could not find the package to set")
	}

	// marshal the packageToSet
	b := k.cdc.MustMarshal(&packageToSet)

	// update the entry's marshaledData field and set the entry
	entry.MarshaledData = b
	common.SetEntry(ctx, k.storeKey, types.PackageKeyPrefix, k.cdc, entry)
	return nil
}

// AddPackage adds a new package to the KVStore. The function returns if the added package is a first version package
func (k Keeper) AddPackage(ctx sdk.Context, packageToAdd types.Package) error {
	// overwrite the packageToAdd's block field with the current block height
	packageToAdd.Block = uint64(ctx.BlockHeight())

	// TODO: verify the CU per epoch field

	// make the package's subscriptions field zero (it's a new package, so no one is subscribed yet)
	packageToAdd.Subscriptions = uint64(0)

	// marshal the packageToAdd
	b := k.cdc.MustMarshal(&packageToAdd)

	// add a new fixated entry with the marshaled packageToAdd
	err := common.AddEntry(ctx, k.storeKey, types.PackageKeyPrefix, types.UniqueIndexKeyPrefix(), k.cdc, packageToAdd.Index, b)
	if err != nil {
		return utils.LavaError(ctx, k.Logger(ctx), "AddPackage_add_fixated_entry_failed", map[string]string{"packageToAdd": packageToAdd.String()}, "could not add new package fixated entry to storage")
	}

	return nil
}

// GetPackageForBlock returns a package or its older version (according to the requested block) from its index
func (k Keeper) GetPackageForBlock(ctx sdk.Context, packageIndex string, block uint64) (types.Package, bool, string) {
	// get the fixation entry that is suits the requested block
	entry, found := common.GetEntryOlderVersionByBlock(ctx, k.storeKey, types.PackageKeyPrefix, k.cdc, packageIndex, block)
	if !found {
		return types.Package{}, found, ""
	}

	// unmarshal the entry's marshaled data to get the package
	umarshaledPackage := types.Package{}
	k.cdc.MustUnmarshal(entry.MarshaledData, &umarshaledPackage)

	return umarshaledPackage, found, entry.GetIndex()
}

// RemovePackage removes a package from the KVStore (essentially, removes the corresponding fixation entry)
func (k Keeper) RemovePackage(
	ctx sdk.Context,
	packageIndex string,
) {
	common.RemoveEntry(ctx, k.storeKey, types.PackageKeyPrefix, packageIndex)
}

// GetAllPackageVersions returns all package versions by the given index
func (k Keeper) GetAllPackageVersions(ctx sdk.Context, packageIndex string) []types.Package {
	// get all the fixation entries by the packageIndex
	entries := common.GetAllEntriesForIndex(ctx, k.storeKey, types.PackageKeyPrefix, k.cdc, packageIndex)

	// go over the entries
	packages := []types.Package{}
	for _, entry := range entries {
		// unmarshal the entry's marshaled data to get the package
		var umarshaledPackage types.Package
		k.cdc.MustUnmarshal(entry.MarshaledData, &umarshaledPackage)

		// add the package to the packages array
		packages = append(packages, umarshaledPackage)
	}

	return packages
}

// Function to get a package's latest version
func (k Keeper) GetPackageLatestVersion(ctx sdk.Context, packageIndex string) (*types.Package, error) {
	// get the requested package's latest fixated entry
	latestPackage, found, _ := k.GetPackageForBlock(ctx, packageIndex, uint64(ctx.BlockHeight()))
	if !found {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_package_by_epoch", map[string]string{"packageIndex": packageIndex}, "could not get packageFixationEntry with index and epoch")
	}

	return &latestPackage, nil
}

func (k Keeper) AddSubscription(ctx sdk.Context, packageIndex string) error {
	latestPackage, err := k.GetPackageLatestVersion(ctx, packageIndex)
	if err != nil {
		return err
	}

	latestPackage.Subscriptions++

	k.SetPackage(ctx, packageIndex, *latestPackage)

	return nil
}

func (k Keeper) SubSubscription(ctx sdk.Context, packageIndex string) error {
	latestPackage, err := k.GetPackageLatestVersion(ctx, packageIndex)
	if err != nil {
		return err
	}

	latestPackage.Subscriptions--

	k.SetPackage(ctx, packageIndex, *latestPackage)

	return nil
}
