package simulation

import (
	"math/rand"

	simtypes "github.com/Finschia/finschia-sdk/types/simulation"

	"github.com/Finschia/ibc-go/v3/modules/core/03-connection/types"
)

// GenConnectionGenesis returns the default connection genesis state.
func GenConnectionGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
