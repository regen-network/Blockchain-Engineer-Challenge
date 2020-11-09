package data

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_, _, _ sdk.MsgRequest = &MsgCreatepostRequest{}, &MsgCreateCommentRequest{}
)

func (m *MsgCreatepostRequest) ValidateBasic() error {
	return nil
}

func (m *MsgCreatepostRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}

func (m *MsgCreateCommentRequest) ValidateBasic() error {
	return nil
}

func (m *MsgCreateCommentRequest) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{addr}
}
