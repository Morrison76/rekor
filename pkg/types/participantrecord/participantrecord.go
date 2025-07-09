package participantrecord

import (
	"context"
	"errors"
	"fmt"

	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/types"
)

const (
	KIND        = "participantrecord"
	APIVersion  = "0.0.1"
)

type BaseParticipantRecordType struct {
	types.RekorType
}

func init() {
	types.TypeMap.Store(KIND, New)
}

func New() types.TypeImpl {
	ptr := BaseParticipantRecordType{}
	ptr.Kind = KIND
	ptr.VersionMap = VersionMap
	return &ptr
}

var VersionMap = types.NewSemVerEntryFactoryMap()

func (pt BaseParticipantRecordType) UnmarshalEntry(pe models.ProposedEntry) (types.EntryImpl, error) {
	if pe == nil {
		return nil, errors.New("proposed entry cannot be nil")
	}

	entry, ok := pe.(*models.Participantrecord)
	if !ok {
		return nil, errors.New("cannot unmarshal non-participantrecord types")
	}

	return pt.VersionedUnmarshal(entry, *entry.APIVersion)
}

func (pt *BaseParticipantRecordType) CreateProposedEntry(ctx context.Context, version string, props types.ArtifactProperties) (models.ProposedEntry, error) {
	if version == "" {
		version = pt.DefaultVersion()
	}
	ei, err := pt.VersionedUnmarshal(nil, version)
	if err != nil {
		return nil, fmt.Errorf("fetching participantrecord version implementation: %w", err)
	}
	return ei.CreateFromArtifactProperties(ctx, props)
}

func (pt BaseParticipantRecordType) DefaultVersion() string {
	return APIVersion
}
