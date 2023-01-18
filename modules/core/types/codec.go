package types

import (
	codectypes "github.com/line/lbm-sdk/codec/types"

	clienttypes "github.com/line/ibc-go/v3/modules/core/02-client/types"
	connectiontypes "github.com/line/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/line/ibc-go/v3/modules/core/04-channel/types"
	commitmenttypes "github.com/line/ibc-go/v3/modules/core/23-commitment/types"
	solomachinetypes "github.com/line/ibc-go/v3/modules/light-clients/06-solomachine/types"
	localhosttypes "github.com/line/ibc-go/v3/modules/light-clients/09-localhost/types"
	ibcoctypes "github.com/line/ibc-go/v3/modules/light-clients/99-ostracon/types"
)

// RegisterInterfaces registers x/ibc interfaces into protobuf Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	clienttypes.RegisterInterfaces(registry)
	connectiontypes.RegisterInterfaces(registry)
	channeltypes.RegisterInterfaces(registry)
	solomachinetypes.RegisterInterfaces(registry)
	ibcoctypes.RegisterInterfaces(registry)
	localhosttypes.RegisterInterfaces(registry)
	commitmenttypes.RegisterInterfaces(registry)
}
