package types_test

import (
	"testing"

	"github.com/stretchr/testify/suite"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Finschia/finschia-sdk/codec"
	sdk "github.com/Finschia/finschia-sdk/types"

	clienttypes "github.com/Finschia/ibc-go/v3/modules/core/02-client/types"
	"github.com/Finschia/ibc-go/v3/modules/core/exported"
	"github.com/Finschia/ibc-go/v3/testing/simapp"
)

const (
	height = 4
)

var clientHeight = clienttypes.NewHeight(0, 10)

type LocalhostTestSuite struct {
	suite.Suite

	cdc   codec.Codec
	ctx   sdk.Context
	store sdk.KVStore
}

func (suite *LocalhostTestSuite) SetupTest() {
	isCheckTx := false
	app := simapp.Setup(isCheckTx)

	suite.cdc = app.AppCodec()
	suite.ctx = app.BaseApp.NewContext(isCheckTx, tmproto.Header{Height: 1, ChainID: "ibc-chain"})
	suite.store = app.IBCKeeper.ClientKeeper.ClientStore(suite.ctx, exported.Localhost)
}

func TestLocalhostTestSuite(t *testing.T) {
	suite.Run(t, new(LocalhostTestSuite))
}
