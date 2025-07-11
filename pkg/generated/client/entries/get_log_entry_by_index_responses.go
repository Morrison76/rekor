// Code generated by go-swagger; DO NOT EDIT.

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
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Morrison76/rekor/pkg/generated/models"
)

// GetLogEntryByIndexReader is a Reader for the GetLogEntryByIndex structure.
type GetLogEntryByIndexReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogEntryByIndexReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogEntryByIndexOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogEntryByIndexNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetLogEntryByIndexDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLogEntryByIndexOK creates a GetLogEntryByIndexOK with default headers values
func NewGetLogEntryByIndexOK() *GetLogEntryByIndexOK {
	return &GetLogEntryByIndexOK{}
}

/*
GetLogEntryByIndexOK describes a response with status code 200, with default header values.

the entry in the transparency log requested along with an inclusion proof
*/
type GetLogEntryByIndexOK struct {
	Payload models.LogEntry
}

// IsSuccess returns true when this get log entry by index o k response has a 2xx status code
func (o *GetLogEntryByIndexOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log entry by index o k response has a 3xx status code
func (o *GetLogEntryByIndexOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log entry by index o k response has a 4xx status code
func (o *GetLogEntryByIndexOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log entry by index o k response has a 5xx status code
func (o *GetLogEntryByIndexOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log entry by index o k response a status code equal to that given
func (o *GetLogEntryByIndexOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get log entry by index o k response
func (o *GetLogEntryByIndexOK) Code() int {
	return 200
}

func (o *GetLogEntryByIndexOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndexOK %s", 200, payload)
}

func (o *GetLogEntryByIndexOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndexOK %s", 200, payload)
}

func (o *GetLogEntryByIndexOK) GetPayload() models.LogEntry {
	return o.Payload
}

func (o *GetLogEntryByIndexOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogEntryByIndexNotFound creates a GetLogEntryByIndexNotFound with default headers values
func NewGetLogEntryByIndexNotFound() *GetLogEntryByIndexNotFound {
	return &GetLogEntryByIndexNotFound{}
}

/*
GetLogEntryByIndexNotFound describes a response with status code 404, with default header values.

The content requested could not be found
*/
type GetLogEntryByIndexNotFound struct {
}

// IsSuccess returns true when this get log entry by index not found response has a 2xx status code
func (o *GetLogEntryByIndexNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log entry by index not found response has a 3xx status code
func (o *GetLogEntryByIndexNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log entry by index not found response has a 4xx status code
func (o *GetLogEntryByIndexNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log entry by index not found response has a 5xx status code
func (o *GetLogEntryByIndexNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get log entry by index not found response a status code equal to that given
func (o *GetLogEntryByIndexNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get log entry by index not found response
func (o *GetLogEntryByIndexNotFound) Code() int {
	return 404
}

func (o *GetLogEntryByIndexNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndexNotFound", 404)
}

func (o *GetLogEntryByIndexNotFound) String() string {
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndexNotFound", 404)
}

func (o *GetLogEntryByIndexNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogEntryByIndexDefault creates a GetLogEntryByIndexDefault with default headers values
func NewGetLogEntryByIndexDefault(code int) *GetLogEntryByIndexDefault {
	return &GetLogEntryByIndexDefault{
		_statusCode: code,
	}
}

/*
GetLogEntryByIndexDefault describes a response with status code -1, with default header values.

There was an internal error in the server while processing the request
*/
type GetLogEntryByIndexDefault struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this get log entry by index default response has a 2xx status code
func (o *GetLogEntryByIndexDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get log entry by index default response has a 3xx status code
func (o *GetLogEntryByIndexDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get log entry by index default response has a 4xx status code
func (o *GetLogEntryByIndexDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get log entry by index default response has a 5xx status code
func (o *GetLogEntryByIndexDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get log entry by index default response a status code equal to that given
func (o *GetLogEntryByIndexDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get log entry by index default response
func (o *GetLogEntryByIndexDefault) Code() int {
	return o._statusCode
}

func (o *GetLogEntryByIndexDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndex default %s", o._statusCode, payload)
}

func (o *GetLogEntryByIndexDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /api/v1/log/entries][%d] getLogEntryByIndex default %s", o._statusCode, payload)
}

func (o *GetLogEntryByIndexDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLogEntryByIndexDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
