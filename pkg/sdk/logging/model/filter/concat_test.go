// Copyright © 2019 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package filter_test

import (
	"testing"

	"github.com/banzaicloud/logging-operator/pkg/sdk/logging/model/filter"
	"github.com/banzaicloud/logging-operator/pkg/sdk/logging/model/render"
	"github.com/ghodss/yaml"
	"github.com/stretchr/testify/require"
)

func TestConcat(t *testing.T) {
	CONFIG := []byte(`
partial_key: "partial_message"
separator: ""
n_lines: 10

`)
	expected := `
<filter **>
  @type concat
  @id test
  key message
  n_lines 10
  partial_key partial_message
</filter>
`
	parser := &filter.Concat{}
	require.NoError(t, yaml.Unmarshal(CONFIG, parser))
	test := render.NewOutputPluginTest(t, parser)
	test.DiffResult(expected)
}
