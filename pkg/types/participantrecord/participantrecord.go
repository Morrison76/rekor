package participantrecord

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/sigstore/rekor/pkg/generated/models"
	"github.com/sigstore/rekor/pkg/types"
)

const Kind = "participantrecord"
const APIVersion = "0.0.1"

// ==== Тип, який буде зареєстровано у TypeMap ====

type Type struct{}

func init() {
	types.TypeMap.Store(Kind, &Type{})
}

// ==== Конкретна реалізація EntryImpl ====

type Entry struct {
	ParticipantID string `json:"participantId"`
	PrimaryPK     string `json:"primaryPK"`
	AlternatePK   string `json:"alternatePK,omitempty"`
}

// ==== Реалізація TypeImpl ====

func (t *Type) CreateProposedEntry(_ context.Context, _ string, _ types.ArtifactProperties) (models.ProposedEntry, error) {
	return nil, errors.New("CreateProposedEntry not supported for participantrecord")
}

func (t *Type) DefaultVersion() string {
	return APIVersion
}

func (t *Type) SupportedVersions() []string {
	return []string{APIVersion}
}

func (t *Type) IsSupportedVersion(version string) bool {
	return version == APIVersion
}

func (t *Type) UnmarshalEntry(pe models.ProposedEntry) (types.EntryImpl, error) {
	obj, ok := pe.(*models.Participantrecord)
	if !ok {
		return nil, errors.New("invalid proposed entry type for participantrecord")
	}

	if obj.Spec == nil {
		return nil, errors.New("missing spec field")
	}

	entry := &Entry{}
	if err := types.DecodeEntry(obj.Spec, entry); err != nil {
		return nil, fmt.Errorf("error decoding entry: %w", err)
	}
	if err := entry.Validate(); err != nil {
		return nil, fmt.Errorf("entry validation failed: %w", err)
	}
	return entry, nil
}

// ==== Реалізація EntryImpl ====

func (e *Entry) ArtifactHash() (string, error) {
	return "", errors.New("participantrecord does not support ArtifactHash")
}

func (e *Entry) APIVersion() string {
	return APIVersion
}

func (e *Entry) IndexKeys() ([]string, error) {
	return []string{
		fmt.Sprintf("participantId:%s", e.ParticipantID),
		fmt.Sprintf("primaryPK:%s", e.PrimaryPK),
	}, nil
}

func (e *Entry) Canonicalize(_ context.Context) ([]byte, error) {
	canonical := map[string]string{
		"participantId": e.ParticipantID,
		"primaryPK":     e.PrimaryPK,
	}
	if e.AlternatePK != "" {
		canonical["alternatePK"] = e.AlternatePK
	}
	return json.Marshal(canonical)
}

func (e *Entry) Unmarshal(pe models.ProposedEntry) error {
	obj, ok := pe.(*models.Participantrecord)
	if !ok {
		return errors.New("invalid proposed entry type for participantrecord")
	}
	if obj.Spec == nil {
		return errors.New("missing spec field")
	}
	return types.DecodeEntry(obj.Spec, e)
}

func (e *Entry) CreateFromArtifactProperties(_ context.Context, _ types.ArtifactProperties) (models.ProposedEntry, error) {
	return nil, errors.New("not supported for participantrecord")
}

func (e *Entry) Verifier() (interface{}, error) {
	return nil, nil // ми не використовуємо публічний ключ як PKI.Verifier
}

func (e *Entry) Insertable() (bool, error) {
	if err := e.Validate(); err != nil {
		return false, err
	}
	return true, nil
}

func (e *Entry) HasExternalEntities() bool {
	return false
}

func (e *Entry) FetchExternalEntities(_ context.Context) error {
	return nil
}

func (e *Entry) Validate() error {
	if e.ParticipantID == "" {
		return errors.New("participantId is required")
	}
	if e.PrimaryPK == "" {
		return errors.New("primaryPK is required")
	}
	return nil
}
