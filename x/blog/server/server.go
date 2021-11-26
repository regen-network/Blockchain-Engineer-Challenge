package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/regen-network/bec/x/blog"
)

type serverImpl struct {
	cdc      codec.Codec
	storeKey sdk.StoreKey
}

func newServer(cdc codec.Codec, storeKey sdk.StoreKey) serverImpl {
	s := serverImpl{cdc: cdc, storeKey: storeKey}

	return s
}

func RegisterServices(cdc codec.Codec, storeKey sdk.StoreKey, configurator module.Configurator) {
	impl := newServer(cdc, storeKey)
	blog.RegisterMsgServer(configurator.MsgServer(), impl)
	blog.RegisterQueryServer(configurator.QueryServer(), impl)
}
