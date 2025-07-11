// Code generated by go-swagger; DO NOT EDIT.



package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
  "github.com/go-openapi/strfmt"
  	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsistencyProof consistency proof
//
// swagger:model ConsistencyProof
      type ConsistencyProof struct {
  
  
    // hashes
// Required: true
Hashes []string `json:"hashes"`

  
    // The hash value stored at the root of the merkle tree at the time the proof was generated
// Required: true
// Pattern: ^[0-9a-fA-F]{64}$
RootHash *string `json:"rootHash"`

  
  
}
  
    
  
  
  
// Validate validates this consistency proof
func (m *ConsistencyProof) Validate(formats strfmt.Registry) error {
  var res []error
  
  
  

  
    
      if err := m.validateHashes(formats); err != nil {
        res = append(res, err)
      }
    
  
    
      if err := m.validateRootHash(formats); err != nil {
        res = append(res, err)
      }
    
  
  
  

  if len(res) > 0 {
    return errors.CompositeValidationError(res...)
  }
  return nil
}

  
    
      
      
      
      

      
func (m *ConsistencyProof) validateHashes(formats strfmt.Registry) error {
        
    
  
    if err := validate.Required("hashes", "body", m.Hashes); err != nil {
      return err
    }
  
  
  
  
  
  
  
      for i := 0; i < len(m.Hashes); i++ {
        
      
  
  
  
  
  if err := validate.Pattern("hashes"+ "." + strconv.Itoa(i), "body", m.Hashes[i], `^[0-9a-fA-F]{64}$`); err != nil {
    return err
  }
  
  
  
  
  


      }



  return nil
}
      
    
  
    
      
      
      
      

      
func (m *ConsistencyProof) validateRootHash(formats strfmt.Registry) error {
        
      
  
  if err := validate.Required("rootHash", "body", m.RootHash); err != nil {
    return err
  }
  
  
  
  if err := validate.Pattern("rootHash", "body", *m.RootHash, `^[0-9a-fA-F]{64}$`); err != nil {
    return err
  }
  
  
  
  
  



  return nil
}
      
    
  
  

  

// ContextValidate validates this consistency proof based on context it is used 
func (m *ConsistencyProof) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
  return nil
}
  
// MarshalBinary interface implementation
func (m *ConsistencyProof) MarshalBinary() ([]byte, error) {
  if m == nil {
    return nil, nil
  }
  return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyProof) UnmarshalBinary(b []byte) error {
  var res ConsistencyProof
  if err := swag.ReadJSON(b, &res); err != nil {
    return err
  }
  *m = res
  return nil
}



