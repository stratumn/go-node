// Copyright © 2017-2018 Stratumn SAS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protector

import (
	"context"
	"io/ioutil"

	json "github.com/gibson042/canonicaljson-go"
	"github.com/pkg/errors"
	cryptopb "github.com/stratumn/alice/pb/crypto"

	logging "gx/ipfs/QmSpJByNKFX1sCsHBEp3R73FL4NF6FnQTEGyNAXHm2GS52/go-log"
	"gx/ipfs/QmcJukH2sAFjY3HdBKq35WDzWoL3UUu2gt9wdfqZTUyM74/go-libp2p-peer"
	"gx/ipfs/Qme1knMqwt1hKZbc1BmQFmnm9f36nyQGwXxPGVpVJ9rMK5/go-libp2p-crypto"
)

var log = logging.Logger("bootstrap")

var (
	// ErrInvalidConfig is returned when an invalid configuration file
	// is provided.
	ErrInvalidConfig = errors.New("invalid configuration file")

	// ErrInvalidSignature is returned when an existing configuration file
	// contains an invalid signature.
	ErrInvalidSignature = errors.New("invalid configuration signature")

	// ErrInvalidNetworkState is returned when trying to set an
	// invalid network state.
	ErrInvalidNetworkState = errors.New("invalid network state")
)

// NewNetworkConfig creates a signed NetworkConfig.
func NewNetworkConfig(networkState NetworkState) *NetworkConfig {
	networkConfig := &NetworkConfig{
		NetworkState: networkState,
		Participants: make(map[string]*PeerAddrs),
	}

	return networkConfig
}

// Sign signs the current configuration data.
func (c *NetworkConfig) Sign(ctx context.Context, privKey crypto.PrivKey) error {
	event := log.EventBegin(ctx, "NetworkConfig.Sign")
	defer event.Done()

	c.Signature = nil

	b, err := json.Marshal(c)
	if err != nil {
		event.SetError(err)
		return errors.WithStack(err)
	}

	c.Signature, err = cryptopb.Sign(privKey, b)
	if err != nil {
		event.SetError(err)
		return errors.WithStack(err)
	}

	return nil
}

// ValidateSignature verifies that the network configuration
// was signed by a given peer.
func (c *NetworkConfig) ValidateSignature(ctx context.Context, peerID peer.ID) bool {
	if c == nil {
		return false
	}

	if c.Signature == nil {
		return false
	}

	pk, err := crypto.UnmarshalPublicKey(c.Signature.PublicKey)
	if err != nil {
		return false
	}

	if !peerID.MatchesPublicKey(pk) {
		return false
	}

	signature := c.Signature
	c.Signature = nil

	payload, err := json.Marshal(c)
	c.Signature = signature
	if err != nil {
		return false
	}

	return c.Signature.Verify(payload)
}

// LoadFromFile loads the given configuration and validates it.
func (c *NetworkConfig) LoadFromFile(ctx context.Context, path string, signerID peer.ID) (err error) {
	event := log.EventBegin(ctx, "NetworkConfig.Load", logging.Metadata{"path": path})
	defer func() {
		if err != nil {
			event.SetError(err)
		}

		event.Done()
	}()

	configBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return ErrInvalidConfig
	}

	var confData NetworkConfig
	err = json.Unmarshal(configBytes, &confData)
	if err != nil {
		return ErrInvalidConfig
	}

	if !confData.ValidateSignature(ctx, signerID) {
		return ErrInvalidSignature
	}

	_, ok := NetworkState_name[int32(confData.NetworkState)]
	if !ok {
		return ErrInvalidNetworkState
	}

	c.NetworkState = confData.NetworkState
	c.Participants = confData.Participants
	c.Signature = confData.Signature

	return nil
}

// SaveToFile signs the configuration data and writes it to disk.
func (c *NetworkConfig) SaveToFile(ctx context.Context, path string, privKey crypto.PrivKey) error {
	event := log.EventBegin(ctx, "NetworkConfig.Save", logging.Metadata{"path": path})
	defer event.Done()

	err := c.Sign(ctx, privKey)
	if err != nil {
		event.SetError(err)
		return err
	}

	signedBytes, err := json.Marshal(c)
	if err != nil {
		event.SetError(err)
		return errors.WithStack(err)
	}

	err = ioutil.WriteFile(path, signedBytes, 0644)
	if err != nil {
		event.SetError(err)
	}

	return err
}