// Copyright 2019, OpenTelemetry Authors
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

// Package tests contains test cases. To run the tests go to tests directory and run:
// TESTBED_CONFIG=local.yaml go test -v

package tests

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector/testbed/testbed"
	scenarios "github.com/open-telemetry/opentelemetry-collector/testbed/tests"
)

// TestMain is used to initiate setup, execution and tear down of testbed.
func TestMain(m *testing.M) {
	testbed.DoTestMain(m)
}

func TestTrace10kSPS(t *testing.T) {
	tests := []struct {
		name         string
		sender       testbed.DataSender
		receiver     testbed.DataReceiver
		resourceSpec testbed.ResourceSpec
	}{
		{
			"OpenCensus",
			testbed.NewOCTraceDataSender(testbed.GetAvailablePort(t)),
			testbed.NewOCDataReceiver(testbed.GetAvailablePort(t)),
			testbed.ResourceSpec{
				ExpectedMaxCPU: 42,
				ExpectedMaxRAM: 71,
			},
		},
		{
			"SAPM",
			NewSapmDataSender(testbed.GetAvailablePort(t)),
			NewSapmDataReceiver(testbed.GetAvailablePort(t)),
			testbed.ResourceSpec{
				ExpectedMaxCPU: 71,
				ExpectedMaxRAM: 140,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			scenarios.Scenario10kItemsPerSecond(
				t,
				test.sender,
				test.receiver,
				testbed.LoadOptions{},
				test.resourceSpec,
			)
		})
	}
}
