package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	// this line is used by starport scaffolding # 1
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDelegate{}, "dualstaking/Delegate", nil)
	cdc.RegisterConcrete(&MsgRedelegate{}, "dualstaking/Redelegate", nil)
	cdc.RegisterConcrete(&MsgUnbond{}, "dualstaking/Unbond", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRedelegate{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUnbond{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
