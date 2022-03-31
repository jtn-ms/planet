package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreatePost{}, "blog/CreatePost", nil)
	cdc.RegisterConcrete(&MsgUpdatePost{}, "blog/UpdatePost", nil)
	cdc.RegisterConcrete(&MsgDeletePost{}, "blog/DeletePost", nil)
	cdc.RegisterConcrete(&MsgCreateSentPost{}, "blog/CreateSentPost", nil)
	cdc.RegisterConcrete(&MsgUpdateSentPost{}, "blog/UpdateSentPost", nil)
	cdc.RegisterConcrete(&MsgDeleteSentPost{}, "blog/DeleteSentPost", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreatePost{},
		&MsgUpdatePost{},
		&MsgDeletePost{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateSentPost{},
		&MsgUpdateSentPost{},
		&MsgDeleteSentPost{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
