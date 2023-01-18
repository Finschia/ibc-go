package ostracon

import (
	"github.com/line/ibc-go/v3/modules/light-clients/99-ostracon/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
