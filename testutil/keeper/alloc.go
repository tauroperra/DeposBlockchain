package keeper

// import (
// 	"testing"

// 	"github.com/cosmos/cosmos-sdk/codec"
// 	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	"github.com/cosmos/cosmos-sdk/store"
// 	storetypes "github.com/cosmos/cosmos-sdk/store/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/trdtech/deposnetwork/x/alloc/keeper"
// 	"github.com/trdtech/deposnetwork/x/alloc/types"
// 	"github.com/stretchr/testify/require"
// 	"github.com/tendermint/tendermint/libs/log"
// 	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
// 	tmdb "github.com/tendermint/tm-db"
// )

// func AllocKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
// 	storeKey := sdk.NewKVStoreKey(types.StoreKey)
// 	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

// 	db := tmdb.NewMemDB()
// 	stateStore := store.NewCommitMultiStore(db)
// 	stateStore.MountStoreWithDB(storeKey, sdk.StoreTypeIAVL, db)
// 	stateStore.MountStoreWithDB(memStoreKey, sdk.StoreTypeMemory, nil)
// 	require.NoError(t, stateStore.LoadLatestVersion())

// 	registry := codectypes.NewInterfaceRegistry()
// 	k := keeper.NewKeeper(
// 		codec.NewProtoCodec(registry),
// 		storeKey,
// 		memStoreKey,
// 		nil,
// 		nil,
// 		nil,
// 		nil,
// 	)

// 	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
// 	return k, ctx
// }
