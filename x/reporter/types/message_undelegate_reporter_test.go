package types

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tellor-io/layer/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestMsgChangeReporter_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgChangeReporter
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgChangeReporter{
				DelegatorAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgChangeReporter{
				DelegatorAddress: sample.AccAddress(),
				ReporterAddress:  sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
