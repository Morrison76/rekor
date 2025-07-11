// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PkgTypesHashedrekordV001HashedrekordV001Schema Hashed Rekor v0.0.1 Schema
//
// # Schema for Hashed Rekord object
//
// swagger:model pkgTypesHashedrekordV001HashedrekordV001Schema
type PkgTypesHashedrekordV001HashedrekordV001Schema struct {

	// data
	// Required: true
	Data *PkgTypesHashedrekordV001HashedrekordV001SchemaData `json:"data"`

	// signature
	// Required: true
	Signature *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature `json:"signature"`
}

// Validate validates this pkg types hashedrekord v001 hashedrekord v001 schema
func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) validateSignature(formats strfmt.Registry) error {

	if err := validate.Required("signature", "body", m.Signature); err != nil {
		return err
	}

	if m.Signature != nil {
		if err := m.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pkg types hashedrekord v001 hashedrekord v001 schema based on the context it is used
func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if m.Data != nil {

		if err := m.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if m.Signature != nil {

		if err := m.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001Schema) UnmarshalBinary(b []byte) error {
	var res PkgTypesHashedrekordV001HashedrekordV001Schema
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesHashedrekordV001HashedrekordV001SchemaData Information about the content associated with the entry
//
// swagger:model PkgTypesHashedrekordV001HashedrekordV001SchemaData
type PkgTypesHashedrekordV001HashedrekordV001SchemaData struct {

	// hash
	Hash *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash `json:"hash,omitempty"`
}

// Validate validates this pkg types hashedrekord v001 hashedrekord v001 schema data
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) validateHash(formats strfmt.Registry) error {
	if swag.IsZero(m.Hash) { // not required
		return nil
	}

	if m.Hash != nil {
		if err := m.Hash.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pkg types hashedrekord v001 hashedrekord v001 schema data based on the context it is used
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHash(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) contextValidateHash(ctx context.Context, formats strfmt.Registry) error {

	if m.Hash != nil {

		if swag.IsZero(m.Hash) { // not required
			return nil
		}

		if err := m.Hash.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data" + "." + "hash")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("data" + "." + "hash")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaData) UnmarshalBinary(b []byte) error {
	var res PkgTypesHashedrekordV001HashedrekordV001SchemaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash Specifies the hash algorithm and value for the content
//
// swagger:model PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash
type PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash struct {

	// The hashing function used to compute the hash value
	// Required: true
	// Enum: ["sha256","sha384","sha512"]
	Algorithm *string `json:"algorithm"`

	// The hash value for the content, as represented by a lower case hexadecimal string
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this pkg types hashedrekord v001 hashedrekord v001 schema data hash
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) Validate(formats strfmt.Registry) error {
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

var pkgTypesHashedrekordV001HashedrekordV001SchemaDataHashTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sha256","sha384","sha512"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pkgTypesHashedrekordV001HashedrekordV001SchemaDataHashTypeAlgorithmPropEnum = append(pkgTypesHashedrekordV001HashedrekordV001SchemaDataHashTypeAlgorithmPropEnum, v)
	}
}

const (

	// PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha256 captures enum value "sha256"
	PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha256 string = "sha256"

	// PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha384 captures enum value "sha384"
	PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha384 string = "sha384"

	// PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha512 captures enum value "sha512"
	PkgTypesHashedrekordV001HashedrekordV001SchemaDataHashAlgorithmSha512 string = "sha512"
)

// prop value enum
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pkgTypesHashedrekordV001HashedrekordV001SchemaDataHashTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("data"+"."+"hash"+"."+"algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"hash"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this pkg types hashedrekord v001 hashedrekord v001 schema data hash based on context it is used
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash) UnmarshalBinary(b []byte) error {
	var res PkgTypesHashedrekordV001HashedrekordV001SchemaDataHash
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesHashedrekordV001HashedrekordV001SchemaSignature Information about the detached signature associated with the entry
//
// swagger:model PkgTypesHashedrekordV001HashedrekordV001SchemaSignature
type PkgTypesHashedrekordV001HashedrekordV001SchemaSignature struct {

	// Specifies the content of the signature inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`

	// public key
	PublicKey *PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey `json:"publicKey,omitempty"`
}

// Validate validates this pkg types hashedrekord v001 hashedrekord v001 schema signature
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this pkg types hashedrekord v001 hashedrekord v001 schema signature based on the context it is used
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {

		if swag.IsZero(m.PublicKey) { // not required
			return nil
		}

		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signature" + "." + "publicKey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signature" + "." + "publicKey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignature) UnmarshalBinary(b []byte) error {
	var res PkgTypesHashedrekordV001HashedrekordV001SchemaSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey The public key that can verify the signature; this can also be an X509 code signing certificate that contains the raw public key information
//
// swagger:model PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey
type PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey struct {

	// Specifies the content of the public key or code signing certificate inline within the document
	// Format: byte
	Content strfmt.Base64 `json:"content,omitempty"`
}

// Validate validates this pkg types hashedrekord v001 hashedrekord v001 schema signature public key
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pkg types hashedrekord v001 hashedrekord v001 schema signature public key based on context it is used
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey) UnmarshalBinary(b []byte) error {
	var res PkgTypesHashedrekordV001HashedrekordV001SchemaSignaturePublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
