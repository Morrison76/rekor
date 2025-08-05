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
	"github.com/Morrison76/rekor/pkg/types/participantrecord"
	"github.com/Morrison76/rekor/pkg/pki"
)

const APIVERSION = "0.0.1"

func init() {
	if err := participantrecord.VersionMap.SetEntryFactory(APIVERSION, NewEntry); err != nil {
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
	if v.Obj.ParticipantID == nil {
		return nil, fmt.Errorf("ParticipantID is nil")
	}
	return []string{
		fmt.Sprintf("participantid:%s", strings.ToLower(*v.Obj.ParticipantID)),
	}, nil
}

func (v *V001Entry) Canonicalize(_ context.Context) ([]byte, error) {
	canonicalEntry := models.ParticipantrecordV001Schema{
		ParticipantID:   v.Obj.ParticipantID,
		PrimaryPK:       v.Obj.PrimaryPK,
		AlternatePK:     v.Obj.AlternatePK,
		EntryCreatedAt:  v.Obj.EntryCreatedAt,
		EncryptionType:  v.Obj.EncryptionType,
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

const (
	EncryptionTypeDefault = "DEFAULT"
	EncryptionTypeRSA     = "RSA"
)

var allowedEncryptionTypes = []string{EncryptionTypeDefault, EncryptionTypeRSA}

func (v *V001Entry) validate() error {
	if strings.TrimSpace(swag.StringValue(v.Obj.ParticipantID)) == "" {
		return errors.New("missing participantId")
	}

	encType := strings.TrimSpace(v.Obj.EncryptionType)
	isAllowed := false
	for _, t := range allowedEncryptionTypes {
		if strings.EqualFold(encType, t) {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return fmt.Errorf("invalid encryptionType '%s', allowed values: %v", encType, allowedEncryptionTypes)
	}

	if strings.EqualFold(encType, EncryptionTypeRSA) {
		if strings.TrimSpace(swag.StringValue(v.Obj.PrimaryPK)) == "" {
			return errors.New("missing primaryPK for RSA encryption type")
		}
	}

	return nil
}


func (v *V001Entry) CreateFromArtifactProperties(_ context.Context, _ types.ArtifactProperties) (models.ProposedEntry, error) {
	return nil, errors.New("not supported for participantrecord")
}

func (v *V001Entry) Verifiers() ([]pki.PublicKey, error) {
	return nil, nil
}

func (v *V001Entry) ArtifactHash() (string, error) {
	return "", nil
}

func (v *V001Entry) Insertable() (bool, error) {
	return true, nil
}
