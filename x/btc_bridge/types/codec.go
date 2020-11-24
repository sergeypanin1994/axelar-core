package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTrackAddress{}, "btcbridge/TrackAddress", nil)
	cdc.RegisterConcrete(MsgTrackPubKey{}, "btcbridge/MsgTrackPubKey", nil)
	cdc.RegisterConcrete(MsgVerifyTx{}, "btcbridge/VerifyTx", nil)
	cdc.RegisterConcrete(MsgRawTx{}, "btcbridge/RawTx", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "btcbridge/Withdraw", nil)
	cdc.RegisterConcrete(&MsgVoteVerifiedTx{}, "btcbridge/MsgVoteVerifiedTx", nil)

	//	TODO: Replace bool with full bitcoin tx data when btcbridge uses write-in voting pattern
	cdc.RegisterConcrete(true, "btcbridge/VotingData", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}