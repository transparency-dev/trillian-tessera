// Copyright 2025 The Tessera authors. All Rights Reserved.
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

// Package antispam provides internal APIs that should not be called or implemented
// by clients.
package antispam

import (
	"github.com/transparency-dev/tessera/core"
	"github.com/transparency-dev/tessera/internal/stream"
)

// Antispam describes the contract that an antispam implementation must meet in order to be used via the
// WithAntispam option.
type Antispam interface {
	// Decorator must return a function which knows how to decorate an Appender's Add function in order
	// to return an index previously assigned to an entry with the same identity hash, if one exists, or
	// delegate to the next Add function in the chain otherwise.
	Decorator() func(core.AddFn) core.AddFn
	// Follower should return a structure which will populate the anti-spam index by tailing the contents
	// of the log, using the provided function to turn entry bundles into identity hashes.
	Follower(func(entryBundle []byte) ([][]byte, error)) stream.Follower
}
