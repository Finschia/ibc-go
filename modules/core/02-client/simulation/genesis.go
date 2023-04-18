package simulation

import (
	"math/rand"

	simtypes "github.com/Finschia/finschia-sdk/types/simulation"

	"github.com/Finschia/ibc-go/v3/modules/core/02-client/types"
)

// GenClientGenesis returns the default client genesis state.
func GenClientGenesis(_ *rand.Rand, _ []simtypes.Account) types.GenesisState {
	return types.DefaultGenesisState()
}
