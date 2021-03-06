// Copyright © 2017-2018 Stratumn SAS
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package protector

import (
	"context"

	"github.com/stratumn/go-node/core/monitoring"
	"github.com/stratumn/go-node/core/protector/pb"

	"github.com/libp2p/go-libp2p-peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p-peerstore"
)

// ProtectUpdater wraps a NetworkConfig implementation and updates a
// protector when the configuration changes.
type ProtectUpdater struct {
	NetworkConfig
	peerStore   peerstore.Peerstore
	protect     Protector
	protectChan chan NetworkUpdate
}

// WrapWithProtectUpdater wraps a NetworkConfig implementation and updates
// the given protector when the configuration changes.
func WrapWithProtectUpdater(
	networkConfig NetworkConfig,
	protect Protector,
	peerStore peerstore.Peerstore,
) NetworkConfig {
	conf := ProtectUpdater{
		NetworkConfig: networkConfig,
		peerStore:     peerStore,
		protect:       protect,
		protectChan:   make(chan NetworkUpdate),
	}

	// This go routine has the same lifetime as the ProtectUpdater object,
	// so it makes sense to launch it here. When the ProtectUpdater object
	// is collected by the GC, the channel is closed which stops
	// this go routine.
	go protect.ListenForUpdates(conf.protectChan)

	return &conf
}

// AddPeer adds a peer to the network configuration
// and updates the protector and peer store.
func (c *ProtectUpdater) AddPeer(ctx context.Context, peerID peer.ID, addrs []multiaddr.Multiaddr) error {
	ctx, span := monitoring.StartSpan(ctx, "protector.config_protect", "AddPeer")
	defer span.End()

	err := c.NetworkConfig.AddPeer(ctx, peerID, addrs)
	if err != nil {
		return err
	}

	c.peerStore.AddAddrs(peerID, addrs, peerstore.PermanentAddrTTL)
	c.protectChan <- CreateAddNetworkUpdate(peerID)

	return nil
}

// RemovePeer removes a peer from the network configuration
// and updates the protector.
func (c *ProtectUpdater) RemovePeer(ctx context.Context, peerID peer.ID) error {
	ctx, span := monitoring.StartSpan(ctx, "protector.config_protect", "RemovePeer")
	defer span.End()

	err := c.NetworkConfig.RemovePeer(ctx, peerID)
	if err != nil {
		return err
	}

	c.protectChan <- CreateRemoveNetworkUpdate(peerID)

	return nil
}

// SetNetworkState sets the current state of the network protection
// and updates the protector if it's interested in state changes.
func (c *ProtectUpdater) SetNetworkState(ctx context.Context, networkState pb.NetworkState) error {
	ctx, span := monitoring.StartSpan(ctx, "protector.config_protect", "SetNetworkState")
	defer span.End()

	err := c.NetworkConfig.SetNetworkState(ctx, networkState)
	if err != nil {
		return err
	}

	stateAwareProtector, ok := c.protect.(NetworkStateWriter)
	if ok {
		err := stateAwareProtector.SetNetworkState(ctx, networkState)
		if err != nil {
			span.SetUnknownError(err)
			return err
		}
	}

	return nil
}

// Reset clears the current configuration and applies the given one.
// It assumes that the incoming configuration signature has been validated.
// It updates the protector accordingly.
func (c *ProtectUpdater) Reset(ctx context.Context, networkConfig *pb.NetworkConfig) error {
	ctx, span := monitoring.StartSpan(ctx, "protector.config_protect", "Reset")
	defer span.End()

	err := c.NetworkConfig.Reset(ctx, networkConfig)
	if err != nil {
		return err
	}

	stateAwareProtector, ok := c.protect.(NetworkStateWriter)
	if ok {
		err := stateAwareProtector.SetNetworkState(ctx, networkConfig.NetworkState)
		if err != nil {
			span.SetUnknownError(err)
			return err
		}
	}

	for _, peerID := range c.protect.AllowedPeers(ctx) {
		c.protectChan <- CreateRemoveNetworkUpdate(peerID)
	}

	for _, peerID := range c.AllowedPeers(ctx) {
		c.protectChan <- CreateAddNetworkUpdate(peerID)
	}

	return nil
}
