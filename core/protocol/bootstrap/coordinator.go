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

package bootstrap

import "context"

// CoordinatorHandler is the handler for the coordinator
// of a private network.
type CoordinatorHandler struct{}

// NewCoordinatorHandler returns a Handler for a coordinator node.
func NewCoordinatorHandler() (Handler, error) {
	return &CoordinatorHandler{}, nil
}

// Close doesn't do anything.
func (h *CoordinatorHandler) Close(context.Context) {}
