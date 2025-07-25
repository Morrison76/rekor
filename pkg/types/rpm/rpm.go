//
// Copyright 2021 The Sigstore Authors.
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

package rpm

import (
	"context"
	"errors"
	"fmt"

	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/types"
)

const (
	KIND = "rpm"
)

type BaseRPMType struct {
	types.RekorType
}

func init() {
	types.TypeMap.Store(KIND, New)
}

func New() types.TypeImpl {
	brt := BaseRPMType{}
	brt.Kind = KIND
	brt.VersionMap = VersionMap
	return &brt
}

var VersionMap = types.NewSemVerEntryFactoryMap()

func (brt *BaseRPMType) UnmarshalEntry(pe models.ProposedEntry) (types.EntryImpl, error) {
	if pe == nil {
		return nil, errors.New("proposed entry cannot be nil")
	}

	rpm, ok := pe.(*models.Rpm)
	if !ok {
		return nil, errors.New("cannot unmarshal non-RPM types")
	}

	return brt.VersionedUnmarshal(rpm, *rpm.APIVersion)
}

func (brt *BaseRPMType) CreateProposedEntry(ctx context.Context, version string, props types.ArtifactProperties) (models.ProposedEntry, error) {
	if version == "" {
		version = brt.DefaultVersion()
	}
	ei, err := brt.VersionedUnmarshal(nil, version)
	if err != nil {
		return nil, fmt.Errorf("fetching RPM version implementation: %w", err)
	}
	return ei.CreateFromArtifactProperties(ctx, props)
}

func (brt BaseRPMType) DefaultVersion() string {
	return "0.0.1"
}
