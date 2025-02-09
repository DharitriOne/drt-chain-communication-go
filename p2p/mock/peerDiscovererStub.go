package mock

// PeerDiscovererStub -
type PeerDiscovererStub struct {
	BootstrapCalled func() error
	CloseCalled     func() error
}

// Bootstrap -
func (pds *PeerDiscovererStub) Bootstrap() error {
	if pds.BootstrapCalled != nil {
		return pds.BootstrapCalled()
	}
	return nil
}

// Name -
func (pds *PeerDiscovererStub) Name() string {
	return "PeerDiscovererStub"
}

// IsInterfaceNil returns true if there is no value under the interface
func (pds *PeerDiscovererStub) IsInterfaceNil() bool {
	return pds == nil
}
