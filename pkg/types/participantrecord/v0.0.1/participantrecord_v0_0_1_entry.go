package participantrecord

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/go-openapi/swag"

	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/types"
)

const APIVERSION = "0.0.1"

func init() {
	if err := VersionMap.SetEntryFactory(APIVERSION, NewEntry); err != nil {
		panic(err)
	}
}

type V001Entry struct {
	Obj models.ParticipantrecordV001Schema
}

func NewEntry() types.EntryImpl {
	return &V001Entry{}
}

func (v *V001Entry) APIVersion() string {
	return APIVERSION
}

func (v *V001Entry) IndexKeys() ([]string, error) {
	return []string{
		fmt.Sprintf("participantId:%s", v.Obj.ParticipantId),
		fmt.Sprintf("primaryPK:%s", v.Obj.PrimaryPk),
	}, nil
}

func (v *V001Entry) Canonicalize(_ context.Context) ([]byte, error) {
	// canonical entry
	canonicalEntry := models.ParticipantrecordV001Schema{
		ParticipantId: v.Obj.ParticipantId,
		PrimaryPk:     v.Obj.PrimaryPk,
		AlternatePk:   v.Obj.AlternatePk,
	}

	obj := models.Participantrecord{
		APIVersion: swag.String(APIVERSION),
		Spec:       &canonicalEntry,
	}

	return json.Marshal(&obj)
}

func (v *V001Entry) Unmarshal(pe models.ProposedEntry) error {
	obj, ok := pe.(*models.Participantrecord)
	if !ok {
		return errors.New("cannot unmarshal non participantrecord type")
	}

	if err := types.DecodeEntry(obj.Spec, &v.Obj); err != nil {
		return err
	}

	return v.validate()
}

func (v *V001Entry) validate() error {
	if strings.TrimSpace(v.Obj.ParticipantId) == "" {
		return errors.New("missing participantId")
	}
	if strings.TrimSpace(v.Obj.PrimaryPk) == "" {
		return errors.New("missing primaryPK")
	}
	return nil
}

func (v *V001Entry) CreateFromArtifactProperties(_ context.Context, _ types.ArtifactProperties) (models.ProposedEntry, error) {
	return nil, errors.New("not supported for participantrecord")
}

func (v *V001Entry) Verifiers() ([]types.PublicKey, error) {
	return nil, nil
}

func (v *V001Entry) ArtifactHash() (string, error) {
	return "", nil
}

func (v *V001Entry) Insertable() (bool, error) {
	return true, nil
}
