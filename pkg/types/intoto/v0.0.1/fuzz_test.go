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

package intoto

import (
	"context"
	"sync"
	"testing"

	fuzz "github.com/AdamKorcz/go-fuzz-headers-1"
	"github.com/go-openapi/swag"

	fuzzUtils "github.com/Morrison76/rekor/pkg/fuzz"
	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/types"
	"github.com/Morrison76/rekor/pkg/types/intoto"
)

var initter sync.Once

func FuzzIntotoCreateProposedEntry(f *testing.F) {
	f.Fuzz(func(t *testing.T, propsData []byte) {
		initter.Do(fuzzUtils.SetFuzzLogger)

		version := "0.0.1"

		ff := fuzz.NewConsumer(propsData)

		props, cleanup, err := fuzzUtils.CreateProps(ff, "intotoV001")
		if err != nil {
			t.Skip()
		}
		defer func() {
			for _, c := range cleanup {
				c()
			}
		}()

		it := intoto.New()
		entry, err := it.CreateProposedEntry(context.Background(), version, props)
		if err != nil {
			t.Skip()
		}
		ei, err := types.CreateVersionedEntry(entry)
		if err != nil {
			t.Skip()
		}

		if ok, err := ei.Insertable(); !ok || err != nil {
			t.Errorf("entry created via CreateProposedEntry should be insertable: %v", err)
		}

		if _, err := types.CanonicalizeEntry(context.Background(), ei); err != nil {
			t.Errorf("valid insertable entry should be able to be canonicalized: %v", err)
		}

		_, _ = ei.IndexKeys()
	})
}

func FuzzIntotoUnmarshalAndCanonicalize(f *testing.F) {
	f.Fuzz(func(t *testing.T, entryData []byte) {
		initter.Do(fuzzUtils.SetFuzzLogger)

		ff := fuzz.NewConsumer(entryData)

		targetV001 := &models.IntotoV001Schema{}

		if err := ff.GenerateStruct(targetV001); err != nil {
			t.Skip()
		}

		targetEntry := &models.Intoto{
			APIVersion: swag.String(APIVERSION),
			Spec:       targetV001,
		}

		ei, err := types.UnmarshalEntry(targetEntry)
		if err != nil {
			t.Skip()
		}

		if _, err := types.CanonicalizeEntry(context.Background(), ei); err != nil {
			t.Errorf("error canonicalizing unmarshalled entry: %v", err)
		}
	})
}
