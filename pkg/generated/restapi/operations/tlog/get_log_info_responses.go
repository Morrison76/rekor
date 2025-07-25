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

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Morrison76/rekor/pkg/generated/models"
)

// GetLogInfoOKCode is the HTTP code returned for type GetLogInfoOK
const GetLogInfoOKCode int = 200

/*
GetLogInfoOK A JSON object with the root hash and tree size as properties

swagger:response getLogInfoOK
*/
type GetLogInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.LogInfo `json:"body,omitempty"`
}

// NewGetLogInfoOK creates GetLogInfoOK with default headers values
func NewGetLogInfoOK() *GetLogInfoOK {

	return &GetLogInfoOK{}
}

// WithPayload adds the payload to the get log info o k response
func (o *GetLogInfoOK) WithPayload(payload *models.LogInfo) *GetLogInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log info o k response
func (o *GetLogInfoOK) SetPayload(payload *models.LogInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetLogInfoDefault There was an internal error in the server while processing the request

swagger:response getLogInfoDefault
*/
type GetLogInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLogInfoDefault creates GetLogInfoDefault with default headers values
func NewGetLogInfoDefault(code int) *GetLogInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLogInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get log info default response
func (o *GetLogInfoDefault) WithStatusCode(code int) *GetLogInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get log info default response
func (o *GetLogInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get log info default response
func (o *GetLogInfoDefault) WithPayload(payload *models.Error) *GetLogInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get log info default response
func (o *GetLogInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLogInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
