//
// Copyright 2022 The Sigstore Authors.
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

//go:build e2e

package main

import (
	"sync"
	"testing"

	"github.com/Morrison76/rekor/pkg/util"
)

var (
	once sync.Once
)

func TestLogInfo(t *testing.T) {
	once.Do(func() {
		util.SetupTestData(t)
	})
	// TODO: figure out some way to check the length, add something, and make sure the length increments!
	out := util.RunCli(t, "loginfo")
	util.OutputContains(t, out, "Verification Successful!")
}
