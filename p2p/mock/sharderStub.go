package mock

import (
	"github.com/DharitriOne/drt-chain-communication-go/p2p"
	"github.com/DharitriOne/drt-chain-core-go/core"
)

// SharderStub -
type SharderStub struct {
	SetPeerShardResolverCalled func(psp p2p.PeerShardResolver) error
	SetSeedersCalled           func(addresses []string)
	IsSeederCalled             func(pid core.PeerID) bool
}

// SetPeerShardResolver -
func (ss *SharderStub) SetPeerShardResolver(psp p2p.PeerShardResolver) error {
	if ss.SetPeerShardResolverCalled != nil {
		return ss.SetPeerShardResolverCalled(psp)
	}

	return nil
}

// SetSeeders -
func (ss *SharderStub) SetSeeders(addresses []string) {
	if ss.SetSeedersCalled != nil {
		ss.SetSeedersCalled(addresses)
	}
}

// IsSeeder -
func (ss *SharderStub) IsSeeder(pid core.PeerID) bool {
	if ss.IsSeederCalled != nil {
		return ss.IsSeederCalled(pid)
	}

	return false
}

// IsInterfaceNil -
func (ss *SharderStub) IsInterfaceNil() bool {
	return ss == nil
}
