package discovery_test

import (
	"testing"

	"github.com/DharitriOne/drt-chain-communication-go/p2p/libp2p/discovery"
	"github.com/DharitriOne/drt-chain-core-go/core/check"
	"github.com/stretchr/testify/assert"
)

func TestNilDiscoverer(t *testing.T) {
	t.Parallel()

	nd := discovery.NewNilDiscoverer()

	assert.False(t, check.IfNil(nd))
	assert.Equal(t, discovery.NullName, nd.Name())
	assert.Nil(t, nd.Bootstrap())
}
