// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PkgTypesIntotoV002IntotoV002Schema intoto v0.0.2 Schema
//
// # Schema for intoto object
//
// swagger:model pkgTypesIntotoV002IntotoV002Schema
type PkgTypesIntotoV002IntotoV002Schema struct {

	// content
	// Required: true
	Content *PkgTypesIntotoV002IntotoV002SchemaContent `json:"content"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema
func (m *PkgTypesIntotoV002IntotoV002Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002Schema) validateContent(formats strfmt.Registry) error {

	if err := validate.Required("content", "body", m.Content); err != nil {
		return err
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pkg types intoto v002 intoto v002 schema based on the context it is used
func (m *PkgTypesIntotoV002IntotoV002Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002Schema) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {

		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002Schema) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesIntotoV002IntotoV002SchemaContent pkg types intoto v002 intoto v002 schema content
//
// swagger:model PkgTypesIntotoV002IntotoV002SchemaContent
type PkgTypesIntotoV002IntotoV002SchemaContent struct {

	// envelope
	// Required: true
	Envelope *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope `json:"envelope"`

	// hash
	Hash *PkgTypesIntotoV002IntotoV002SchemaContentHash `json:"hash,omitempty"`

	// payload hash
	PayloadHash *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash `json:"payloadHash,omitempty"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema content
func (m *PkgTypesIntotoV002IntotoV002SchemaContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvelope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayloadHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) validateEnvelope(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"envelope", "body", m.Envelope); err != nil {
		return err
	}

	if m.Envelope != nil {
		if err := m.Envelope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "envelope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "envelope")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) validatePayloadHash(formats strfmt.Registry) error {
	if swag.IsZero(m.PayloadHash) { // not required
		return nil
	}

	if m.PayloadHash != nil {
		if err := m.PayloadHash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pkg types intoto v002 intoto v002 schema content based on the context it is used
func (m *PkgTypesIntotoV002IntotoV002SchemaContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnvelope(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayloadHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) contextValidateEnvelope(ctx context.Context, formats strfmt.Registry) error {

	if m.Envelope != nil {

		if err := m.Envelope.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "envelope")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "envelope")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {

		if swag.IsZero(m.Hash) { // not required
			return nil
		}

		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContent) contextValidatePayloadHash(ctx context.Context, formats strfmt.Registry) error {

	if m.PayloadHash != nil {

		if swag.IsZero(m.PayloadHash) { // not required
			return nil
		}

		if err := m.PayloadHash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "payloadHash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "payloadHash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContent) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002SchemaContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesIntotoV002IntotoV002SchemaContentEnvelope dsse envelope
//
// swagger:model PkgTypesIntotoV002IntotoV002SchemaContentEnvelope
type PkgTypesIntotoV002IntotoV002SchemaContentEnvelope struct {

	// payload of the envelope
	// Format: byte
	Payload strfmt.Base64 `json:"payload,omitempty"`

	// type describing the payload
	// Required: true
	PayloadType *string `json:"payloadType"`

	// collection of all signatures of the envelope's payload
	// Required: true
	// Min Items: 1
	Signatures []*PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0 `json:"signatures"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema content envelope
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePayloadType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) validatePayloadType(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"envelope"+"."+"payloadType", "body", m.PayloadType); err != nil {
		return err
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) validateSignatures(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"envelope"+"."+"signatures", "body", m.Signatures); err != nil {
		return err
	}

	iSignaturesSize := int64(len(m.Signatures))

	if err := validate.MinItems("content"+"."+"envelope"+"."+"signatures", "body", iSignaturesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this pkg types intoto v002 intoto v002 schema content envelope based on the context it is used
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {

			if swag.IsZero(m.Signatures[i]) { // not required
				return nil
			}

			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + "envelope" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelope) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002SchemaContentEnvelope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0 a signature of the envelope's payload along with the public key for the signature
//
// swagger:model PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0
type PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0 struct {

	// optional id of the key used to create the signature
	Keyid string `json:"keyid,omitempty"`

	// public key that corresponds to this signature
	// Required: true
	// Format: byte
	PublicKey *strfmt.Base64 `json:"publicKey"`

	// signature of the payload
	// Required: true
	// Format: byte
	Sig *strfmt.Base64 `json:"sig"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema content envelope signatures items0
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) validatePublicKey(formats strfmt.Registry) error {

	if err := validate.Required("publicKey", "body", m.PublicKey); err != nil {
		return err
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) validateSig(formats strfmt.Registry) error {

	if err := validate.Required("sig", "body", m.Sig); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pkg types intoto v002 intoto v002 schema content envelope signatures items0 based on context it is used
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002SchemaContentEnvelopeSignaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesIntotoV002IntotoV002SchemaContentHash Specifies the hash algorithm and value encompassing the entire signed envelope
//
// swagger:model PkgTypesIntotoV002IntotoV002SchemaContentHash
type PkgTypesIntotoV002IntotoV002SchemaContentHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: ["sha256"]
	Algorithm *string `json:"algorithm"`

	// The hash value for the archive
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema content hash
func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pkgTypesIntotoV002IntotoV002SchemaContentHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pkgTypesIntotoV002IntotoV002SchemaContentHashTypeAlgorithmPropEnum = append(pkgTypesIntotoV002IntotoV002SchemaContentHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// PkgTypesIntotoV002IntotoV002SchemaContentHashAlgorithmSha256 captures enum value "sha256"
	PkgTypesIntotoV002IntotoV002SchemaContentHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pkgTypesIntotoV002IntotoV002SchemaContentHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pkg types intoto v002 intoto v002 schema content hash based on the context it is used
func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentHash) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002SchemaContentHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash Specifies the hash algorithm and value covering the payload within the DSSE envelope
//
// swagger:model PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash
type PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: ["sha256"]
	Algorithm *string `json:"algorithm"`

	// The hash value of the payload
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this pkg types intoto v002 intoto v002 schema content payload hash
func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var pkgTypesIntotoV002IntotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pkgTypesIntotoV002IntotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum = append(pkgTypesIntotoV002IntotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// PkgTypesIntotoV002IntotoV002SchemaContentPayloadHashAlgorithmSha256 captures enum value "sha256"
	PkgTypesIntotoV002IntotoV002SchemaContentPayloadHashAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pkgTypesIntotoV002IntotoV002SchemaContentPayloadHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("content"+"."+"payloadHash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("content"+"."+"payloadHash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this pkg types intoto v002 intoto v002 schema content payload hash based on the context it is used
func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash) UnmarshalBinary(b []byte) error {
	var res PkgTypesIntotoV002IntotoV002SchemaContentPayloadHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
