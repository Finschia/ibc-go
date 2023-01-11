package utils

import (
	"context"

	"github.com/line/lbm-sdk/client"
	"github.com/line/lbm-sdk/codec"
	sdkerrors "github.com/line/lbm-sdk/types/errors"
	octypes "github.com/line/ostracon/types"

	"github.com/line/ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/line/ibc-go/modules/core/23-commitment/types"
	host "github.com/line/ibc-go/modules/core/24-host"
	ibcclient "github.com/line/ibc-go/modules/core/client"
	"github.com/line/ibc-go/modules/core/exported"
	ibcoctypes "github.com/line/ibc-go/modules/light-clients/99-ostracon/types"
)

// QueryClientState returns a client state. If prove is true, it performs an ABCI store query
// in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryClientState(
	clientCtx client.Context, clientID string, prove bool,
) (*types.QueryClientStateResponse, error) {
	if prove {
		return QueryClientStateABCI(clientCtx, clientID)
	}

	queryClient := types.NewQueryClient(clientCtx)
	req := &types.QueryClientStateRequest{
		ClientId: clientID,
	}

	return queryClient.ClientState(context.Background(), req)
}

// QueryClientStateABCI queries the store to get the light client state and a merkle proof.
func QueryClientStateABCI(
	clientCtx client.Context, clientID string,
) (*types.QueryClientStateResponse, error) {
	key := host.FullClientStateKey(clientID)

	value, proofBz, proofHeight, err := ibcclient.QueryOstraconProof(clientCtx, key)
	if err != nil {
		return nil, err
	}

	// check if client exists
	if len(value) == 0 {
		return nil, sdkerrors.Wrap(types.ErrClientNotFound, clientID)
	}

	cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

	clientState, err := types.UnmarshalClientState(cdc, value)
	if err != nil {
		return nil, err
	}

	anyClientState, err := types.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	clientStateRes := types.NewQueryClientStateResponse(anyClientState, proofBz, proofHeight)
	return clientStateRes, nil
}

// QueryConsensusState returns a consensus state. If prove is true, it performs an ABCI store
// query in order to retrieve the merkle proof. Otherwise, it uses the gRPC query client.
func QueryConsensusState(
	clientCtx client.Context, clientID string, height exported.Height, prove, latestHeight bool,
) (*types.QueryConsensusStateResponse, error) {
	if prove {
		return QueryConsensusStateABCI(clientCtx, clientID, height)
	}

	queryClient := types.NewQueryClient(clientCtx)
	req := &types.QueryConsensusStateRequest{
		ClientId:       clientID,
		RevisionNumber: height.GetRevisionNumber(),
		RevisionHeight: height.GetRevisionHeight(),
		LatestHeight:   latestHeight,
	}

	return queryClient.ConsensusState(context.Background(), req)
}

// QueryConsensusStateABCI queries the store to get the consensus state of a light client and a
// merkle proof of its existence or non-existence.
func QueryConsensusStateABCI(
	clientCtx client.Context, clientID string, height exported.Height,
) (*types.QueryConsensusStateResponse, error) {
	key := host.FullConsensusStateKey(clientID, height)

	value, proofBz, proofHeight, err := ibcclient.QueryOstraconProof(clientCtx, key)
	if err != nil {
		return nil, err
	}

	// check if consensus state exists
	if len(value) == 0 {
		return nil, sdkerrors.Wrap(types.ErrConsensusStateNotFound, clientID)
	}

	cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

	cs, err := types.UnmarshalConsensusState(cdc, value)
	if err != nil {
		return nil, err
	}

	anyConsensusState, err := types.PackConsensusState(cs)
	if err != nil {
		return nil, err
	}

	return types.NewQueryConsensusStateResponse(anyConsensusState, proofBz, proofHeight), nil
}

// QueryOstraconHeader takes a client context and returns the appropriate
// ostracon header
func QueryOstraconHeader(clientCtx client.Context) (ibcoctypes.Header, int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	info, err := node.ABCIInfo(context.Background())
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	var height int64
	if clientCtx.Height != 0 {
		height = clientCtx.Height
	} else {
		height = info.Response.LastBlockHeight
	}

	commit, err := node.Commit(context.Background(), &height)
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	page := 1
	count := 10_000

	validators, err := node.Validators(context.Background(), &height, &page, &count)
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	page = 0
	count = 10_000
	voters, err := node.Voters(context.Background(), &height, &page, &count)
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	protoCommit := commit.SignedHeader.ToProto()
	protoValset, err := octypes.NewValidatorSet(validators.Validators).ToProto()
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	protoVoterSet, err := octypes.WrapValidatorsToVoterSet(voters.Voters).ToProto()
	if err != nil {
		return ibcoctypes.Header{}, 0, err
	}

	header := ibcoctypes.Header{
		SignedHeader: protoCommit,
		ValidatorSet: protoValset,
		VoterSet:     protoVoterSet,
	}

	return header, height, nil
}

// QuerySelfConsensusState takes a client context and returns the appropriate
// ostracon consensus state
func QuerySelfConsensusState(clientCtx client.Context) (*ibcoctypes.ConsensusState, int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return &ibcoctypes.ConsensusState{}, 0, err
	}

	info, err := node.ABCIInfo(context.Background())
	if err != nil {
		return &ibcoctypes.ConsensusState{}, 0, err
	}

	var height int64
	if clientCtx.Height != 0 {
		height = clientCtx.Height
	} else {
		height = info.Response.LastBlockHeight
	}

	commit, err := node.Commit(context.Background(), &height)
	if err != nil {
		return &ibcoctypes.ConsensusState{}, 0, err
	}

	page := 1
	count := 10_000

	nextHeight := height + 1
	nextVals, err := node.Validators(context.Background(), &nextHeight, &page, &count)
	if err != nil {
		return &ibcoctypes.ConsensusState{}, 0, err
	}

	state := &ibcoctypes.ConsensusState{
		Timestamp:          commit.Time,
		Root:               commitmenttypes.NewMerkleRoot(commit.AppHash),
		NextValidatorsHash: octypes.NewValidatorSet(nextVals.Validators).Hash(),
	}

	return state, height, nil
}
