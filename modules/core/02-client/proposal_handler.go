package client

import (
	sdk "github.com/Finschia/finschia-sdk/types"
	sdkerrors "github.com/Finschia/finschia-sdk/types/errors"
	govtypes "github.com/Finschia/finschia-sdk/x/gov/types"

	"github.com/Finschia/ibc-go/v3/modules/core/02-client/keeper"
	"github.com/Finschia/ibc-go/v3/modules/core/02-client/types"
)

// NewClientProposalHandler defines the 02-client proposal handler
func NewClientProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.ClientUpdateProposal:
			return k.ClientUpdateProposal(ctx, c)
		case *types.UpgradeProposal:
			return k.HandleUpgradeProposal(ctx, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized ibc proposal content type: %T", c)
		}
	}
}
