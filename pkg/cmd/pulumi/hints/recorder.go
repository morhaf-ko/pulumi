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
	"github.com/pulumi/pulumi/pkg/v3/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Recorder returns the recorder that turns the hints providers report for their failures, such as dynamic
// credentials, into this process's hints. Commands set it as the HintRecorder of the plugin contexts they create.
func Recorder() plugin.HintRecorder { return defaultManager }

// RecordHint implements plugin.HintRecorder. It offers the CLI hint for a hint that a provider sent.
func (m *manager) RecordHint(pkg tokens.Package, hint plugin.Hint) {
	if hint != plugin.HintDynamicCredentials {
		return
	}
	if h, ok := escCredentials[pkg]; ok {
		m.offer(h)
	}
}
