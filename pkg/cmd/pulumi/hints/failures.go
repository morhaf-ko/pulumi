// Copyright 2026, Pulumi Corporation.
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

package hints

import (
	"context"

	"github.com/pulumi/pulumi/pkg/v3/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// WithFailureRecorder returns a context whose providers report the hints for their failures, such as dynamic
// credentials, to this process's hints.
func WithFailureRecorder(ctx context.Context) context.Context {
	return plugin.WithHintRecorder(ctx, defaultManager)
}

// RecordHint implements plugin.HintRecorder. It offers the CLI hint for a hint that a provider sent.
func (m *manager) RecordHint(pkg tokens.Package, hint plugin.Hint) {
	if hint != plugin.HintDynamicCredentials {
		return
	}
	if h, ok := escCredentials[pkg]; ok {
		m.offer(h)
	}
}
