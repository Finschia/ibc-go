package simulation

import (
	"fmt"

	"github.com/Finschia/finschia-sdk/types/kv"

	clientsim "github.com/Finschia/ibc-go/v3/modules/core/02-client/simulation"
	connectionsim "github.com/Finschia/ibc-go/v3/modules/core/03-connection/simulation"
	channelsim "github.com/Finschia/ibc-go/v3/modules/core/04-channel/simulation"
	host "github.com/Finschia/ibc-go/v3/modules/core/24-host"
	"github.com/Finschia/ibc-go/v3/modules/core/keeper"
)

// NewDecodeStore returns a decoder function closure that unmarshals the KVPair's
// Value to the corresponding ibc type.
func NewDecodeStore(k keeper.Keeper) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		if res, found := clientsim.NewDecodeStore(k.ClientKeeper, kvA, kvB); found {
			return res
		}

		if res, found := connectionsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		if res, found := channelsim.NewDecodeStore(k.Codec(), kvA, kvB); found {
			return res
		}

		panic(fmt.Sprintf("invalid %s key prefix: %s", host.ModuleName, string(kvA.Key)))
	}
}
