/*
Copyright The Rekor Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/Morrison76/rekor/pkg/generated/models"
	"github.com/Morrison76/rekor/pkg/generated/restapi/operations/pubkey"
)

func GetPublicKeyHandler(params pubkey.GetPublicKeyParams) middleware.Responder {
	treeID := swag.StringValue(params.TreeID)
	pk, err := api.logRanges.PublicKey(treeID)
	if err != nil {
		return handleRekorAPIError(params, http.StatusBadRequest, err, "")
	}
	return pubkey.NewGetPublicKeyOK().WithPayload(pk)
}

// handlers for APIs that may be disabled in a given instance

func GetPublicKeyNotImplementedHandler(_ pubkey.GetPublicKeyParams) middleware.Responder {
	err := &models.Error{
		Code:    http.StatusNotImplemented,
		Message: "Get Public Key API not enabled in this Rekor instance",
	}

	return pubkey.NewGetPublicKeyDefault(http.StatusNotImplemented).WithPayload(err)
}
