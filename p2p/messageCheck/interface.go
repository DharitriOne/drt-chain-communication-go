package messagecheck

import "github.com/DharitriOne/drt-chain-core-go/core"

type p2pSigner interface {
	Verify(payload []byte, pid core.PeerID, signature []byte) error
	IsInterfaceNil() bool
}
