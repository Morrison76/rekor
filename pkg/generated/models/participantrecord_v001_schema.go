// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ParticipantrecordV001Schema Participantrecord v0.0.1 Schema
//
// # Schema for Participantrecord v0.0.1 object
//
// swagger:model participantrecordV001Schema
type ParticipantrecordV001Schema struct {

	// Alternate public key (optional)
	AlternatePK string `json:"AlternatePK,omitempty"`

	// ID of the participant
	// Required: true
	ParticipantID *string `json:"participantid"`

	// Primary public key
	// Required: true
	PrimaryPK *string `json:"PrimaryPK"`
}

// Validate validates this participantrecord v001 schema
func (m *ParticipantrecordV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParticipantID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryPK(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ParticipantrecordV001Schema) validateParticipantID(formats strfmt.Registry) error {

	if err := validate.Required("ParticipantID", "body", m.ParticipantID); err != nil {
		return err
	}

	return nil
}

func (m *ParticipantrecordV001Schema) validatePrimaryPK(formats strfmt.Registry) error {

	if err := validate.Required("PrimaryPK", "body", m.PrimaryPK); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this participantrecord v001 schema based on context it is used
func (m *ParticipantrecordV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParticipantrecordV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParticipantrecordV001Schema) UnmarshalBinary(b []byte) error {
	var res ParticipantrecordV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
