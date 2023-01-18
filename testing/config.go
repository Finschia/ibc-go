package ibctesting

import (
	"time"

	connectiontypes "github.com/line/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/line/ibc-go/v3/modules/core/04-channel/types"
	"github.com/line/ibc-go/v3/modules/core/exported"
	ibcoctypes "github.com/line/ibc-go/v3/modules/light-clients/99-ostracon/types"
	"github.com/line/ibc-go/v3/testing/mock"
)

type ClientConfig interface {
	GetClientType() string
}

type OstraconConfig struct {
	TrustLevel                   ibcoctypes.Fraction
	TrustingPeriod               time.Duration
	UnbondingPeriod              time.Duration
	MaxClockDrift                time.Duration
	AllowUpdateAfterExpiry       bool
	AllowUpdateAfterMisbehaviour bool
}

func NewOstraconConfig() *OstraconConfig {
	return &OstraconConfig{
		TrustLevel:                   DefaultTrustLevel,
		TrustingPeriod:               TrustingPeriod,
		UnbondingPeriod:              UnbondingPeriod,
		MaxClockDrift:                MaxClockDrift,
		AllowUpdateAfterExpiry:       false,
		AllowUpdateAfterMisbehaviour: false,
	}
}

func (tmcfg *OstraconConfig) GetClientType() string {
	return exported.Ostracon
}

type ConnectionConfig struct {
	DelayPeriod uint64
	Version     *connectiontypes.Version
}

func NewConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		DelayPeriod: DefaultDelayPeriod,
		Version:     ConnectionVersion,
	}
}

type ChannelConfig struct {
	PortID  string
	Version string
	Order   channeltypes.Order
}

func NewChannelConfig() *ChannelConfig {
	return &ChannelConfig{
		PortID:  mock.PortID,
		Version: DefaultChannelVersion,
		Order:   channeltypes.UNORDERED,
	}
}
